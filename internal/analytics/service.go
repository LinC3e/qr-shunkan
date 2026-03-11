package analytics

import "context"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetStats(ctx context.Context, qrID string) (*Stats, error) {
	return s.repo.GetStats(ctx, qrID)
}

func (s *Service) GetDailyStats(ctx context.Context, qrID string) ([]DailyStats, error) {
	return s.repo.GetDailyStats(ctx, qrID)
}

func (s *Service) RegisterScan(ctx context.Context, qrID string, ip string, ua string) error {
	return s.repo.CreateScan(ctx, qrID, ip, ua)
}
