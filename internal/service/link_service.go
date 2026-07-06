package service

import (
	"context"
	"time"

	"github.com/stormbels/linkshelf/internal/domain"
	"github.com/stormbels/linkshelf/internal/storage"
)

type LinkService struct {
	storage storage.LinkStorage
}

func NewLinkService(storage storage.LinkStorage) *LinkService {
	return &LinkService{storage: storage}
}

func (s *LinkService) CreateLink(ctx context.Context, rawURL string, rawTitle string, rawTags string) error {
	input, err := normalizeLinkInput(rawURL, rawTitle, rawTags)
	if err != nil {
		return err
	}

	finalURL, err := resolveReachableURL(ctx, input.URL)
	if err != nil {
		return err
	}
	input.URL = finalURL

	exists, err := s.storage.ExistsByURL(ctx, input.URL)
	if err != nil {
		return err
	}
	if exists {
		return ErrLinkAlreadyExists
	}

	link := input.toDomainLink()
	link.CreatedAt = time.Now()

	return s.storage.Save(ctx, link)
}

func (s *LinkService) UpdateLink(ctx context.Context, id int64, rawURL string, rawTitle string, rawTags string) error {
	input, err := normalizeLinkInput(rawURL, rawTitle, rawTags)
	if err != nil {
		return err
	}

	currentLink, found, err := s.findLinkByID(ctx, id)
	if err != nil {
		return err
	}

	finalURL, err := resolveReachableURL(ctx, input.URL)
	if err != nil {
		return err
	}
	input.URL = finalURL

	if found && !linkInputHasChanges(currentLink, input) {
		return ErrLinkNotChanged
	}

	if err := s.ensureURLIsFreeForUpdate(ctx, id, input.URL); err != nil {
		return err
	}

	updatedLink := input.toDomainLink()

	return s.storage.Update(ctx, id, updatedLink)
}

func (s *LinkService) ListLinks(ctx context.Context) ([]domain.Link, error) {
	return s.storage.List(ctx)
}

func (s *LinkService) DeleteLink(ctx context.Context, id int64) error {
	return s.storage.Delete(ctx, id)
}
