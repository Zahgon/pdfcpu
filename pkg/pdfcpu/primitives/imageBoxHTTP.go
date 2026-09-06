package primitives

import (
	"context"
	"net"
	"net/http"
	"net/url"
	"time"
)

func imageBoxRemoteURL(s string) (*url.URL, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func validateImageBoxRemoteURL(u *url.URL) error { _ = "STUB: not implemented"; return nil }

func rejectPrivateImageBoxHost(host string) error { _ = "STUB: not implemented"; return nil }

func rejectPrivateImageBoxIP(host string, ip net.IP) error { _ = "STUB: not implemented"; return nil }

func imageBoxBlockedIP(ip net.IP) bool { _ = "STUB: not implemented"; return false }

func (pdf *PDF) imageBoxHTTPClient() *http.Client { _ = "STUB: not implemented"; return nil }

func imageBoxTransport(timeout time.Duration) *http.Transport {
	_ = "STUB: not implemented"
	return nil
}

func imageBoxRedirect(req *http.Request, via []*http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func imageBoxDialContext(dialer *net.Dialer) func(context.Context, string, string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return nil
}

func rejectImageBoxIPs(host string, ips []net.IPAddr) error { _ = "STUB: not implemented"; return nil }
