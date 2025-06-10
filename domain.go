package domain

import (
	"context"
	"net"
	"strings"
)

func getResolver(ctx context.Context, dns ...string) *net.Resolver {
	if len(dns) > 0 && dns[0] != "" {
		return &net.Resolver{
			PreferGo: true,
			Dial: func(_ context.Context, network, address string) (net.Conn, error) {
				d := net.Dialer{}
				dnsServer := dns[0]
				if !strings.Contains(dnsServer, ":") {
					dnsServer += ":53"
				}
				return d.DialContext(ctx, network, dnsServer)
			},
		}
	}
	return net.DefaultResolver
}
