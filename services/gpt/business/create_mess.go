package business

import (
	"context"
	"github.com/viettranx/service-context/core"
	"lang-gpt-api/services/gpt/entity"
)

func (biz *gptBusiness) CreateMessageBiz(ctx context.Context, body entity.CreateGptContexts) error {
	requester := core.GetRequester(ctx)

	uid, _ := core.FromBase58(requester.GetSubject())
	requesterId := int(uid.GetLocalID()) // task owner id, id of who creates this new task

	err := biz.gptRepository.CreateMess(ctx, entity.GptContexts{
		UserId:      requesterId,
		SendMessage: body.SendMessage,
		GptMessage:  body.GptMessage,
	})

	if err != nil {
		return core.ErrInternalServerError.WithError(entity.ErrCannotCreateGptMessage.Error())
	}

	return nil
}
