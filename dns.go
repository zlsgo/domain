package domain

import (
	"context"
)

// GetDns get domain dns
func GetDns(ctx context.Context, domain string, dns ...string) ([]string, error) {
	return getDns(ctx, domain, "ip", dns...)
}

func GetDnsIPv4(ctx context.Context, domain string, dns ...string) ([]string, error) {
	return getDns(ctx, domain, "ip4", dns...)
}

func GetDnsIPv6(ctx context.Context, domain string, dns ...string) ([]string, error) {
	return getDns(ctx, domain, "ip6", dns...)
}

func getDns(ctx context.Context, domain string, network string, dns ...string) ([]string, error) {
	pDomain := parseDomain(domain)

	ipAddrs, err := getResolver(ctx, dns...).LookupIP(ctx, network, pDomain)
	if err != nil {
		return nil, err
	}

	if len(ipAddrs) == 0 {
		return []string{}, nil
	}

	ips := make([]string, len(ipAddrs))
	for i, ip := range ipAddrs {
		ips[i] = ip.String()
	}

	return ips, nil
}
