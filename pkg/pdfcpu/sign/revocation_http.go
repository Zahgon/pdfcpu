package sign

import (
	"context"
	"net"
	"net/http"
	"net/url"
	"time"
)

const (
	defaultRevocationHTTPTimeout = 10 * time.Second
	maxRevocationRedirects       = 10
)

type revocationResolver interface {
	LookupIPAddr(context.Context, string) ([]net.IPAddr, error)
}

type revocationDialer func(context.Context, string, string) (net.Conn, error)

func normalizeRevocationHost(host string) string { _ = "STUB: not implemented"; return "" }

func allowedRevocationHostSet(hosts []string) map[string]bool {
	_ = "STUB: not implemented"
	return nil
}

func validateRevocationURL(u *url.URL) error { _ = "STUB: not implemented"; return nil }

func validateRevocationURLString(s string) error { _ = "STUB: not implemented"; return nil }

func revocationBlockedIP(ip net.IP) bool { _ = "STUB: not implemented"; return false }

func validateRevocationIPs(host string, ips []net.IPAddr, allowed map[string]bool) error {
	_ = "STUB: not implemented"
	return nil
}

func revocationDialContext(
	resolver revocationResolver,
	dial revocationDialer,
	allowed map[string]bool,
) func(context.Context, string, string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return nil
}

func revocationRedirect(req *http.Request, via []*http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func revocationHTTPClient(timeout time.Duration, allowedHosts []string) *http.Client {
	_ = "STUB: not implemented"
	return nil
}
