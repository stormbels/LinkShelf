package service

import (
	"context"
	"net/http"
	"net/url"
	"time"
)

func resolveReachableURL(ctx context.Context, linkURL string) (string, error) {
	if httpsURL, ok := upgradeHTTPToHTTPS(linkURL); ok {
		finalURL, err := resolveFinalURL(ctx, httpsURL)
		if err == nil {
			return finalURL, nil
		}
	}

	return resolveFinalURL(ctx, linkURL)
}

func upgradeHTTPToHTTPS(linkURL string) (string, bool) {
	parsedURL, err := url.Parse(linkURL)
	if err != nil || parsedURL.Scheme != "http" {
		return "", false
	}

	parsedURL.Scheme = "https"
	return parsedURL.String(), true
}

func resolveFinalURL(ctx context.Context, linkURL string) (string, error) {
	client := http.Client{
		Timeout: 4 * time.Second,
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodHead, linkURL, nil)
	if err != nil {
		return "", ErrInvalidURL
	}

	response, err := client.Do(request)
	if err != nil {
		return "", ErrBrokenURL
	}
	defer response.Body.Close()

	finalURL, _, err := parseLinkURL(response.Request.URL.String())
	if err != nil {
		return "", err
	}

	return finalURL, nil
}
