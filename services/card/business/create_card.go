package business

import (
	"context"
	"github.com/viettranx/service-context/core"
	"lang-gpt-api/services/card/entity"
)

func (c *cardBusiness) CreateCard(ctx context.Context, data entity.Card) error {
	if err := c.cardRepository.Create(ctx, data); err != nil {
		return core.ErrInternalServerError.WithError(entity.ErrCannotCreateCard.Error())
	}

	return nil
}
