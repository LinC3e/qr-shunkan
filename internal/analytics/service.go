package analytics

import "context"

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) RegisterScan(ctx context.Context, qrID string, ip string, ua string) error {
	return s.repo.CreateScan(ctx, qrID, ip, ua)
}
