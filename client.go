package libwhois

import "time"

const DefaultWhoisPort = 43

type SimpleWhoisClient interface {
	Request(host string, query string) (string, error)
	Request2(host string, port int, query string) (string, error)
	RequestWithTimeout(host string, port int, query string, timeout time.Duration) (string, error)
}

func NewSimpleWhoisClient() SimpleWhoisClient {
	return &simpleWhoisClientImpl{}
}
