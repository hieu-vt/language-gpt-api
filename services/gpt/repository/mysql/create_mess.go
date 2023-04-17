package mysql

import (
	"context"
	"lang-gpt-api/services/gpt/entity"
)

func (s *mysqlRepo) CreateMess(ctx context.Context, body entity.GptContexts) error {
	db := s.db.GetDB()
	if err := db.Create(&body).Error; err != nil {
		return err
	}

	return nil
}
