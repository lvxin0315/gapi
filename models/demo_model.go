package models

import "github.com/jinzhu/gorm"

type DemoModel struct {
	gorm.Model
	Name string
	Age  uint
}

func (DemoModel) TableName() string {
	return "demo"
}

func (model *DemoModel) GetModel() interface{} {
	return new(DemoModel)
}

func (model *DemoModel) GetModelList() interface{} {
	var demoList []*DemoModel
	return demoList
}
