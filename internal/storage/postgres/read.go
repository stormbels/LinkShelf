package postgres

import (
	"context"
	"fmt"

	"github.com/stormbels/linkshelf/internal/domain"
)

func (s *Storage) List(ctx context.Context) ([]domain.Link, error) {
	const query = `
		SELECT id, url, title, description, tags, created_at
		FROM links
		ORDER BY created_at DESC, id DESC
	`

	rows, err := s.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("list links: %w", err)
	}
	defer rows.Close()

	links := make([]domain.Link, 0)
	for rows.Next() {
		var link domain.Link
		if err := rows.Scan(
			&link.ID,
			&link.URL,
			&link.Title,
			&link.Description,
			&link.Tags,
			&link.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan link: %w", err)
		}

		links = append(links, link)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate links: %w", err)
	}

	return links, nil
}

func (s *Storage) ExistsByURL(ctx context.Context, rawURL string) (bool, error) {
	const query = `
		SELECT EXISTS(
			SELECT 1
			FROM links
			WHERE url = $1
		)
	`

	var exists bool
	if err := s.pool.QueryRow(ctx, query, rawURL).Scan(&exists); err != nil {
		return false, fmt.Errorf("check link exists by url: %w", err)
	}

	return exists, nil
}
