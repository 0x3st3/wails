package application

import "testing"

func TestUriMatchesProxy(t *testing.T) {
	const (
		host = "proxy.example.com"
		port = 8080
	)
	tests := []struct {
		name string
		uri  string
		want bool
	}{
		// The proxy's own challenge, in the forms WebView2 may serialize it.
		{"absolute uri", "http://proxy.example.com:8080", true},
		{"trailing slash", "http://proxy.example.com:8080/", true},
		{"https scheme", "https://proxy.example.com:8080", true},
		{"uppercase host", "HTTP://PROXY.EXAMPLE.COM:8080", true},
		{"trailing dot host", "http://proxy.example.com.:8080", true},
		{"authority only", "proxy.example.com:8080", true},
		{"scheme relative", "//proxy.example.com:8080", true},
		{"with userinfo", "http://user:pass@proxy.example.com:8080", true},

		// Anything that is not the proxy must be refused.
		{"different host same port", "http://origin.example.com:8080", false},
		{"same host different port", "http://proxy.example.com:9090", false},
		{"default port vs 8080", "http://proxy.example.com", false},
		{"proxy host as subdomain", "http://proxy.example.com.evil.test:8080", false},
		{"proxy host as prefix", "http://proxy.example.com.evil:8080", false},
		{"proxy name in path", "http://evil.test:8080/proxy.example.com", false},
		{"proxy name in userinfo", "http://proxy.example.com@evil.test:8080", false},
		{"proxy name in query", "http://evil.test:8080/?h=proxy.example.com", false},
		{"proxy name in fragment", "http://evil.test:8080/#proxy.example.com", false},
		{"empty", "", false},
		{"garbage", "not a uri", false},
		{"scheme only", "http://", false},
		{"bare host no port", "proxy.example.com", false},
		{"mailto", "mailto:someone@proxy.example.com", false},
		{"file scheme", "file:///proxy.example.com", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uriMatchesProxy(tt.uri, host, port); got != tt.want {
				t.Fatalf("uriMatchesProxy(%q, %q, %d) = %v, want %v", tt.uri, host, port, got, tt.want)
			}
		})
	}
}

func TestUriMatchesProxyDefaultPorts(t *testing.T) {
	if !uriMatchesProxy("http://proxy.example.com", "proxy.example.com", 80) {
		t.Fatal("an http proxy on port 80 must match a uri that omits the default port")
	}
	if !uriMatchesProxy("https://proxy.example.com", "proxy.example.com", 443) {
		t.Fatal("an https proxy on port 443 must match a uri that omits the default port")
	}
	if uriMatchesProxy("https://proxy.example.com", "proxy.example.com", 80) {
		t.Fatal("an https uri must not resolve to the http default port")
	}
}

func TestUriMatchesProxyIPv6(t *testing.T) {
	if !uriMatchesProxy("http://[2001:db8::1]:8080", "2001:db8::1", 8080) {
		t.Fatal("an ipv6 proxy must match")
	}
	if uriMatchesProxy("http://[2001:db8::2]:8080", "2001:db8::1", 8080) {
		t.Fatal("a different ipv6 address must not match")
	}
}

func TestUriMatchesProxyRejectsEmptyProxyHost(t *testing.T) {
	if uriMatchesProxy("http://proxy.example.com:8080", "", 8080) {
		t.Fatal("an unconfigured proxy host must never match")
	}
}
