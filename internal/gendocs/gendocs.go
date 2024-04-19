package main

import (
	"fmt"
	"os"
	"phish.report/IOK/schema"
)

func main() {
	s, err := os.Create("schema.md")
	if err != nil {
		panic(err)
	}
	defer s.Close()

	s.WriteString(`| Field | Description | Example |
|-------|-------------|---------|
`)

	for _, field := range schema.Fields {
		s.WriteString(fmt.Sprintf("|`%s`|%s|`%s`|\n", field.SigmaName, field.Description, field.Example))
	}
}
