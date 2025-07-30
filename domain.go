package domain

import (
	"context"
	"net"
	"strings"
)

var defaultClient = NewClient()

// GetDns performs a DNS lookup for the given domain and returns a slice of IP addresses.
// It uses a default client. For custom DNS servers or more control, use NewClient().
func GetDns(ctx context.Context, domain string) ([]string, error) {
	return defaultClient.GetDns(ctx, domain)
}

// GetDnsIPv4 performs a DNS lookup for the given domain and returns a slice of IPv4 addresses.
func GetDnsIPv4(ctx context.Context, domain string) ([]string, error) {
	return defaultClient.GetDnsIPv4(ctx, domain)
}

// GetDnsIPv6 performs a DNS lookup for the given domain and returns a slice of IPv6 addresses.
func GetDnsIPv6(ctx context.Context, domain string) ([]string, error) {
	return defaultClient.GetDnsIPv6(ctx, domain)
}

// GetCNAME performs a CNAME lookup for the given domain.
func GetCNAME(ctx context.Context, domain string) (string, error) {
	return defaultClient.GetCNAME(ctx, domain)
}

// GetTxt retrieves the TXT records for a given domain.
func GetTxt(ctx context.Context, domain string) ([]string, error) {
	return defaultClient.GetTxt(ctx, domain)
}

// GetMX retrieves the MX records for a given domain.
func GetMX(ctx context.Context, domain string) ([]*net.MX, error) {
	return defaultClient.GetMX(ctx, domain)
}

// GetNS retrieves the NS records for a given domain.
func GetNS(ctx context.Context, domain string) ([]*net.NS, error) {
	return defaultClient.GetNS(ctx, domain)
}

// GetSRV retrieves the SRV records for a given service, proto, and name.
func GetSRV(ctx context.Context, service, proto, name string) (string, []*net.SRV, error) {
	return defaultClient.GetSRV(ctx, service, proto, name)
}

// LookupAddr performs a reverse DNS lookup for the given IP address.
func LookupAddr(ctx context.Context, addr string) ([]string, error) {
	return defaultClient.LookupAddr(ctx, addr)
}

// GetMulti performs DNS lookups for multiple domains concurrently.
func GetMulti(ctx context.Context, domains []string) (map[string][]string, []error) {
	return defaultClient.GetMulti(ctx, domains)
}

func parseDomain(url string) string {
	s := strings.Split(url, "://")
	if len(s) == 2 {
		return strings.SplitN(s[1], "/", 2)[0]
	}
	return strings.SplitN(url, "/", 2)[0]
}