package qr

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

type QR struct {
	ID        uuid.UUID `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type Repository interface {
	Create(ctx context.Context, content string) (*QR, error)
	GetByID(ctx context.Context, id uuid.UUID) (*QR, error)
}

type postgresRepo struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &postgresRepo{db: db}
}

func (r *postgresRepo) Create(ctx context.Context, content string) (*QR, error) {
	qr := &QR{
		ID:        uuid.New(),
		Content:   content,
		CreatedAt: time.Now().UTC(),
	}

	query := `
		INSERT INTO qrs (id, content, created_at)
		VALUES ($1, $2, $3)
		RETURNING id, content, created_at`

	err := r.db.QueryRowContext(ctx, query, qr.ID, qr.Content, qr.CreatedAt).
		Scan(&qr.ID, &qr.Content, &qr.CreatedAt)
	if err != nil {
		return nil, err
	}

	return qr, nil
}

func (r *postgresRepo) GetByID(ctx context.Context, id uuid.UUID) (*QR, error) {
	qr := &QR{}

	query := `SELECT id, content, created_at FROM qrs WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(&qr.ID, &qr.Content, &qr.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("qr not found.")
	}
	if err != nil {
		return nil, err
	}

	return qr, nil
}