package main

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"phish.report/IOK/schema"
	"strings"
)

func main() {
	l, err := os.Create("logsource.yml")
	if err != nil {
		panic(err)
	}
	defer l.Close()

	i, err := os.Create("input.go")
	if err != nil {
		panic(err)
	}
	defer i.Close()

	l.WriteString(`title: Sigma config for use with the phish.report/IOK library

backends:
  - github.com/bradleyjkemp/sigma-go

fieldmappings:
`)

	i.WriteString(`package iok

import "net/url"

//go:generate go run ./internal/genlogsource
//go:generate go fmt input.go

type Input struct {
`)

	for _, field := range schema.Fields {
		goIdentifier := toGoIdentifier(field.SigmaName, field.Derived)
		jsonPath := "$." + goIdentifier
		if field.Type == schema.StringList {
			jsonPath += "[*]"
		}
		fmt.Fprintf(l, "  %s: %s\n", field.SigmaName, jsonPath)

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
		fmt.Fprintf(i, "  %s %s // %s\n", goIdentifier, gotype, field.Description)
	}
	i.WriteString("}\n")
}

func toGoIdentifier(s, derived string) string {
	parts := strings.Split(s, ".")
	for i, part := range parts {
		parts[i] = cases.Title(language.English, cases.Compact).String(part)
	}
	if derived != "" {
		parts[0] = strings.ToLower(parts[0])
	}
	return strings.Join(parts, "")
}
