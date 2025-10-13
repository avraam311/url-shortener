package url

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/avraam311/url-shortener/internal/models/domain"
	"github.com/avraam311/url-shortener/internal/models/dto"
)

const urlLength = 13

func (s *ServiceURL) CreateShortURL(ctx context.Context, fullURL *dto.FullURL) (string, error) {
	var shortURL string
	// to avoid collisions
	uniqueShortURL := false
	for !uniqueShortURL {
		shortURL = generateShortURL(fullURL.URL)
		unique, err := s.repo.CheckIfShortURLIsUnique(ctx, shortURL)
		if err != nil {
			return "", fmt.Errorf("service/create_short_url.go - %w", err)
		}
		uniqueShortURL = unique
	}

	url := domain.URL{
		FullURL:  fullURL.URL,
		ShortURL: shortURL,
	}

	err := s.repo.SaveShortURL(ctx, &url)
	if err != nil {
		return "", fmt.Errorf("service/create_short_url.go - %w", err)
	}

	return shortURL, nil
}

func generateShortURL(fullURL string) string {
	hash := sha256.Sum256([]byte(fullURL))
	encoded := base64.URLEncoding.EncodeToString(hash[:])
	shortURL := encoded[:urlLength]

	return shortURL
}
