package domain

import (
	"context"
	"net"
)

// GetNS retrieves the NS records for a given domain.
func (c *Client) GetNS(ctx context.Context, domain string) ([]*net.NS, error) {
	nss, err := c.resolver.LookupNS(ctx, parseDomain(domain))
	if err != nil {
		return nil, err
	}

	if len(nss) == 0 {
		return []*net.NS{}, nil
	}

	return nss, nil
}
