package mysql

import (
	"context"
	"github.com/pkg/errors"
	"lang-gpt-api/services/task/entity"
)

func (repo *mysqlRepo) UpdateTask(ctx context.Context, id int, data *entity.TaskDataUpdate) error {
	if err := repo.db.Table(data.TableName()).Where("id = ?", id).Updates(data).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
