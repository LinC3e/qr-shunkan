package utils

/* Device & Browser */

import "strings"

func ParseUserAgent(ua string) (device string, browser string) {

	device = "desktop"

	if strings.Contains(strings.ToLower(ua), "mobile") {
		device = "mobile"
	}

	switch {
	case strings.Contains(ua, "Chrome"):
		browser = "Chrome"
	case strings.Contains(ua, "Firefox"):
		browser = "Firefox"
	case strings.Contains(ua, "Safari"):
		browser = "Safari"
	default:
		browser = "Other"
	}

	return
}
