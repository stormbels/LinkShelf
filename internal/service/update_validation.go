package service

import (
	"context"

	"github.com/stormbels/linkshelf/internal/domain"
)

func (s *LinkService) ensureURLIsFreeForUpdate(ctx context.Context, currentLinkID int64, finalURL string) error {
	links, err := s.storage.List(ctx)
	if err != nil {
		return err
	}

	for _, link := range links {
		if link.URL == finalURL && link.ID != currentLinkID {
			return ErrLinkAlreadyExists
		}
	}

	return nil
}

func linkInputHasChanges(current domain.Link, input linkInput) bool {
	if current.URL != input.URL {
		return true
	}

	if current.Title != input.Title {
		return true
	}

	return !tagsEqual(current.Tags, input.Tags)
}

func tagsEqual(left []string, right []string) bool {
	if len(left) != len(right) {
		return false
	}

	for i := range left {
		if left[i] != right[i] {
			return false
		}
	}

	return true
}
