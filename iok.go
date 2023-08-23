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
	Title    string   // Title is the title of the page as it would be shown in a browser
	Hostname string   // Hostname is the hostname that the page was served from
	DOM      string   // DOM contains the HTML contents of the primary page *after* it has loaded
	HTML     string   // HTML contains the HTML response of the primary page
	JS       []string // JS contains all JavaScript on the page, whether an embedded script or loaded from a file
	CSS      []string // CSS contains all CSS on the page, whether an embedded stylesheet or loaded from a file
	Cookies  []string // Cookies contains all cookies set both by the initial page load and any subsequent requests
	Headers  []string // Headers contains the headers of the initial page load
	Requests []string // Requests contains a list of all URLs requested during the page load
}

func GetMatches(input Input) ([]sigma.Rule, error) {
	return GetMatchesForRules(input, evaluators)
}

func GetMatchesForRules(input Input, rules []*evaluator.RuleEvaluator) ([]sigma.Rule, error) {
	matches := []sigma.Rule{}
	ruleInput := convertInput(input)
	for _, rule := range rules {
		result, err := rule.Matches(context.Background(), ruleInput)
		if err != nil {
			return nil, fmt.Errorf("error evaluating %s: %w", rule.Title, err)
		}

		if result.Match {
			matches = append(matches, rule.Rule)
		}
	}
	return matches, nil
}

func convertInput(input Input) evaluator.Event {
	return map[string]interface{}{
		"title":    input.Title,
		"hostname": input.Hostname,
		"dom":      input.DOM,
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

var Rules = map[string]*evaluator.RuleEvaluator{}
var RawRules = map[string][]byte{}

func ParseRule(path string, contents []byte) (*evaluator.RuleEvaluator, error) {
	config, err := sigma.ParseConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}
	rule, err := sigma.ParseRule(contents)
	if err != nil {
		return nil, fmt.Errorf("failed to parse rule %s: %w", path, err)
	}
	if rule.ID == "" {
		rule.ID, _, _ = strings.Cut(filepath.Base(path), ".")
	}

	return evaluator.ForRule(rule, evaluator.WithConfig(config), evaluator.CaseSensitive), nil
}

func init() {
	err := fs.WalkDir(indicators, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		contents, _ := indicators.ReadFile(path)

		rule, err := ParseRule(path, contents)
		if err != nil {
			return err
		}
		Rules[rule.ID] = rule
		RawRules[rule.ID] = contents

		evaluators = append(evaluators, rule)
		return nil
	})
	if err != nil {
		panic(err)
	}
}
