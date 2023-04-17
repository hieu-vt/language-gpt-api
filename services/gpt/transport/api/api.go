package api

import (
	"context"
	sctx "github.com/viettranx/service-context"
	"github.com/viettranx/service-context/core"
	"lang-gpt-api/services/gpt/entity"
)

type ServiceContext interface {
	sctx.ServiceContext
	Business
}

type Business interface {
	GetMessagesBiz(ctx context.Context, paging *core.Paging) (error, []entity.GptContexts)
	CreateMessageBiz(ctx context.Context, body entity.CreateGptContexts) error
}

type api struct {
	serviceCtx sctx.ServiceContext
	business   Business
}

func NewAPI(serviceCtx sctx.ServiceContext, business Business) *api {
	return &api{
		serviceCtx: serviceCtx,
		business:   business,
	}
}
