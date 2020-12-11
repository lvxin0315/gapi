package services

import (
	"github.com/jinzhu/gorm"
	"github.com/lvxin0315/gapi/db"
	"github.com/lvxin0315/gapi/models"
	"github.com/lvxin0315/gapi/tools"
)

/**
 * @Description 保存
 **/
func SaveDemo(m *models.DemoModel) error {
	return db.DefaultSqliteDB(func(db *gorm.DB) error {
		return db.Save(m).Error
	})
}

/**
 * @Description 删除
 **/
func DeleteDemoByIds(ids ...uint) error {
	return db.DefaultSqliteDB(func(db *gorm.DB) error {
		return db.Delete(&models.DemoModel{}, "id in (?)", ids).Error
	})
}

/**
 * @Description 查询
 **/
func GetDemo(id uint) (data models.DemoModel, err error) {
	err = db.DefaultSqliteDB(func(db *gorm.DB) error {
		return db.Where("id = ?", id).First(&data).Error
	})
	return
}

/**
 * @Description 列表
 **/
func GetDemoList(where map[string]interface{}) (dataList []models.DemoModel, err error) {
	err = db.DefaultSqliteDB(func(db *gorm.DB) error {
		return db.Where(where).Find(&dataList).Error
	})
	return
}

/**
 * @Description 列表（分页）
 **/
func GetDemoListPage(where map[string]interface{}, page uint, pageSize uint) (dataList []models.DemoModel, pagination tools.Pagination, err error) {
	offset := 0
	if page > 1 {
		offset = int((page - 1) * pageSize)
	}
	err = db.DefaultSqliteDB(func(db *gorm.DB) error {
		theDB := db.Model(&models.DemoModel{}).Where(where)
		//total
		total := 0
		err := theDB.Count(&total).Error
		if err != nil {
			return err
		}
		pagination, err = tools.NewPagination(uint(total), page, pageSize, where)
		if err != nil {
			return err
		}
		//dataList
		return theDB.Offset(offset).Limit(pageSize).Find(&dataList).Error
	})
	return
}
