package analytics

import "time"

type Scan struct {
	ID        int64
	QRID      string
	IPAddress string
	UserAgent string
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
