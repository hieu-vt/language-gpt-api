package entity

import "errors"

var (
	ErrCannotCreateFlashcard = errors.New("cannot create flashcard")
	ErrCannotListFlashCard   = errors.New("cannot list flashcard")
)
