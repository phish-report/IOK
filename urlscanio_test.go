package iok

import (
	"context"
	"github.com/bradleyjkemp/cupaloy/v2"
	"net/http"
	"sort"
	"testing"
)

func TestInputFromURLScan(t *testing.T) {
	tests := []string{
		"be1f3938-ec14-4c0e-8f0a-45fb1ea423b3",
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

func sortField(f []string) {
	sort.Slice(f, func(i, j int) bool {
		return f[i] < f[j]
	})
}
