package domain

import (
	"context"
)

func GetTxt(ctx context.Context, domain string, dns ...string) ([]string, error) {
	r := getResolver(ctx, dns...)
	txts, err := r.LookupTXT(ctx, domain)
	if err != nil {
		return nil, err
	}

	if len(txts) == 0 {
		return nil, nil
	}

	return txts, nil
}
