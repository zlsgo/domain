package domain

import (
	"context"
	"errors"
	"net"
	"strings"
	"testing"
)

func TestGetDns(t *testing.T) {
	ctx := context.Background()
	client := &Client{
		resolver: testResolver{
			lookupIP: func(_ context.Context, network, host string) ([]net.IP, error) {
				if host != "www.google.com" {
					t.Fatalf("unexpected host: %s", host)
				}

				switch network {
				case "ip":
					return []net.IP{net.ParseIP("1.1.1.1"), net.ParseIP("2606:4700:4700::1111")}, nil
				case "ip4":
					return []net.IP{net.ParseIP("1.1.1.1")}, nil
				case "ip6":
					return []net.IP{net.ParseIP("2606:4700:4700::1111")}, nil
				default:
					t.Fatalf("unexpected network: %s", network)
					return nil, nil
				}
			},
		},
	}

	dns, err := client.GetDns(ctx, "https://www.google.com")
	if err != nil {
		t.Fatal(err)
	}
	if len(dns) != 2 {
		t.Fatalf("expected 2 IPs, but got %d", len(dns))
	}

	dns, err = client.GetDns(ctx, "https://www.google.com/xxx")
	if err != nil {
		t.Fatal(err)
	}
	if len(dns) != 2 {
		t.Fatalf("expected 2 IPs, but got %d", len(dns))
	}

	dns, err = client.GetDns(ctx, "www.google.com")
	if err != nil {
		t.Fatal(err)
	}
	if len(dns) != 2 {
		t.Fatalf("expected 2 IPs, but got %d", len(dns))
	}

	ips, err := client.GetDnsIPv4(ctx, "www.google.com")
	if err != nil {
		t.Fatal(err)
	}
	if len(ips) != 1 || ips[0] != "1.1.1.1" {
		t.Fatalf("expected [1.1.1.1], but got %v", ips)
	}

	ips, err = client.GetDnsIPv6(ctx, "www.google.com")
	if err != nil {
		t.Fatal(err)
	}
	if len(ips) != 1 || ips[0] != "2606:4700:4700::1111" {
		t.Fatalf("expected [2606:4700:4700::1111], but got %v", ips)
	}
}

func TestGetCNAME(t *testing.T) {
	client := NewClient()
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
	client := &Client{
		resolver: testResolver{
			lookupMX: func(context.Context, string) ([]*net.MX, error) {
				return []*net.MX{{Host: "mx.qq.com.", Pref: 10}}, nil
			},
		},
	}

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
	client := &Client{
		resolver: testResolver{
			lookupNS: func(context.Context, string) ([]*net.NS, error) {
				return []*net.NS{{Host: "ns1.qq.com."}}, nil
			},
		},
	}

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
	client := &Client{
		resolver: testResolver{
			lookupIP: func(_ context.Context, network, host string) ([]net.IP, error) {
				switch host {
				case "www.google.com":
					return []net.IP{net.ParseIP("1.1.1.1")}, nil
				case "www.github.com":
					return []net.IP{net.ParseIP("2.2.2.2")}, nil
				default:
					return nil, &net.DNSError{Err: "no such host", Name: host, IsNotFound: true}
				}
			},
		},
	}

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
	client := NewClient()
	ips, err := client.GetDns(context.Background(), "http://1.1.1.1:8080/path")
	if err != nil {
		t.Fatal(err)
	}
	if len(ips) != 1 || ips[0] != "1.1.1.1" {
		t.Fatalf("expected [1.1.1.1], but got %v", ips)
	}
}

func TestGetDnsIPv6WithLiteralIPv4ReturnsAddrError(t *testing.T) {
	client := NewClient()
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
	client := NewClient()
	ips, err := client.GetDns(context.Background(), "https://[2606:4700:4700::1111]:8443/dns-query")
	if err != nil {
		t.Fatal(err)
	}
	if len(ips) != 1 || ips[0] != "2606:4700:4700::1111" {
		t.Fatalf("expected [2606:4700:4700::1111], but got %v", ips)
	}
}

func TestGetDnsIPv4WithLiteralIPv4URL(t *testing.T) {
	client := NewClient()
	ips, err := client.GetDnsIPv4(context.Background(), "http://1.1.1.1:8080/path")
	if err != nil {
		t.Fatal(err)
	}
	if len(ips) != 1 || ips[0] != "1.1.1.1" {
		t.Fatalf("expected [1.1.1.1], but got %v", ips)
	}
}

func TestGetDnsIPv4WithLiteralIPv6URLReturnsAddrError(t *testing.T) {
	client := NewClient()
	_, err := client.GetDnsIPv4(context.Background(), "https://[2606:4700:4700::1111]:8443/dns-query")
	if err == nil {
		t.Fatal("expected error, but got nil")
	}

	var addrErr *net.AddrError
	if !errors.As(err, &addrErr) {
		t.Fatalf("expected net.AddrError, but got %T", err)
	}
}

type testResolver struct {
	lookupAddr  func(ctx context.Context, addr string) ([]string, error)
	lookupCNAME func(ctx context.Context, host string) (string, error)
	lookupIP    func(ctx context.Context, network, host string) ([]net.IP, error)
	lookupMX    func(ctx context.Context, name string) ([]*net.MX, error)
	lookupNS    func(ctx context.Context, name string) ([]*net.NS, error)
	lookupSRV   func(ctx context.Context, service, proto, name string) (string, []*net.SRV, error)
	lookupTXT   func(ctx context.Context, name string) ([]string, error)
}

func (r testResolver) LookupAddr(ctx context.Context, addr string) ([]string, error) {
	return r.lookupAddr(ctx, addr)
}

func (r testResolver) LookupCNAME(ctx context.Context, host string) (string, error) {
	return r.lookupCNAME(ctx, host)
}

func (r testResolver) LookupIP(ctx context.Context, network, host string) ([]net.IP, error) {
	return r.lookupIP(ctx, network, host)
}

func (r testResolver) LookupMX(ctx context.Context, name string) ([]*net.MX, error) {
	return r.lookupMX(ctx, name)
}

func (r testResolver) LookupNS(ctx context.Context, name string) ([]*net.NS, error) {
	return r.lookupNS(ctx, name)
}

func (r testResolver) LookupSRV(ctx context.Context, service, proto, name string) (string, []*net.SRV, error) {
	return r.lookupSRV(ctx, service, proto, name)
}

func (r testResolver) LookupTXT(ctx context.Context, name string) ([]string, error) {
	return r.lookupTXT(ctx, name)
}
