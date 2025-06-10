package domain

import (
	"context"
	"strings"
)

// GetDns performs a DNS lookup for the given domain and returns all available IP addresses (both IPv4 and IPv6).
func GetDns(ctx context.Context, domain string, dns ...string) ([]string, error) {
	return NewClient(dns...).GetDns(ctx, domain)
}

// GetDnsIPv4 performs a DNS lookup for the given domain and returns only IPv4 addresses.
func GetDnsIPv4(ctx context.Context, domain string, dns ...string) ([]string, error) {
	return NewClient(dns...).GetDnsIPv4(ctx, domain)
}

// GetDnsIPv6 performs a DNS lookup for the given domain and returns only IPv6 addresses.
func GetDnsIPv6(ctx context.Context, domain string, dns ...string) ([]string, error) {
	return NewClient(dns...).GetDnsIPv6(ctx, domain)
}

// GetTxt retrieves all TXT records for the given domain.
func GetTxt(ctx context.Context, domain string, dns ...string) ([]string, error) {
	return NewClient(dns...).GetTxt(ctx, domain)
}

func parseDomain(url string) string {
	s := strings.Split(url, "://")
	if len(s) == 2 {
		return strings.SplitN(s[1], "/", 2)[0]
	}
	return strings.SplitN(url, "/", 2)[0]
}
