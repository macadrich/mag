package discovery

import (
	"fmt"
	"net"
	"strings"
)

// ServiceRecord contains the basic description of a service,
// contains instance name, service type & domain
type ServiceRecord struct {
	Instance            string `json:"name"`
	Service             string `json:"type"`
	Domain              string `json:"domain"`
	serviceName         string
	serviceInstanceName string
	serviceTypeName     string
}

// ServiceName returns complete service name
func (s *ServiceRecord) ServiceName() string {
	if s.serviceName == "" {
		s.serviceName = fmt.Sprintf("%s.%s.", trimDot(s.Service), trimDot(s.Domain))
	}
	return s.serviceName
}

// ServiceInstanceName returns complete service instance name
func (s *ServiceRecord) ServiceInstanceName() string {
	if s.Instance == "" {
		return ""
	}

	if s.serviceInstanceName == "" {
		s.serviceInstanceName = fmt.Sprintf("%s.%s", trimDot(s.Instance), s.ServiceName())
	}
	return s.serviceInstanceName
}

// ServiceTypeName -
func (s *ServiceRecord) ServiceTypeName() string {
	if s.serviceTypeName == "" {
		domain := "local"
		if len(s.Domain) > 0 {
			domain = trimDot(s.Domain)
		}
		s.serviceTypeName = fmt.Sprintf("_services._dns-sd._udp.%s.", domain)
	}
	return s.serviceTypeName
}

// NewServiceRecord constructs a ServiceRecord structure by given arguments
func NewServiceRecord(instance, service, domain string) *ServiceRecord {
	return &ServiceRecord{instance, service, domain, "", "", ""}
}

// LookupParams contains configurable properties to create a service discovery request
type LookupParams struct {
	ServiceRecord
	Entries chan<- *ServiceEntry
}

// NewLookupParams constructs a LookupParams structure by given arguments
func NewLookupParams(instance, service, domain string, entries chan<- *ServiceEntry) *LookupParams {
	return &LookupParams{
		*NewServiceRecord(instance, service, domain),
		entries,
	}
}

// ServiceEntry represents a browse/lookup result for client API
type ServiceEntry struct {
	ServiceRecord
	HostName string   `json:"hostname"`
	Port     int      `json:"port"`
	Text     []string `json:"text"`
	TTL      uint32   `json:"ttl"`
	AddrIPv4 net.IP   `json:"-"`
	AddrIPv6 net.IP   `json:"-"`
}

// NewServiceEntry constructs a ServiceEntry structure by given arguments
func NewServiceEntry(instance, service, domain string) *ServiceEntry {
	return &ServiceEntry{
		*NewServiceRecord(instance, service, domain),
		"",
		0,
		[]string{},
		0,
		nil,
		nil,
	}
}

// trimDot is used to trim the dots from the start or end of a string
func trimDot(s string) string {
	return strings.Trim(s, ".")
}
