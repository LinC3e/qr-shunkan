package qr

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, qr *QR) error
	GetByID(ctx context.Context, id uuid.UUID) (*QR, error)
}

type postgresRepo struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &postgresRepo{db: db}
}

func (r *postgresRepo) Create(ctx context.Context, qr *QR) error {

	query := `
	INSERT INTO qrs (
		id, slug, content, type,
		scan_count, is_active,
		created_at, updated_at
	)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		qr.ID,
		qr.Slug,
		qr.Content,
		qr.Type,
		qr.ScanCount,
		qr.IsActive,
		qr.CreatedAt,
		qr.UpdatedAt,
	)

	return err
}

func (r *postgresRepo) GetByID(ctx context.Context, id uuid.UUID) (*QR, error) {

	query := `
	SELECT
		id, slug, content, type,
		scan_count, is_active,
		created_at, updated_at
	FROM qrs
	WHERE id = $1
	`

	qr := &QR{}

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&qr.ID,
		&qr.Slug,
		&qr.Content,
		&qr.Type,
		&qr.ScanCount,
		&qr.IsActive,
		&qr.CreatedAt,
		&qr.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return qr, nil
}