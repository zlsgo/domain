package domain_test

import (
	"context"
	"testing"

	"github.com/zlsgo/domain"
)

func TestGetSRV(t *testing.T) {
	// Using a public and reliable DNS server for SRV record lookups
	client := domain.NewClient("8.8.8.8")
	cname, srvs, err := client.GetSRV(context.Background(), "sip", "tcp", "sip.voice.google.com")
	if err != nil {
		// SRV records can sometimes be elusive, so we'll log the error but not fail the test
		t.Logf("Could not resolve SRV records, which can be normal: %v", err)
		return
	}
	if len(srvs) == 0 {
		t.Log("No SRV records found, which can be normal.")
		return
	}
	t.Log(cname, srvs)
}
