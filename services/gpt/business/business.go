package business

import (
	"context"
	"github.com/viettranx/service-context/core"
	"lang-gpt-api/services/gpt/entity"
)

type GptRepository interface {
	CreateMess(ctx context.Context, body entity.GptContexts) error
	FindMessages(ctx context.Context, userId int, paging *core.Paging, moreKeys ...string) (error, []entity.GptContexts)
}

type gptBusiness struct {
	gptRepository GptRepository
}

func NewGptBusiness(repo GptRepository) *gptBusiness {
	return &gptBusiness{gptRepository: repo}
}
