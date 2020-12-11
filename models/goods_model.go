package models

import "github.com/jinzhu/gorm"

type GoodsModel struct {
	gorm.Model
	Name       string
	Stock      int
	Price      float64
	CategoryId uint
}

func (GoodsModel) TableName() string {
	return "goods"
}
