package iok

import (
	"context"
	"github.com/bradleyjkemp/cupaloy/v2"
	"net/http"
	"sort"
	"testing"
)

func TestInputsFromURLScan(t *testing.T) {
	tests := []string{
		"4392cf46-0bfc-406d-9bdc-6a53ad08a4b0",
		"be1f3938-ec14-4c0e-8f0a-45fb1ea423b3",
		"55545c59-1b1c-452b-9b33-0b5d63a0d825",
	}
	for _, tt := range tests {
		t.Run(tt, func(t *testing.T) {
			inputs, err := InputsFromURLScan(context.Background(), tt, http.DefaultClient)
			if err != nil {
				t.Fatal(err)
			}

			for _, input := range inputs {
				// because resources are fetched in parallel the output is non-deterministic, so we need to sort each field
				sortField(input.Title)
				sortField(input.JS)
				sortField(input.CSS)
				sortField(input.Cookies)
				sortField(input.Headers)
				sortField(input.Requests)
			}

			cupaloy.SnapshotT(t, inputs)
		})
	}
}

func sortField(f []string) {
	sort.Slice(f, func(i, j int) bool {
		return f[i] < f[j]
	})
}
