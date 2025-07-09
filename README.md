# LibWhois

[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)
[![GoDoc](https://pkg.go.dev/badge/github.com/iryndin/libwhois.svg)](https://pkg.go.dev/github.com/iryndin/libwhois)
[![Go Report Card](https://goreportcard.com/badge/github.com/iryndin/libwhois)](https://goreportcard.com/report/github.com/iryndin/libwhois)

## 1. Overview

Golang utilities to work with WHOIS: fetch whois info, parse it

## 2. Examples

### 2.1. Get whois host for zone

```go
    whoisHost, exists := libwhois.GetZoneWhoisHost("com")
    if exists {
        fmt.Printf("%s: %s\n", zone, whois)
    } else {
        fmt.Printf("%s: not found\n", zone)
    }
```

### 2.2. Get whois hosts for all zones

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


