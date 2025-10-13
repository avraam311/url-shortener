package url

import (
	"context"

	"github.com/avraam311/url-shortener/internal/models/domain"
)

type RepositoryURL interface {
	SaveShortURL(context.Context, *domain.URL) error
	GetFullURL(context.Context, string) (string, error)
	CheckIfShortURLIsUnique(context.Context, string) (bool, error)
}

type ServiceURL struct {
	repo RepositoryURL
}

func NewService(repo RepositoryURL) *ServiceURL {
	return &ServiceURL{
		repo: repo,
	}
}
