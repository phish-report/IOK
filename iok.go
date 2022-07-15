package iok

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/bradleyjkemp/sigma-go"
	"github.com/bradleyjkemp/sigma-go/evaluator"
)

//go:embed indicators
var indicators embed.FS

//go:embed logsource.yml
var config []byte

var evaluators []*evaluator.RuleEvaluator

type Input struct {
	HTML     string
	JS       []string
	CSS      []string
	Cookies  []string
	Headers  []string
	Requests []string
}

func GetMatches(input Input) ([]sigma.Rule, error) {
	matches := []sigma.Rule{}
	for _, evaluator := range evaluators {
		result, err := evaluator.Matches(context.Background(), convertInput(input))
		if err != nil {
			return nil, fmt.Errorf("error evaluating %s: %w", evaluator.Title, err)
		}

		if result.Match {
			matches = append(matches, evaluator.Rule)
		}
	}
	return matches, nil
}

func convertInput(input Input) evaluator.Event {
	return map[string]interface{}{
		"html":     input.HTML,
		"js":       toInterfaceSlice(input.JS),
		"css":      toInterfaceSlice(input.CSS),
		"cookies":  toInterfaceSlice(input.Cookies),
		"headers":  toInterfaceSlice(input.Headers),
		"requests": toInterfaceSlice(input.Requests),
	}
}

func toInterfaceSlice(s []string) []interface{} {
	i := make([]interface{}, 0, len(s))
	for _, str := range s {
		i = append(i, str)
	}
	return i
}

var Rules = map[string]sigma.Rule{}
var RawRules = map[string][]byte{}

func init() {
	config, err := sigma.ParseConfig(config)
	if err != nil {
		panic(err)
	}

	err = fs.WalkDir(indicators, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		contents, _ := indicators.ReadFile(path)
		rule, err := sigma.ParseRule(contents)
		if err != nil {
			panic(err)
		}
		if rule.ID == "" {
			rule.ID, _, _ = strings.Cut(filepath.Base(path), ".")
		}

		Rules[rule.ID] = rule
		RawRules[rule.ID] = contents

		evaluators = append(evaluators, evaluator.ForRule(rule, evaluator.WithConfig(config)))
		return nil
	})
	if err != nil {
		panic(err)
	}
}
