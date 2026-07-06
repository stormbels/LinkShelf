package service

import (
	"net/url"
	"strings"

	"github.com/stormbels/linkshelf/internal/domain"
)

type linkInput struct {
	URL   string
	Title string
	Tags  []string
}

func (input linkInput) toDomainLink() domain.Link {
	return domain.Link{
		URL:   input.URL,
		Title: input.Title,
		Tags:  input.Tags,
	}
}

func normalizeLinkInput(rawURL string, rawTitle string, rawTags string) (linkInput, error) {
	cleanURL, host, err := parseLinkURL(rawURL)
	if err != nil {
		return linkInput{}, err
	}

	title := strings.TrimSpace(rawTitle)
	if title == "" {
		title = host
	}

	return linkInput{
		URL:   cleanURL,
		Title: title,
		Tags:  parseTags(rawTags),
	}, nil
}

func parseLinkURL(rawURL string) (cleanURL string, host string, err error) {
	cleanURL = strings.TrimSpace(rawURL)
	if cleanURL == "" {
		return "", "", ErrEmptyURL
	}

	parsedURL, err := url.ParseRequestURI(cleanURL)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		return "", "", ErrInvalidURL
	}

	parsedURL.Path = strings.TrimRight(parsedURL.Path, "/")

	return parsedURL.String(), parsedURL.Host, nil
}

func parseTags(rawTags string) []string {
	parts := strings.Split(rawTags, ",")
	tags := make([]string, 0, len(parts))
	seenTags := make(map[string]struct{}, len(parts))

	for _, part := range parts {
		tag := strings.ToLower(strings.TrimSpace(part))
		if tag == "" {
			continue
		}

		if _, exists := seenTags[tag]; exists {
			continue
		}

		seenTags[tag] = struct{}{}
		tags = append(tags, tag)
	}

	return tags
}
