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

//go:generate go run ./internal/genlogsource

type Input struct {
`)

	for _, field := range schema.Fields {
		jsonPath := "$." + toGoIdentifier(field.SigmaName)
		if field.Type == schema.StringListField {
			jsonPath += "[*]"
		}
		fmt.Fprintf(l, "  %s: %s\n", field.SigmaName, jsonPath)

		var gotype string
		switch field.Type {
		case schema.StringListField:
			gotype = "[]string"
		default:
			gotype = "string"
		}
		fmt.Fprintf(i, "  %s %s // %s\n", toGoIdentifier(field.SigmaName), gotype, field.Description)
	}
	i.WriteString("}\n")
}

func toGoIdentifier(s string) string {
	parts := strings.Split(s, ".")
	for i, part := range parts {
		parts[i] = cases.Title(language.English, cases.Compact).String(part)
	}
	return strings.Join(parts, "")
}
