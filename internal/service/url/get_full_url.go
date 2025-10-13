package url

import (
	"context"
	"fmt"
)

func (s *ServiceURL) GetFullURL(ctx context.Context, shortURL string) (string, error) {
	fullURL, err := s.repo.GetFullURL(ctx, shortURL)
	if err != nil {
		return "", fmt.Errorf("service/get/full_url.go - %w", err)
	}

	return fullURL, nil
}
