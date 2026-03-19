package domain

import (
	"context"
	"testing"
)

func TestLookupAddr(t *testing.T) {
	client := &Client{
		resolver: testResolver{
			lookupAddr: func(context.Context, string) ([]string, error) {
				return []string{"dns.google."}, nil
			},
		},
	}
	names, err := client.LookupAddr(context.Background(), "8.8.8.8")
	if err != nil {
		t.Fatal(err)
	}
	if len(names) == 0 {
		t.Fatal("expected names, but got none")
	}
	t.Log(names)
}
