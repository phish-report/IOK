package iok

import (
	"context"
	"github.com/bradleyjkemp/cupaloy/v2"
	"net/http"
	"net/url"
	"slices"
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
			slices.Sort(input.Title)
			slices.Sort(input.Title)
			slices.Sort(input.JS)
			slices.Sort(input.CSS)
			slices.Sort(input.Cookie)
			slices.Sort(input.Header)
			slices.SortFunc(input.Request, func(a, b *url.URL) int {
				switch {
				case a.String() < b.String():
					return -1
				case a.String() == b.String():
					return 0
				default:
					return 1
				}
			})

			cupaloy.SnapshotT(t, input)
			t.Run("converted", func(t *testing.T) {
				cupaloy.SnapshotT(t, convertInput(input))
			})
		})
	}
}
