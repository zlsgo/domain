package domain_test

import (
	"context"
	"strings"
	"testing"

	"github.com/zlsgo/domain"
)

func TestGetDns(t *testing.T) {
	ctx := context.Background()

	client := domain.NewClient()
	dns, err := client.GetDns(ctx, "https://www.google.com")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(dns)

	dns, err = client.GetDns(ctx, "https://www.google.com/xxx")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(dns)

	dns, err = client.GetDns(ctx, "www.google.com")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(dns)

	client = domain.NewClient("8.8.8.8")
	ips, err := client.GetDns(ctx, "www.google.com")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ips)

	ips, err = client.GetDnsIPv4(ctx, "www.google.com")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ips)

	ips, err = client.GetDnsIPv6(ctx, "www.google.com")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ips)
}

func TestGetCNAME(t *testing.T) {
	client := domain.NewClient()
	cname, err := client.GetCNAME(context.Background(), "www.github.com")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasSuffix(cname, "github.com.") {
		t.Fatalf("expected CNAME to end with github.com, but got %s", cname)
	}
	t.Log(cname)
}

func TestGetMX(t *testing.T) {
	client := domain.NewClient("8.8.8.8")
	mxs, err := client.GetMX(context.Background(), "qq.com")
	if err != nil {
		t.Fatal(err)
	}
	if len(mxs) == 0 {
		t.Fatal("expected MX records, but got none")
	}
	t.Log(mxs)
}

func TestGetNS(t *testing.T) {
	client := domain.NewClient("8.8.8.8")
	nss, err := client.GetNS(context.Background(), "qq.com")
	if err != nil {
		t.Fatal(err)
	}
	if len(nss) == 0 {
		t.Fatal("expected NS records, but got none")
	}
	t.Log(nss)
}

func TestGetMulti(t *testing.T) {
	client := domain.NewClient()
	domains := []string{"www.google.com", "www.github.com", "invalid-domain-for-test"}
	results, errs := client.GetMulti(context.Background(), domains)
	if len(errs) != 1 {
		t.Fatalf("expected 1 error, but got %d", len(errs))
	}
	if len(results) != 2 {
		t.Fatalf("expected 2 results, but got %d", len(results))
	}
	t.Log(results)
	t.Log(errs)
}
