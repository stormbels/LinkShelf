package domain

import "time"

type Link struct {
	ID          int64
	URL         string
	Title       string
	Description string
	Tags        []string
	CreatedAt   time.Time
}
