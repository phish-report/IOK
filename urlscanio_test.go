package iok

import (
	"context"
	"github.com/bradleyjkemp/cupaloy/v2"
	"github.com/bradleyjkemp/sigma-go"
	"net/http"
	"sort"
	"testing"
)

func TestInputFromURLScan(t *testing.T) {
	tests := []string{
		"be1f3938-ec14-4c0e-8f0a-45fb1ea423b3",
		"55545c59-1b1c-452b-9b33-0b5d63a0d825",
	}
	for _, tt := range tests {
		t.Run(tt, func(t *testing.T) {
			input, err := InputFromURLScan(context.Background(), tt, http.DefaultClient)
			if err != nil {
				t.Fatal(err)
			}

			// because resources are fetched in parallel the output is non-deterministic, so we need to sort each field
			sortField(input.Title)
			sortField(input.JS)
			sortField(input.CSS)
			sortField(input.Cookies)
			sortField(input.Headers)
			sortField(input.Requests)

			cupaloy.SnapshotT(t, input)
		})
	}
}

func BenchmarkUrlscanGetMatches(b *testing.B) {
	b.StopTimer()
	input, err := InputFromURLScan(context.Background(), "67514436-4198-46c1-8e8b-5ddbc03098f2", http.DefaultClient)
	if err != nil {
		b.Fatal(err)
	}
	b.StartTimer()
	var matches []sigma.Rule
	for i := 0; i < b.N; i++ {
		matches, err = GetMatches(input)
		if err != nil {
			b.Fatal(err)
		}
	}
	if len(matches) == 0 {
		b.Fatal(0)
	}
}

func sortField(f []string) {
	sort.Slice(f, func(i, j int) bool {
		return f[i] < f[j]
	})
}
