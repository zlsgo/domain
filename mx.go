package domain

import (
	"context"
	"net"
)

// GetMX retrieves the MX records for a given domain.
func (c *Client) GetMX(ctx context.Context, domain string) ([]*net.MX, error) {
	mxs, err := c.resolver.LookupMX(ctx, parseDomain(domain))
	if err != nil {
		return nil, err
	}

	if len(mxs) == 0 {
		return []*net.MX{}, nil
	}

	return mxs, nil
}
