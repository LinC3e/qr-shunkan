package qr

import (
	"context"
	"errors"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

var (
	ErrEmptyContent = errors.New("the content cannot be empty")
	ErrInvalidURL   = errors.New("valid URL required (ej: https://example.com)")
	ErrNoScheme     = errors.New("URL must use http or https")
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateQR(ctx context.Context, content string) (*QR, error) {

	content = strings.TrimSpace(content)

	if content == "" {
		return nil, ErrEmptyContent
	}

	parsedURL, err := url.ParseRequestURI(content)
	if err != nil {

		if strings.HasPrefix(content, "www.") || strings.Contains(content, ".") {
			test := "https://" + content

			parsedURL, err = url.ParseRequestURI(test)
			if err == nil {
				content = test
			}
		}

		if err != nil {
			return nil, ErrInvalidURL
		}
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return nil, ErrNoScheme
	}

	slug, err := gonanoid.New(6)
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()

	qr := &QR{
		ID:        uuid.New(),
		Slug:      slug,
		Content:   content,
		Type:      "url",
		IsActive:  true,
		ScanCount: 0,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err = s.repo.Create(ctx, qr)
	if err != nil {
		return nil, err
	}

	return qr, nil
}

func (s *Service) GetQR(ctx context.Context, id string) (*QR, error) {

	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid ID")
	}

	return s.repo.GetByID(ctx, uid)
}

func (s *Service) ResolveQR(ctx context.Context, slug string) (*QR, error) {

	qr, err := s.repo.GetBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}

	if !qr.IsActive {
		return nil, errors.New("qr inactive")
	}

	if qr.MaxScans != nil && qr.ScanCount >= *qr.MaxScans {
		return nil, errors.New("scan limit reached")
	}

	if qr.ExpiresAt != nil && time.Now().After(*qr.ExpiresAt) {
		return nil, errors.New("qr expired")
	}

	err = s.repo.IncrementScan(ctx, qr.ID.String())
	if err != nil {
		return nil, err
	}

	return qr, nil
}