package mysql

import (
	"context"
	"lang-gpt-api/services/card/entity"
)

func (repo *mysqlRepo) Create(ctx context.Context, data entity.Card) error {
	db := repo.db.GetDB()

	if err := db.Create(&data).Error; err != nil {
		return err
	}

	return nil

}
