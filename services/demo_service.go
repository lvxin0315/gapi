package services

import (
	"github.com/jinzhu/gorm"
	"github.com/lvxin0315/gapi/db"
	"github.com/lvxin0315/gapi/models"
)

type DemoService struct {
}

func NewDemoService() *DemoService {
	return new(DemoService)
}

func (service *DemoService) Save(data *models.DemoModel) error {
	return db.DefaultSqliteDB(func(db *gorm.DB) error {
		return db.Save(data).Error
	})
}

func (service *DemoService) Find(dataList *[]*models.DemoModel, where ...interface{}) error {
	err := db.DefaultSqliteDB(func(db *gorm.DB) error {
		return db.Order("id DESC").Find(dataList, where...).Error
	})
	return err
}

func (service *DemoService) First(data *models.DemoModel) error {
	return db.DefaultSqliteDB(func(db *gorm.DB) error {
		return db.First(data).Error
	})
}

func (service *DemoService) Delete(data *models.DemoModel) error {
	return db.DefaultSqliteDB(func(db *gorm.DB) error {
		return db.Delete(data).Error
	})
}

//update xxx set xxx=xxx where id = xxx
func (service *DemoService) UpdateWithId(id uint, data map[string]interface{}) error {
	return db.DefaultSqliteDB(func(db *gorm.DB) error {
		return db.Model(models.DemoModel{}).Where("id = ?", id).Update(data).Error
	})
}
