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

	client := domain.NewClient()
	dns, err := client.GetDns(ctx, "https://www.google.com")
	if err != nil {
		tt.Fatal(err)
	}
	tt.Log(dns)

	dns, err = client.GetDns(ctx, "https://www.google.com/xxx")
	if err != nil {
		tt.Fatal(err)
	}
	tt.Log(dns)

	dns, err = client.GetDns(ctx, "www.google.com")
	if err != nil {
		tt.Fatal(err)
	}
	tt.Log(dns)

	client = domain.NewClient("8.8.8.8")
	ips, err := client.GetDns(ctx, "www.google.com")
	if err != nil {
		tt.Fatal(err)
	}
	tt.Log(ips)

	ips, err = client.GetDnsIPv4(ctx, "www.google.com")
	if err != nil {
		tt.Fatal(err)
	}
	tt.Log(ips)

	ips, err = client.GetDnsIPv6(ctx, "www.google.com")
	if err != nil {
		tt.Fatal(err)
	}
	tt.Log(ips)
}
