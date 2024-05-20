package main

import (
	"fmt"
	. "github.com/dave/jennifer/jen"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"phish.report/IOK/schema"
	"slices"
	"strings"
)

func main() {
	l, err := os.Create("logsource.yml")
	if err != nil {
		panic(err)
	}
	defer l.Close()

	input := NewFile("iok")
	input.Comment("//go:generate go run ./internal/genlogsource")
	input.Comment("//go:generate go fmt input.go")
	input.Comment("//go:generate go run ./internal/gendocs")

	l.WriteString(`title: Sigma config for use with the phish.report/IOK library

backends:
  - github.com/bradleyjkemp/sigma-go

fieldmappings:
`)

	mappings := map[string]string{}

	structFields := []Code{}
	setters := []Code{}
	for _, field := range schema.Fields {
		goIdentifier := toGoIdentifier(field.SigmaName, field.Deprecated)
		jsonPath := "$." + goIdentifier
		if field.Type == schema.StringList {
			jsonPath += "[*]"
		}
		mappings[field.SigmaName] = jsonPath

		if field.Alias != "" {
			jsonPath = mappings[field.Alias]
		}
		fmt.Fprintf(l, "  %s: %s\n", field.SigmaName, jsonPath)

		if field.Alias != "" {
			continue // don't add this field to the go struct
		}
		gotype := field.GoType
		if field.GoType == "" {
			switch field.Type {
			case schema.String:
				gotype = "string"
			case schema.StringList:
				gotype = "[]string"
			case schema.Number:
				gotype = "int"
			case schema.NumberList:
				gotype = "[]int"

			default:
				panic("unknown field type: " + field.Type)
			}
		}

		comment := field.Description
		if field.Derived != "" {
			comment += fmt.Sprintf(" (derived from %s)", field.Derived)
		}

		structFields = append(structFields, Id(goIdentifier).Id(gotype).Comment(comment))

		if field.Alias != "" {
			// don't generate a setter
			continue
		}

		setterName := ""
		switch {
		case field.Derived != "" && strings.HasPrefix(gotype, "[]"):
			setterName = "add" + goIdentifier
		case field.Derived == "" && strings.HasPrefix(gotype, "[]"):
			setterName = "Add" + goIdentifier
		case field.Derived != "" && !strings.HasPrefix(gotype, "[]"):
			setterName = "set" + goIdentifier
		case field.Derived == "" && !strings.HasPrefix(gotype, "[]"):
			setterName = "Set" + goIdentifier
		}

		setters = append(setters,
			Func().Params(Id("i").Id("*Input")).Id(setterName).Do(func(s *Statement) {
				if strings.HasPrefix(gotype, "[]") {
					s.Params(Id("v").Id("..." + strings.TrimPrefix(gotype, "[]")))
				} else {
					s.Params(Id("v").Id(gotype))
				}
			}).
				BlockFunc(func(g *Group) {
					g.Id("i").Dot(goIdentifier).Op("=").Do(func(s *Statement) {
						if strings.HasPrefix(gotype, "[]") {
							s.Append(Id("i").Dot(goIdentifier), Id("v").Op("..."))
						} else {
							s.Id("v")
						}
					})
					// This field is now set in the input, ensure we mark it as supported
					g.Id("i").Dot("Supported").Op("=").Append(Id("i").Dot("Supported"), Lit(field.SigmaName))
					g.Qual("slices", "Sort").Call(Id("i").Dot("Supported"))
					g.Id("i").Dot("Supported").Op("=").Qual("slices", "Compact").Call(Id("i").Dot("Supported"))
					if field.PostSetter != nil {
						field.PostSetter(g)
					}
				}))
	}

	input.Type().Id("Input").Struct(
		slices.Concat([]Code{
			Id("ipToASN").Func().Params(Qual("context", "Context"), Qual("net", "IP")).Id("int"),
		}, structFields)...)
	for _, setter := range setters {
		input.Add(setter)
	}

	err = input.Save("input.go")
	if err != nil {
		panic(err)
	}
}

func toGoIdentifier(s string, private bool) string {
	parts := strings.Split(s, ".")
	for i, part := range parts {
		if len(part) <= 3 || part == "cname" || part == "html" {
			parts[i] = cases.Upper(language.English, cases.Compact).String(part)
		} else {
			parts[i] = cases.Title(language.English, cases.Compact).String(part)
		}
	}
	if private {
		// lowercasing makes this unexported
		parts[0] = strings.ToLower(parts[0])
	}
	return strings.Join(parts, "")
}
