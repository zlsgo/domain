package domain_test

import (
	"context"
	"errors"
	"net"
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

func TestGetDnsWithLiteralIPv4URL(t *testing.T) {
	client := domain.NewClient()
	ips, err := client.GetDns(context.Background(), "http://1.1.1.1:8080/path")
	if err != nil {
		t.Fatal(err)
	}
	if len(ips) != 1 || ips[0] != "1.1.1.1" {
		t.Fatalf("expected [1.1.1.1], but got %v", ips)
	}
}

func TestGetDnsIPv6WithLiteralIPv4ReturnsAddrError(t *testing.T) {
	client := domain.NewClient()
	_, err := client.GetDnsIPv6(context.Background(), "1.1.1.1")
	if err == nil {
		t.Fatal("expected error, but got nil")
	}

	var addrErr *net.AddrError
	if !errors.As(err, &addrErr) {
		t.Fatalf("expected net.AddrError, but got %T", err)
	}
}

func TestGetDnsWithLiteralIPv6URL(t *testing.T) {
	client := domain.NewClient()
	ips, err := client.GetDns(context.Background(), "https://[2606:4700:4700::1111]:8443/dns-query")
	if err != nil {
		t.Fatal(err)
	}
	if len(ips) != 1 || ips[0] != "2606:4700:4700::1111" {
		t.Fatalf("expected [2606:4700:4700::1111], but got %v", ips)
	}
}

func TestGetDnsIPv4WithLiteralIPv4URL(t *testing.T) {
	client := domain.NewClient()
	ips, err := client.GetDnsIPv4(context.Background(), "http://1.1.1.1:8080/path")
	if err != nil {
		t.Fatal(err)
	}
	if len(ips) != 1 || ips[0] != "1.1.1.1" {
		t.Fatalf("expected [1.1.1.1], but got %v", ips)
	}
}

func TestGetDnsIPv4WithLiteralIPv6URLReturnsAddrError(t *testing.T) {
	client := domain.NewClient()
	_, err := client.GetDnsIPv4(context.Background(), "https://[2606:4700:4700::1111]:8443/dns-query")
	if err == nil {
		t.Fatal("expected error, but got nil")
	}

	var addrErr *net.AddrError
	if !errors.As(err, &addrErr) {
		t.Fatalf("expected net.AddrError, but got %T", err)
	}
}
