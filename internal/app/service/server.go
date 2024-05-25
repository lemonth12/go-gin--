package service

import (
	"fmt"
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
	fmt.Println("3333333333333333333333333333333333")
	aa := sv.db.Find()
	fmt.Println("*************************************", aa)
	return aa
}
