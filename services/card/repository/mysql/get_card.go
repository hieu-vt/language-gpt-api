package mysql

import (
	"context"
	"lang-gpt-api/services/card/entity"
)

func (repo *mysqlRepo) GetByDate(ctx context.Context, userId int, createdAt string) (error, []entity.Card) {
	db := repo.db.GetDB()
	var result []entity.Card

	if err := db.Table(entity.Card{}.TableName()).Where("user_id = ?", userId).Where("date(created_at) = ?", createdAt).Find(&result).Error; err != nil {
		return err, nil
	}

	return nil, result
}
