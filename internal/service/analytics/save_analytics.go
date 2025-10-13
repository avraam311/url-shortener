package analytics

import (
	"context"
	"fmt"

	"github.com/avraam311/url-shortener/internal/models/db"
)

func (s *ServiceAnalytics) SaveAnalytics(ctx context.Context, analytics *db.Analytics) error {
	if err := s.repo.SaveAnalytics(ctx, analytics); err != nil {
		return fmt.Errorf("service/save_analytics.go - %w", err)
	}

	return nil
}
