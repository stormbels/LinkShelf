package service

import (
	"context"
	"strings"

	"github.com/stormbels/linkshelf/internal/domain"
)

func (s *LinkService) SearchLinks(ctx context.Context, rawQuery string) ([]domain.Link, error) {
	links, err := s.storage.List(ctx)
	if err != nil {
		return nil, err
	}

	query := strings.ToLower(strings.TrimSpace(rawQuery))
	if query == "" {
		return links, nil
	}

	foundLinks := make([]domain.Link, 0)
	for _, link := range links {
		if linkMatchesQuery(link, query) {
			foundLinks = append(foundLinks, link)
		}
	}

	return foundLinks, nil
}

func linkMatchesQuery(link domain.Link, query string) bool {
	searchTextParts := []string{link.Title, link.URL}
	searchTextParts = append(searchTextParts, link.Tags...)

	searchText := strings.ToLower(strings.Join(searchTextParts, " "))
	return strings.Contains(searchText, query)
}
