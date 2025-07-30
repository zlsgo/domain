package domain

import (
	"context"
)

// GetCNAME retrieves the CNAME record for a given domain.
func (c *Client) GetCNAME(ctx context.Context, domain string) (string, error) {
	cname, err := c.resolver.LookupCNAME(ctx, parseDomain(domain))
	if err != nil {
		return "", err
	}

	return cname, nil
}
