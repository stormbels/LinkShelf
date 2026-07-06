package storage

import (
	"context"

	"github.com/stormbels/linkshelf/internal/domain"
)

type LinkStorage interface {
	Save(ctx context.Context, link domain.Link) error

	List(ctx context.Context) ([]domain.Link, error)

	ExistsByURL(ctx context.Context, url string) (bool, error)

	Update(ctx context.Context, id int64, link domain.Link) error

	Delete(ctx context.Context, id int64) error
}
