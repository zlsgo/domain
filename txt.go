package domain

import (
	"context"
)

func GetTxt(ctx context.Context, domain string, dns ...string) ([]string, error) {
	r := getResolver(ctx, dns...)
	txts, err := r.LookupTXT(ctx, parseDomain(domain))
	if err != nil {
		return []string{}, err
	}

	if len(txts) == 0 {
		return []string{}, nil
	}

	return txts, nil
}
