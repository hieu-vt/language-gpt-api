package api

import (
	"context"
	sctx "github.com/viettranx/service-context"
	"lang-gpt-api/services/card/entity"
)

type ServiceContext struct {
	sctx.ServiceContext
	Business
}

type Business interface {
	CreateCard(ctx context.Context, data entity.Card) error
	GetCardByDate(ctx context.Context, userId int, date string) (error, []entity.Card)
}

type api struct {
	sc       sctx.ServiceContext
	business Business
}

func NewApi(sc sctx.ServiceContext, biz Business) *api {
	return &api{
		sc:       sc,
		business: biz,
	}
}
