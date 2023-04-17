package mysql

import (
	"context"
	"github.com/pkg/errors"
	"lang-gpt-api/services/user/entity"
)

func (repo *mysqlRepo) CreateNewUser(ctx context.Context, data *entity.UserDataCreation) error {
	if err := repo.db.Table(data.TableName()).Create(data).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}
