# LibWhois

[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)
[![GoDoc](https://pkg.go.dev/badge/github.com/iryndin/libwhois.svg)](https://pkg.go.dev/github.com/iryndin/libwhois)
[![Go Report Card](https://goreportcard.com/badge/github.com/iryndin/libwhois)](https://goreportcard.com/report/github.com/iryndin/libwhois)

## 1. Overview

Golang utilities to work with WHOIS: fetch whois info, parse it

## 2. Installation

```shell
go get -u github.com/iryndin/libwhois
```

## 3. Examples

### 3.1. Get whois host for zone

```go
    whoisHost, exists := libwhois.GetZoneWhoisHost("com")
    if exists {
        fmt.Printf("%s: %s\n", zone, whois)
    } else {
        fmt.Printf("%s: not found\n", zone)
    }
```

### 3.2. Get whois hosts for all zones

```go
    whoisHostMap := libwhois.GetWhoisHosts()
    allZones := make([]string, 0, len(whoisHostMap))
    for z, w := range whoisHostMap {
        if len(w) > 0 {
            allZones = append(allZones, z)
        }
    }
    sort.Strings(allZones)
    for _, z := range allZones {
        fmt.Printf("%s: %s\n", z, whoisHostMap[z])
    } 
```

### 3.3. Get whois for a domain 

```go
    import (
        "log"
        "github.com/iryndin/libwhois"
    )

    ...
 
    whoisClient := libwhois.NewSimpleWhoisClient()
    whois, err := whoisClient.Request("whois.verisign-grs.com", "chatgpt.com")

    if err != nil {
        log.Fatalf("Error fetching whois response: %v", err)
    }

    print(whois)
```

### 3.4. Get whois via proxy

```go
    import (
        "log"
        "github.com/iryndin/libwhois"
    )

    prx, err := libwhois.CreateSocks5Proxy("38.112.217.17", 5868, "user22", "pass1234")
    if err != nil {
        log.Fatalf("Failed to create proxy: %v", err)
    }

    whoisClient := libwhois.NewProxiedWhoisClient()
    whois, err := whoisClient.Request(prx, "whois.nic.io", "github.io")

    if err != nil {
        log.Fatalf("Error fetching whois response: %v", err)
    }

    print(whois)
```