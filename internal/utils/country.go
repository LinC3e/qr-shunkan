package utils

import (
	"net/netip"

	"github.com/oschwald/geoip2-golang/v2"
)

var geoDB *geoip2.Reader

func InitGeoDB(path string) error {
	var err error
	geoDB, err = geoip2.Open(path)
	return err
}

func GetCountry(ip string) string {

	if geoDB == nil {
		return "Unknown"
	}

	// for local
	if ip == "::1" || ip == "127.0.0.1" {
		ip = "8.8.8.8"
	}

	parsedIP, err := netip.ParseAddr(ip)
	if err != nil {
		return "Unknown"
	}

	record, err := geoDB.Country(parsedIP)
	if err != nil {
		return "Unknown"
	}

	return record.Country.Names.English
}
