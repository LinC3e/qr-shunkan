package qr

import (
	"context"
	"errors"
	"net/url"
	"strings"

	"github.com/google/uuid"
)

var (
	ErrEmptyContent   = errors.New("the content cannot be empty")
	ErrInvalidURL     = errors.New("valid URL: (ej: https://ejemplo.com)")
	ErrNoScheme       = errors.New("URL protocol (http o https)")
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
			// https:// auto
			testWithScheme := "https://" + content
			parsedURL, err = url.ParseRequestURI(testWithScheme)
			if err == nil {
				content = testWithScheme
			}
		}

		if err != nil {
			return nil, ErrInvalidURL
		}
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return nil, ErrNoScheme
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