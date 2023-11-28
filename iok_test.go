package iok

import (
	"io/fs"
	"testing"
)

// A dummy test just to ensure that the iok.go init functions don't panic
func Test(t *testing.T) {}

func BenchmarkInit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fs.WalkDir(indicators, ".", func(path string, d fs.DirEntry, err error) error {
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
	}
}
