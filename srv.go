package domain

import (
	"context"
	"net"
)

// GetSRV retrieves the SRV records for a given service, proto, and name.
func (c *Client) GetSRV(ctx context.Context, service, proto, name string) (string, []*net.SRV, error) {
	cname, srvs, err := c.resolver.LookupSRV(ctx, service, proto, name)
	if err != nil {
		return "", nil, err
	}

	if len(srvs) == 0 {
		return cname, []*net.SRV{}, nil
	}

	return cname, srvs, nil
}
