package domain

import (
	"context"
)

// GetDns performs a DNS lookup for the given domain and returns a slice of IP addresses.
func (c *Client) GetDns(ctx context.Context, domain string) ([]string, error) {
	return c.getDns(ctx, domain, "ip")
}

// GetDnsIPv4 performs a DNS lookup for the given domain and returns a slice of IPv4 addresses.
func (c *Client) GetDnsIPv4(ctx context.Context, domain string) ([]string, error) {
	return c.getDns(ctx, domain, "ip4")
}

// GetDnsIPv6 performs a DNS lookup for the given domain and returns a slice of IPv6 addresses.
func (c *Client) GetDnsIPv6(ctx context.Context, domain string) ([]string, error) {
	return c.getDns(ctx, domain, "ip6")
}

func (c *Client) getDns(ctx context.Context, domain string, network string) ([]string, error) {
	pDomain := parseDomain(domain)

	ipAddrs, err := c.resolver.LookupIP(ctx, network, pDomain)
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
