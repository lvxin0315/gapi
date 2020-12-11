package models

import "github.com/jinzhu/gorm"

type CategoryModel struct {
	gorm.Model
	Title string
}

func (CategoryModel) TableName() string {
	return "category"
}
