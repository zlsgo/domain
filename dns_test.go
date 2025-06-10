package domain_test

import (
	"context"
	"testing"

	"github.com/sohaha/zlsgo"
	"github.com/zlsgo/domain"
)

func TestGetDns(t *testing.T) {
	tt := zlsgo.NewTest(t)
	ctx := context.Background()

	dns, err := domain.GetDns(ctx, "www.google.com")
	if err != nil {
		tt.Fatal(err)
	}
	tt.Log(dns)

	ips, err := domain.GetDns(ctx, "www.google.com", "8.8.8.8")
	if err != nil {
		tt.Fatal(err)
	}
	tt.Log(ips)
}
