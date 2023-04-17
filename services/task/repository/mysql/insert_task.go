package mysql

import (
	"context"
	"github.com/pkg/errors"
	"lang-gpt-api/services/task/entity"
)

func (repo *mysqlRepo) AddNewTask(ctx context.Context, data *entity.TaskDataCreation) error {
	if err := repo.db.Create(data).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
