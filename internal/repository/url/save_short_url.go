package url

import (
	"context"
	"fmt"

	"github.com/avraam311/url-shortener/internal/models/domain"
)

func (r *RepositoryURL) SaveShortURL(ctx context.Context, url *domain.URL) error {
	query := `
		INSERT INTO url (
			full_url, short_url
		) VALUES ($1, $2);
	`

	_, err := r.db.ExecContext(ctx, query, url.FullURL, url.ShortURL)
	if err != nil {
		return fmt.Errorf("repository/save_short_url.go - failed to save short url - %w", err)
	}

	return nil
}
