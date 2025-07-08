package libwhois

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

type simpleWhoisClientImpl struct {
}

func (c simpleWhoisClientImpl) Request(host string, query string) (string, error) {
	return c.Request2(host, DefaultWhoisPort, query)
}

func (simpleWhoisClientImpl) Request2(host string, port int, query string) (string, error) {
	hostPortString := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.Dial("tcp", hostPortString)
	if err != nil {
		return "", fmt.Errorf("failed to connect to '%s': %w", hostPortString, err)
	}
	defer conn.Close()

	return readDataFromConn(conn, query)
}

func (simpleWhoisClientImpl) RequestWithTimeout(host string, port int, query string, timeout time.Duration) (string, error) {
	hostPortString := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", hostPortString, timeout)
	if err != nil {
		return "", fmt.Errorf("failed to connect to '%s': %w", hostPortString, err)
	}
	defer conn.Close()

	return readDataFromConn(conn, query)
}

func readDataFromConn(conn net.Conn, query string) (string, error) {
	// Send query with CRLF as per WHOIS protocol
	_, err := fmt.Fprintf(conn, "%s\r\n", query)
	if err != nil {
		return "", err
	}

	var strBuilder strings.Builder
	strBuilder.Grow(16 * 1024)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		strBuilder.WriteString(scanner.Text() + "\n")
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return strBuilder.String(), nil
}
