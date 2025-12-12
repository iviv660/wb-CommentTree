package model

import "errors"

var (
	ErrNotFound        = errors.New("not found")
	ErrNotFoundComment = errors.New("comment not found")
	ErrInvalidInput    = errors.New("invalid input")
	ErrEmptyBody       = errors.New("empty body")
)
