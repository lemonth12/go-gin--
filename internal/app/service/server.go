package service

import (
	"preject/internal/app/db"
)

type Db interface {
	Find() int
}

type Service struct {
	db Db
}

func NewService() *Service {
	return &Service{
		db: db.NewFindEntity(db.MysqlClass3Client()),
	}
}

func (sv *Service) Logic() int {
	aa := sv.db.Find()
	return aa
}
