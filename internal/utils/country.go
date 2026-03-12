package utils

import (
	"log"
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

	parsedIP, err := netip.ParseAddr(ip)
	if err != nil {
		log.Fatal(err)
	}

	record, err := geoDB.Country(parsedIP)
	if err != nil {
		return "Unknown"
	}

	return record.Country.Names.English
}
