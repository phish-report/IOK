package main

import (
	"context"
	"fmt"
	"github.com/bradleyjkemp/sigma-go/evaluator"
	"log"
	"net/http"
	"os"
	"os/signal"
	iok "phish.report/IOK"
	"strings"
	"syscall"
	"time"
)

func main() {
	rules := os.Args[1:]
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()
	ctx, cancel = signal.NotifyContext(ctx, syscall.SIGTERM)
	defer cancel()

	failed := false
	for _, rule := range rules {
		if err := verifyRule(ctx, rule); err != nil {
			failed = true
			log.Println("Rule", rule, "failed to verify ❌ ", err)
		} else {
			log.Println("Rule", rule, "verified ✅")
		}
	}
	if failed {
		os.Exit(1)
	}
}

func verifyRule(ctx context.Context, ruleFile string) error {
	ruleContents, err := os.ReadFile(ruleFile)
	if err != nil {
		return fmt.Errorf("failed to read rule %s: %w", ruleFile, err)
	}

	rule, err := iok.ParseRule(ruleFile, ruleContents)
	if err != nil {
		return fmt.Errorf("failed to parse rule %s: %w", ruleFile, err)
	}

	for _, reference := range rule.References {
		if !strings.HasPrefix(reference, "https://urlscan.io/result/") {
			continue
		}

		uuid := strings.TrimPrefix(reference, "https://urlscan.io/result/")

		input, err := iok.InputFromURLScan(ctx, uuid, http.DefaultClient)
		if err != nil {
			return fmt.Errorf("failed to get IOK input %s: %w", uuid, err)
		}

		matches, err := iok.GetMatchesForRules(input, []*evaluator.RuleEvaluator{rule})
		if err != nil {
			return fmt.Errorf("failed to run IOK rule %s: %w", ruleFile, err)
		}

		if len(matches) == 0 {
			return fmt.Errorf("expected %s to match %s reference, but it didn't", ruleFile, reference)
		}
	}
	return nil // the rule matched all the urlscan.io results it references
}
