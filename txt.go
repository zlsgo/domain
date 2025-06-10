package domain

import (
	"context"
)

// GetTxt retrieves the TXT records for a given domain.
func (c *Client) GetTxt(ctx context.Context, domain string) ([]string, error) {
	txts, err := c.resolver.LookupTXT(ctx, parseDomain(domain))
	if err != nil {
		return nil, err
	}

	if len(txts) == 0 {
		return []string{}, nil
	}

	return txts, nil
}
