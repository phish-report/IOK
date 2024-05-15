package iok

import (
	"context"
	"embed"
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"io/fs"
	"path/filepath"
	"phish.report/IOK/schema"
	"reflect"
	"strings"

	"github.com/bradleyjkemp/sigma-go"
	"github.com/bradleyjkemp/sigma-go/evaluator"
)

//go:embed indicators
var indicators embed.FS

//go:embed logsource.yml
var config []byte

var evaluators []*evaluator.RuleEvaluator

func GetMatches(input Input) ([]sigma.Rule, error) {
	return GetMatchesForRules(input, evaluators)
}

func GetMatchesForRules(input Input, rules []*evaluator.RuleEvaluator) ([]sigma.Rule, error) {
	matches := []sigma.Rule{}
	ruleInput := convertInput(input)
	var errs []error
	for _, rule := range rules {
		result, err := rule.Matches(context.Background(), ruleInput)
		if err != nil {
			errs = append(errs, fmt.Errorf("error evaluating %s: %w", rule.Title, err))
			continue
		}

		if result.Match {
			matches = append(matches, rule.Rule)
		}
	}
	return matches, errors.Join(errs...)
}

func convertInput(input Input) evaluator.Event {
	for _, field := range schema.Fields {
		supported := false
		for _, s := range input.Supported {
			if field.SigmaName == s {
				supported = true
				break
			}
		}
		if !supported {
			input.unsupported = append(input.unsupported, field.SigmaName)
		}
	}

	event := map[string]interface{}{}
	mapstructure.Decode(input, &event)
	// TODO: do we need to handle []something -> []interface{} conversion?

	// Convert slices into []interface{} (required for sigma-go to work properly)
	for key, val := range event {
		if reflect.TypeOf(val).Kind() != reflect.Slice {
			continue
		}
		event[key] = toInterfaceSlice(val)
	}
	return event
}

func toInterfaceSlice(s any) []interface{} {
	v := reflect.ValueOf(s)
	is := []interface{}{}
	for i := 0; i < v.Len(); i++ {
		is = append(is, v.Index(i).Interface())
	}
	return is
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
