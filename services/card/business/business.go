package business

import (
	"context"
	"lang-gpt-api/services/card/entity"
)

type cardRepository interface {
	Create(ctx context.Context, data entity.Card) error
	GetByDate(ctx context.Context, userId int, createdAt string) (error, []entity.Card)
}

type cardBusiness struct {
	cardRepository cardRepository
}

func NewCardBusiness(repository cardRepository) *cardBusiness {
	return &cardBusiness{cardRepository: repository}
}
