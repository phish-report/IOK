package iok

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"net/http"
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
	Dom     string
	Cookies []http.Cookie
}

func GetMatches(input Input) ([]sigma.Rule, error) {
	matches := []sigma.Rule{}
	for _, evaluator := range evaluators {
		result, err := evaluator.Matches(context.Background(), map[string]interface{}{
			"dom":     input.Dom,
			"cookies": input.Cookies,
		})
		if err != nil {
			return nil, fmt.Errorf("error evaluating %s: %w", evaluator.Title, err)
		}

		if result.Match {
			matches = append(matches, evaluator.Rule)
		}
	}
	return matches, nil
}

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
		evaluators = append(evaluators, evaluator.ForRule(rule, evaluator.WithConfig(config)))
		return nil
	})
	if err != nil {
		panic(err)
	}
}
