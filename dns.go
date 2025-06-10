package domain

import (
	"context"
)

// GetDns get domain dns
func GetDns(ctx context.Context, domain string, dns ...string) ([]string, error) {
	r := getResolver(ctx, dns...)
	ips, err := r.LookupHost(ctx, domain)
	if err != nil {
		return nil, err
	}

	if len(ips) == 0 {
		return nil, nil
	}

	return ips, nil
}
