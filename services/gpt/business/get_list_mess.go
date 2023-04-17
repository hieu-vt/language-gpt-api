package business

import (
	"context"
	"github.com/viettranx/service-context/core"
	"lang-gpt-api/services/gpt/entity"
)

func (biz *gptBusiness) GetMessagesBiz(ctx context.Context, paging *core.Paging) (error, []entity.GptContexts) {
	requester := core.GetRequester(ctx)

	uid, _ := core.FromBase58(requester.GetSubject())
	requesterId := int(uid.GetLocalID()) // task owner id, id of who creates this new task

	err, result := biz.gptRepository.FindMessages(ctx, requesterId, paging)

	if err != nil {
		return core.ErrInternalServerError.WithError(entity.ErrCannotListMessageGpt.Error()), nil
	}

	return nil, result
}
