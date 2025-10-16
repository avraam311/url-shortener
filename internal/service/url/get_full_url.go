package url

import (
	"context"
	"fmt"
)

func (s *ServiceURL) GetFullURL(ctx context.Context, shortURL string) (string, error) {
	fullURL, err := s.repo.GetFullURL(ctx, shortURL)
	if err != nil {
		return "", fmt.Errorf("service/get_full_url.go by %s - %w", shortURL, err)
	}
	if fullURL[:7] != "http://" {
		fullURL = "http://" + fullURL
	}

	return fullURL, nil
}
