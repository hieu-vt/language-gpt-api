package mysql

import (
	"context"
	"github.com/pkg/errors"
	"github.com/viettranx/service-context/core"
	"lang-gpt-api/services/gpt/entity"
)

func (s *mysqlRepo) FindMessages(ctx context.Context, userId int, paging *core.Paging, moreKeys ...string) (error, []entity.GptContexts) {
	db := s.db.GetDB()

	var result []entity.GptContexts
	db = db.Table(entity.GptContexts{}.TableName())

	for i := range moreKeys {
		db.Preload(moreKeys[i])
	}

	db = db.Where("user_id = ?", userId)

	// count total
	if err := db.Select("id").Count(&paging.Total).Error; err != nil {
		return errors.WithStack(err), nil
	}

	//Query data with paging
	if err := db.Select("*").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).
		Error; err != nil {
		return errors.WithStack(err), nil
	}

	return nil, result
}
