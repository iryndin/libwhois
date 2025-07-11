package libwhois

import (
	"fmt"
	"time"

	"golang.org/x/net/proxy"
)

const DefaultWhoisPort = 43

// SimpleWhoisClient whois client allows to submit basic WHOIS calls
type SimpleWhoisClient interface {
	Request(host string, query string) (string, error)
	Request2(host string, port int, query string) (string, error)
	RequestWithTimeout(host string, port int, query string, timeout time.Duration) (string, error)
}

// ProxiedWhoisClient whois client with proxy
type ProxiedWhoisClient interface {
	Request(prx proxy.Dialer, host string, query string) (string, error)
	Request2(prx proxy.Dialer, host string, port int, query string) (string, error)
	RequestWithTimeout(prx proxy.Dialer, host string, port int, query string, timeout time.Duration) (string, error)
}

// NewSimpleWhoisClient create new SimpleWhoisClient
func NewSimpleWhoisClient() SimpleWhoisClient {
	return &simpleWhoisClientImpl{}
}

// NewProxiedWhoisClient create new ProxiedWhoisClient
func NewProxiedWhoisClient() ProxiedWhoisClient {
	return &proxiedWhoisClientImpl{}
}

// CreateSocks5Proxy return a SOCKS5 proxy, accept host and port, and username and password
func CreateSocks5Proxy(host string, port int, user string, password string) (proxy.Dialer, error) {
	hostPortString := fmt.Sprintf("%s:%d", host, port)

	auth := &proxy.Auth{
		User:     user,
		Password: password,
	}

	// Create a SOCKS5 dialer
	return proxy.SOCKS5("tcp", hostPortString, auth, proxy.Direct)
}
