package qr

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateQR(ctx context.Context, content string) (*QR, error) {
	if content == "" {
		return nil, errors.New("content is required")
	}
	return s.repo.Create(ctx, content)
}

func (s *Service) GetQR(ctx context.Context, id string) (*QR, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid ID")
	}
	return s.repo.GetByID(ctx, uid)
}