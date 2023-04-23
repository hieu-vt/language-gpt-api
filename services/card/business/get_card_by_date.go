package business

import (
	"context"
	"github.com/viettranx/service-context/core"
	"lang-gpt-api/services/card/entity"
)

func (c *cardBusiness) GetCardByDate(ctx context.Context, date string) (error, []entity.Card) {
	err, result := c.cardRepository.GetByDate(ctx, date)

	if err != nil {
		return core.ErrInternalServerError.WithError(entity.ErrCannotListCard.Error()), nil
	}

	return nil, result
}
