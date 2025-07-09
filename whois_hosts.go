package libwhois

import (
	_ "embed"
	"encoding/json"
	"log"
)

type whoisData struct {
	Zone  string `json:"zone" `
	Whois string `json:"whois,omitempty"`
}

//go:embed whois.json
var whoisHostsFileContent string

var whoisHostsMap map[string]string

// GetZoneWhoisHost get whois host for a zone
// It returns whois hostname and bool flag which is true if there exist whois host for given zone, or false otherwise.
func GetZoneWhoisHost(zone string) (string, bool) {
	m := GetWhoisHosts()
	result, ok := m[zone]
	return result, ok
}

// GetWhoisHosts return a map where keys are zones and values are whois hosts
func GetWhoisHosts() map[string]string {
	if whoisHostsMap == nil {
		loadWhoisHosts()
	}
	return whoisHostsMap
}

func loadWhoisHosts() {
	var data []whoisData
	if err := json.Unmarshal([]byte(whoisHostsFileContent), &data); err != nil {
		log.Fatal(err)
	}
	whoisHostsMap = make(map[string]string, len(data)*2)
	for _, d := range data {
		if len(d.Whois) > 0 {
			whoisHostsMap[d.Zone] = d.Whois
		}
	}
}
