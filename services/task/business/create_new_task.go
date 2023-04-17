package business

import (
	"context"
	"github.com/viettranx/service-context/core"
	"lang-gpt-api/services/task/entity"
)

func (biz *business) CreateNewTask(ctx context.Context, data *entity.TaskDataCreation) error {
	requester := core.GetRequester(ctx)

	uid, _ := core.FromBase58(requester.GetSubject())
	requesterId := int(uid.GetLocalID()) // task owner id, id of who creates this new task

	data.Prepare(requesterId, entity.StatusDoing)

	if err := data.Validate(); err != nil {
		return core.ErrBadRequest.WithError(err.Error())
	}

	if err := biz.taskRepo.AddNewTask(ctx, data); err != nil {
		return core.ErrInternalServerError.WithError(entity.ErrCannotCreateTask.Error())
	}

	return nil
}
