package domain

import (
	"context"
)

// LookupAddr performs a reverse DNS lookup for the given IP address.
func (c *Client) LookupAddr(ctx context.Context, addr string) ([]string, error) {
	names, err := c.resolver.LookupAddr(ctx, addr)
	if err != nil {
		return nil, err
	}
	if len(names) == 0 {
		return []string{}, nil
	}
	return names, nil
}
