package postgres

import (
	"context"
	"fmt"

	"github.com/stormbels/linkshelf/internal/domain"
)

func (s *Storage) Save(ctx context.Context, link domain.Link) error {
	query := `
		INSERT INTO links (url, title, description, tags, created_at)
		VALUES ($1, $2, $3, $4, $5)
	`

	if _, err := s.pool.Exec(
		ctx,
		query,
		link.URL,
		link.Title,
		link.Description,
		link.Tags,
		link.CreatedAt,
	); err != nil {
		return fmt.Errorf("save link: %w", err)
	}

	return nil
}

func (s *Storage) Update(ctx context.Context, id int64, link domain.Link) error {
	query := `
		UPDATE links
		SET url = $1,
			title = $2,
			description = $3,
			tags = $4
		WHERE id = $5
	`

	if _, err := s.pool.Exec(
		ctx,
		query,
		link.URL,
		link.Title,
		link.Description,
		link.Tags,
		id,
	); err != nil {
		return fmt.Errorf("update link: %w", err)
	}

	return nil
}

func (s *Storage) Delete(ctx context.Context, id int64) error {
	query := `
		DELETE FROM links
		WHERE id = $1
	`

	if _, err := s.pool.Exec(ctx, query, id); err != nil {
		return fmt.Errorf("delete link: %w", err)
	}

	return nil
}
