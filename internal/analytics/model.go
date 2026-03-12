package analytics

import "time"

type Scan struct {
	ID        int64
	QRID      string
	IPAddress string
	UserAgent string
	Country   string
	Device    string
	Browser   string
	ScannedAt time.Time
}

type Stats struct {
	TotalScans int `json:"total_scans"`
	TodayScans int `json:"today_scans"`
}

type DailyStats struct {
	Date  string `json:"date"`
	Scans int    `json:"scans"`
}

type CountryStats struct {
	Country string `json:"country"`
	Scans   int    `json:"scans"`
}

type DeviceStats struct {
	Device string `json:"device"`
	Scans  int    `json:"scans"`
}

type BrowserStats struct {
	Browser string `json:"browser"`
	Scans   int    `json:"scans"`
}
