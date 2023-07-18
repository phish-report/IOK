package iok

import (
	"context"
	"github.com/bradleyjkemp/cupaloy/v2"
	"net/http"
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
			cupaloy.SnapshotT(t, input)
		})
	}
}
