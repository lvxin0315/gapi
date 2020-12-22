package services

import (
	"github.com/jinzhu/gorm"
	"github.com/lvxin0315/gapi/db"
	"github.com/lvxin0315/gapi/models"
	"github.com/sirupsen/logrus"
)

func GetAllEbStoreCategoryByIndex(where map[string]interface{}, limit int) ([]models.EbStoreCategoryModel, error) {
	pid := -1
	if where["pid"] != nil {
		pid = where["pid"].(int)
	}
	logrus.Info("pid:", pid)

	var storeCategoryModelList []models.EbStoreCategoryModel
	err := db.MysqlDB(func(db *gorm.DB) error {
		newDB := db.Model(&models.EbStoreCategoryModel{}).Where("is_show = ?", 1).Attrs("id", "cate_name", "pid", "pic")
		switch pid {
		case -1:
			newDB = newDB.Where("pid = 0")
		case 0:
			newDB = newDB.Where("pid > 0")
		default:
			newDB = newDB.Where("pid = ?", pid)
		}
		if where["name"] != nil && where["name"] != "" {
			newDB = newDB.Where("id = ? OR cate_name LIKE ?", where["name"], "%"+where["name"].(string)+"%")
		}
		return newDB.Order("sort DESC").Limit(limit).Find(&storeCategoryModelList).Error
	})
	return storeCategoryModelList, err
}

func GetAllEbStoreCategoryWithChildren() ([]models.EbStoreCategoryModel, error) {
	var storeCategoryModelList []models.EbStoreCategoryModel
	err := db.MysqlDB(func(db *gorm.DB) error {
		newDB := db.Model(&models.EbStoreCategoryModel{}).Where("is_show = ?", 1).Attrs("id", "cate_name", "pid", "pic")
		newDB = newDB.Where("pid = 0").Preload("EbStoreCategoryModels")
		return newDB.Order("sort DESC,id DESC").Find(&storeCategoryModelList).Error
	})
	return storeCategoryModelList, err
}
