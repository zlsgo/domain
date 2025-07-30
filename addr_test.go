package domain_test

import (
	"context"
	"testing"

	"github.com/zlsgo/domain"
)

func TestLookupAddr(t *testing.T) {
	client := domain.NewClient()
	names, err := client.LookupAddr(context.Background(), "8.8.8.8")
	if err != nil {
		t.Fatal(err)
	}
	if len(names) == 0 {
		t.Fatal("expected names, but got none")
	}
	t.Log(names)
}
