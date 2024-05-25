package db

import (
	"fmt"
	"gorm.io/gorm"
	"preject/internal/app/models/system"
)

type DataTableEntity struct {
	DataBase *gorm.DB
	Model    *system.TableName
}

func NewFindEntity(db *gorm.DB) *DataTableEntity {
	return &DataTableEntity{DataBase: db}
}

func (r *DataTableEntity) Find() int {
	fmt.Println("5555555555555555555555")
	//数据逻辑
	return 1
}
