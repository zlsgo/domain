# Domain

## Usage

This package offers two ways to perform DNS lookups: simple package-level functions for convenience, and a client-based approach for more control.

### Simple Usage (Recommended for single lookups)

You can use the package-level functions for quick and easy lookups. These use a shared default client.

```go
package main

import (
	"context"
	"fmt"
	"github.com/zlsgo/domain"
)

func main() {
	// Get all IPs (IPv4 and IPv6)
	ips, err := domain.GetDns(context.Background(), "www.google.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("IPs:", ips)

	// Get only IPv4
	ipv4s, _ := domain.GetDnsIPv4(context.Background(), "www.google.com")
	fmt.Println("IPv4s:", ipv4s)

	// Get CNAME
	cname, _ := domain.GetCNAME(context.Background(), "www.github.com")
	fmt.Println("CNAME:", cname)
}
```

### Advanced Usage (Recommended for multiple lookups or custom DNS)

For scenarios requiring a custom DNS server or for performing many lookups, it's more efficient to create and reuse a client.

#### Single lookup

```go
import (
	"context"
	"fmt"
	"github.com/zlsgo/domain"
)

func main() {
	// Use system default resolver
	client := domain.NewClient()

	// Or specify a DNS server
	// client := domain.NewClient("8.8.8.8")

	// Get all IPs (IPv4 and IPv6)
	ips, err := client.GetDns(context.Background(), "www.google.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("IPs:", ips)

	// Get only IPv4
	ipv4s, _ := client.GetDnsIPv4(context.Background(), "www.google.com")
	fmt.Println("IPv4s:", ipv4s)

	// Get CNAME
	cname, _ := client.GetCNAME(context.Background(), "www.github.com")
	fmt.Println("CNAME:", cname)

	// Get TXT records
	txts, _ := client.GetTxt(context.Background(), "google.com")
	fmt.Println("TXTs:", txts)

	// Get MX records
	mxs, _ := client.GetMX(context.Background(), "google.com")
	for _, mx := range mxs {
		fmt.Printf("MX: %s %d\n", mx.Host, mx.Pref)
	}

	// Get NS records
	nss, _ := client.GetNS(context.Background(), "google.com")
	for _, ns := range nss {
		fmt.Printf("NS: %s\n", ns.Host)
	}

	// Get SRV records
	cname, srvs, _ := client.GetSRV(context.Background(), "sip", "tcp", "sip.voice.google.com")
	fmt.Println("SRV CNAME:", cname)
	for _, srv := range srvs {
		fmt.Printf("SRV: %s:%d %d %d\n", srv.Target, srv.Port, srv.Priority, srv.Weight)
	}

	// Reverse DNS lookup
	names, _ := client.LookupAddr(context.Background(), "8.8.8.8")
	fmt.Println("Names for 8.8.8.8:", names)
}
```

#### Multiple lookups

```go
	// ... client initialization ...

	domains := []string{"www.google.com", "www.github.com", "invalid-domain-for-test"}
	results, errs := client.GetMulti(context.Background(), domains)
	if len(errs) > 0 {
		fmt.Println("Errors:", errs)
	}

	for d, ips := range results {
		fmt.Printf("Domain: %s, IPs: %v\n", d, ips)
	}
```

