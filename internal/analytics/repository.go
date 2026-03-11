package analytics

import (
	"context"
	"database/sql"
)

type Repository interface {
	GetStats(ctx context.Context, qrID string) (*Stats, error)
	GetDailyStats(ctx context.Context, qrID string) ([]DailyStats, error)
	CreateScan(ctx context.Context, qrID string, ip string, ua string) error
}

type postgresRepo struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &postgresRepo{db: db}
}

func (r *postgresRepo) GetStats(ctx context.Context, qrID string) (*Stats, error) {

	query := `
	SELECT
		COUNT(*) AS total_scans,
		COUNT(*) FILTER (
			WHERE scanned_at >= NOW() - INTERVAL '1 day'
		) AS today_scans
	FROM qr_scans
	WHERE qr_id = $1
	`

	stats := &Stats{}

	err := r.db.QueryRowContext(ctx, query, qrID).Scan(
		&stats.TotalScans,
		&stats.TodayScans,
	)

	if err != nil {
		return nil, err
	}

	return stats, nil
}

func (r *postgresRepo) GetDailyStats(ctx context.Context, qrID string) ([]DailyStats, error) {

	query := `
	SELECT
		DATE(scanned_at) as date,
		COUNT(*) as scans
	FROM qr_scans
	WHERE qr_id = $1
	GROUP BY DATE(scanned_at)
	ORDER BY date ASC
	`

	rows, err := r.db.QueryContext(ctx, query, qrID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stats []DailyStats

	for rows.Next() {

		var s DailyStats

		err := rows.Scan(
			&s.Date,
			&s.Scans,
		)

		if err != nil {
			return nil, err
		}

		stats = append(stats, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return stats, nil
}

func (r *postgresRepo) CreateScan(ctx context.Context, qrID string, ip string, ua string) error {

	query := `
		INSERT INTO qr_scans (qr_id, ip_address, user_agent)
		VALUES ($1,$2,$3)
	`

	_, err := r.db.ExecContext(ctx, query, qrID, ip, ua)
	return err
}
