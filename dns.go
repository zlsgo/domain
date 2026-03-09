package domain

import (
	"context"
	"net"
	"sync"
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

// GetMulti performs DNS lookups for multiple domains concurrently.
// It returns a map where keys are the domains and values are their IP addresses.
// It also returns a slice of errors encountered during the lookups.
func (c *Client) GetMulti(ctx context.Context, domains []string) (map[string][]string, []error) {
	var wg sync.WaitGroup
	results := make(map[string][]string)
	mu := &sync.Mutex{}
	var errs []error

	for _, domain := range domains {
		wg.Add(1)
		go func(d string) {
			defer wg.Done()
			ips, err := c.GetDns(ctx, d)
			if err != nil {
				mu.Lock()
				errs = append(errs, err)
				mu.Unlock()
				return
			}
			mu.Lock()
			results[d] = ips
			mu.Unlock()
		}(domain)
	}

	wg.Wait()

	return results, errs
}

func (c *Client) getDns(ctx context.Context, domain string, network string) ([]string, error) {
	if ip := parseLiteralIP(domain); ip != nil {
		return literalIPResult(ip, network)
	}

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

func literalIPResult(ip net.IP, network string) ([]string, error) {
	switch network {
	case "ip":
		return []string{ip.String()}, nil
	case "ip4":
		if ip.To4() != nil {
			return []string{ip.String()}, nil
		}
	case "ip6":
		if ip.To4() == nil {
			return []string{ip.String()}, nil
		}
	}

	return nil, &net.AddrError{Err: "no suitable address found", Addr: ip.String()}
}
