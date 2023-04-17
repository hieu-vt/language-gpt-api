package entity

import "errors"

var (
	ErrCannotCreateGptMessage = errors.New("cannot create gpt message")
	ErrCannotListMessageGpt   = errors.New("cannot list gpt messages")
)
