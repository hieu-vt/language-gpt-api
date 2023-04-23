package mysql

import "lang-gpt-api/common"

type mysqlRepo struct {
	db common.GormComponent
}

func NewMysqlRepoCard(db common.GormComponent) *mysqlRepo {
	return &mysqlRepo{db: db}
}
