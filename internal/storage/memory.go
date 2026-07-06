package storage

import (
	"context"
	"sort"
	"sync"

	"github.com/stormbels/linkshelf/internal/domain"
)

type MemoryStorage struct {
	mu     sync.RWMutex
	nextID int64
	links  []domain.Link
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		nextID: 1,
		links:  make([]domain.Link, 0),
	}
}

func (s *MemoryStorage) Save(ctx context.Context, link domain.Link) error {
	_ = ctx

	s.mu.Lock()
	defer s.mu.Unlock()

	link.ID = s.nextID
	s.nextID++

	s.links = append(s.links, link)
	return nil
}
func (s *MemoryStorage) ExistsByURL(ctx context.Context, url string) (bool, error) {
	_ = ctx

	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, link := range s.links {
		if link.URL == url {
			return true, nil
		}
	}

	return false, nil
}

func (s *MemoryStorage) Update(ctx context.Context, id int64, updatedLink domain.Link) error {
	_ = ctx

	s.mu.Lock()
	defer s.mu.Unlock()

	for i, link := range s.links {
		if link.ID == id {
			updatedLink.ID = link.ID
			updatedLink.CreatedAt = link.CreatedAt
			s.links[i] = updatedLink
			return nil
		}
	}

	return nil
}

func (s *MemoryStorage) List(ctx context.Context) ([]domain.Link, error) {
	_ = ctx

	s.mu.RLock()
	defer s.mu.RUnlock()

	links := make([]domain.Link, len(s.links))
	copy(links, s.links)

	sort.SliceStable(links, func(i, j int) bool {
		if links[i].CreatedAt.Equal(links[j].CreatedAt) {
			return links[i].ID > links[j].ID
		}

		return links[i].CreatedAt.After(links[j].CreatedAt)
	})

	return links, nil
}

func (s *MemoryStorage) Delete(ctx context.Context, id int64) error {
	_ = ctx

	s.mu.Lock()
	defer s.mu.Unlock()

	for i, link := range s.links {
		if link.ID == id {
			s.links = append(s.links[:i], s.links[i+1:]...)
			return nil
		}
	}

	return nil
}
