package bsw

import (
	"strings"
)

// MX returns the A record for any MX records for a domain.
func MX(domain, serverAddr string) *Tsk {
	t := newTsk("mx")
	servers, err := LookupMX(domain, serverAddr)
	if err != nil {
		t.SetErr(err)
		return t
	}
	for _, s := range servers {
		ip, err := LookupName(s, serverAddr)
		if err != nil || ip == "" {
			continue
		}
		t.AddResult(ip, strings.TrimRight(s, "."))
	}
	return t
}
