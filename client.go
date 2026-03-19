package domain

import (
	"context"
	"net"
	"strings"
)

type resolver interface {
	LookupAddr(ctx context.Context, addr string) ([]string, error)
	LookupCNAME(ctx context.Context, host string) (string, error)
	LookupIP(ctx context.Context, network, host string) ([]net.IP, error)
	LookupMX(ctx context.Context, name string) ([]*net.MX, error)
	LookupNS(ctx context.Context, name string) ([]*net.NS, error)
	LookupSRV(ctx context.Context, service, proto, name string) (string, []*net.SRV, error)
	LookupTXT(ctx context.Context, name string) ([]string, error)
}

// Client is a DNS client for performing lookups.
type Client struct {
	resolver resolver
}

// NewClient creates a new DNS client.
// If a dns server address is provided, it will be used for lookups.
// Otherwise, the system's default resolver will be used.
func NewClient(dns ...string) *Client {
	var r *net.Resolver
	if len(dns) > 0 && dns[0] != "" {
		dnsServer := dns[0]
		r = &net.Resolver{
			PreferGo: true,
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				d := net.Dialer{}
				if !strings.Contains(dnsServer, ":") {
					dnsServer += ":53"
				}
				return d.DialContext(ctx, network, dnsServer)
			},
		}
	} else {
		r = net.DefaultResolver
	}

	return &Client{
		resolver: r,
	}
}
