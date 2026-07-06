package service

import "errors"

var (
	ErrEmptyURL          = errors.New("url is empty")
	ErrInvalidURL        = errors.New("url is invalid")
	ErrBrokenURL         = errors.New("url is broken")
	ErrLinkAlreadyExists = errors.New("link already exists")
	ErrLinkNotChanged    = errors.New("link was not changed")
)
