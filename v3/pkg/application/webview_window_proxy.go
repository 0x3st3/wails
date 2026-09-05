package application

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

func parseWebviewProxyServer(server string) (host string, port int, scheme string, err error) {
	u, err := url.Parse(server)
	if err != nil {
		return "", 0, "", err
	}
	scheme = u.Scheme
	if scheme != "http" && scheme != "https" {
		return "", 0, "", fmt.Errorf("unsupported proxy scheme %q (want http or https)", scheme)
	}
	host = u.Hostname()
	portStr := u.Port()
	if host == "" || portStr == "" {
		return "", 0, "", fmt.Errorf("proxy URL must include host and port")
	}
	port, err = strconv.Atoi(portStr)
	if err != nil {
		return "", 0, "", fmt.Errorf("invalid proxy port %q: %w", portStr, err)
	}
	return host, port, scheme, nil
}

// webviewProxyServerArg returns the credential-stripped "scheme://host:port"
// string for use as a Chromium --proxy-server argument.
func webviewProxyServerArg(server string) (string, error) {
	host, port, scheme, err := parseWebviewProxyServer(server)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s://%s:%d", scheme, host, port), nil
}

// canonicalHost lowercases a host and drops a single trailing dot so that
// "Proxy.Example.com." and "proxy.example.com" compare equal.
func canonicalHost(host string) string {
	return strings.ToLower(strings.TrimSuffix(host, "."))
}

// uriAuthority extracts the host and port a URI addresses, resolving the port
// from the scheme when it is left implicit. It accepts an authority-only value
// ("host:port") as well as an absolute URI.
func uriAuthority(uri string) (host string, port string, ok bool) {
	u, err := url.Parse(uri)
	if err != nil || u.Host == "" {
		// WebView2 is expected to supply an absolute URI. Fall back to an
		// authority-only value rather than refuse the challenge outright.
		u, err = url.Parse("//" + strings.TrimPrefix(uri, "//"))
		if err != nil || u.Host == "" {
			return "", "", false
		}
	}
	host, port = u.Hostname(), u.Port()
	if host == "" {
		return "", "", false
	}
	if port == "" {
		switch strings.ToLower(u.Scheme) {
		case "http":
			port = "80"
		case "https":
			port = "443"
		default:
			return "", "", false
		}
	}
	return host, port, true
}

// uriMatchesProxy reports whether uri addresses the given proxy host and port.
// WebView2 raises BasicAuthenticationRequested for origin and proxy challenges
// alike and provides no auth-type field; for proxy challenges the event's Uri is
// the proxy server's own URI, so it is the only reliable discriminator.
func uriMatchesProxy(uri, proxyHost string, proxyPort int) bool {
	if uri == "" || proxyHost == "" {
		return false
	}
	host, port, ok := uriAuthority(uri)
	if !ok {
		return false
	}
	if canonicalHost(host) != canonicalHost(proxyHost) {
		return false
	}
	return port == strconv.Itoa(proxyPort)
}
