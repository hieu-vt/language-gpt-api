package entity

import "errors"

var (
	ErrCannotCreateCard = errors.New("cannot create card")
	ErrCannotListCard   = errors.New("cannot list cards")
)
