package libwhois

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"golang.org/x/net/proxy"
)

type proxiedWhoisClientImpl struct {
}

func (c proxiedWhoisClientImpl) Request(prx proxy.Dialer, host string, query string) (string, error) {
	return c.Request2(prx, host, DefaultWhoisPort, query)
}

func (proxiedWhoisClientImpl) Request2(prx proxy.Dialer, host string, port int, query string) (string, error) {
	conn, err := prx.Dial("tcp", net.JoinHostPort(host, strconv.Itoa(port)))
	if err != nil {
		return "", fmt.Errorf("whois: connect to proxy or destination server failed: %w", err)
	}
	defer conn.Close()

	return readDataFromConn(conn, query)
}

func (proxiedWhoisClientImpl) RequestWithTimeout(prx proxy.Dialer, host string, port int, query string, timeout time.Duration) (string, error) {
	start := time.Now()
	conn, err := prx.Dial("tcp", net.JoinHostPort(host, strconv.Itoa(port)))
	if err != nil {
		return "", fmt.Errorf("whois: connect to proxy or destination server failed: %w", err)
	}
	defer conn.Close()
	elapsed := time.Since(start)

	_ = conn.SetWriteDeadline(time.Now().Add(timeout - elapsed))

	return readDataFromConn(conn, query)
}
