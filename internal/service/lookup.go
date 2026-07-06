package service

import (
	"context"

	"github.com/stormbels/linkshelf/internal/domain"
)

func (s *LinkService) FindLinkByURL(ctx context.Context, cleanURL string) (domain.Link, bool, error) {
	links, err := s.storage.List(ctx)
	if err != nil {
		return domain.Link{}, false, err
	}

	for _, link := range links {
		if link.URL == cleanURL {
			return link, true, nil
		}
	}

	return domain.Link{}, false, nil
}

func (s *LinkService) FindDuplicateByRawURL(ctx context.Context, rawURL string) (domain.Link, bool, error) {
	input, err := normalizeLinkInput(rawURL, "", "")
	if err != nil {
		return domain.Link{}, false, err
	}

	finalURL, err := resolveReachableURL(ctx, input.URL)
	if err != nil {
		return domain.Link{}, false, err
	}

	return s.FindLinkByURL(ctx, finalURL)
}

func (s *LinkService) findLinkByID(ctx context.Context, id int64) (domain.Link, bool, error) {
	links, err := s.storage.List(ctx)
	if err != nil {
		return domain.Link{}, false, err
	}

	for _, link := range links {
		if link.ID == id {
			return link, true, nil
		}
	}

	return domain.Link{}, false, nil
}
