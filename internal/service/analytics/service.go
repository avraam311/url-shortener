package analytics

import (
	"context"

	"github.com/avraam311/url-shortener/internal/models/db"
)

type RepositoryAnalytics interface {
	GetAnalytics(context.Context, string) ([]*db.Analytics, error)
	SaveAnalytics(context.Context, *db.Analytics) error
}

type ServiceAnalytics struct {
	repo RepositoryAnalytics
}

func NewService(repo RepositoryAnalytics) *ServiceAnalytics {
	return &ServiceAnalytics{
		repo: repo,
	}
}
