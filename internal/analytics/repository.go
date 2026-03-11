package analytics

import (
	"context"
	"database/sql"
)

type Repository interface {
	CreateScan(ctx context.Context, qrID string, ip string, ua string) error
}

type postgresRepo struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &postgresRepo{db: db}
}

func (r *postgresRepo) CreateScan(ctx context.Context, qrID string, ip string, ua string) error {

	query := `
		INSERT INTO qr_scans (qr_id, ip_address, user_agent)
		VALUES ($1,$2,$3)
	`

	_, err := r.db.ExecContext(ctx, query, qrID, ip, ua)
	return err
}
