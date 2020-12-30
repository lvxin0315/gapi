package services

import (
	"github.com/jinzhu/gorm"
	"github.com/lvxin0315/gapi/core/tools"
	"github.com/lvxin0315/gapi/db"
	"github.com/lvxin0315/gapi/models"
)

/**
 * @Description 保存
 **/
func SaveDemo(m *models.DemoModel) error {
	return db.MysqlDB(func(db *gorm.DB) error {
		return db.Save(m).Error
	})
}

/**
 * @Description 删除
 **/
func DeleteDemoByIds(ids ...uint) error {
	return db.MysqlDB(func(db *gorm.DB) error {
		return db.Delete(&models.DemoModel{}, "id in (?)", ids).Error
	})
}

/**
 * @Description 查询
 **/
func GetDemo(id uint) (data models.DemoModel, err error) {
	err = db.MysqlDB(func(db *gorm.DB) error {
		return db.Where("id = ?", id).First(&data).Error
	})
	return
}

/**
 * @Description 列表
 **/
func GetDemoList(where map[string]interface{}, limit int, order string) (dataList []models.DemoModel, err error) {
	if limit <= 0 {
		limit = 1000
	}
	whereSql, valueList, err := tools.WhereBuild(where)
	if err != nil {
		return nil, err
	}
	err = db.MysqlDB(func(db *gorm.DB) error {
		return db.Where(whereSql, valueList...).Limit(limit).Order(order).Find(&dataList).Error
	})
	return
}

/**
 * @Description 列表（分页）
 **/
func GetDemoListPage(where map[string]interface{}, page uint, pageSize uint, order string) (dataList []models.DemoModel, pagination tools.Pagination, err error) {
	offset := 0
	if page > 1 {
		offset = int((page - 1) * pageSize)
	}
	whereSql, valueList, err := tools.WhereBuild(where)
	if err != nil {
		return
	}
	err = db.MysqlDB(func(db *gorm.DB) error {
		theDB := db.Model(&models.DemoModel{}).Where(whereSql, valueList...)
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
		return theDB.Offset(offset).Limit(pageSize).Order(order).Find(&dataList).Error
	})
	return
}

/**
 * @Description 批量更新
 **/
func UpdateDemoByIds(values map[string]interface{}, ids ...uint) error {
	return db.MysqlDB(func(db *gorm.DB) error {
		return db.Model(&models.DemoModel{}).Where("id in (?)", ids).Update(values).Error
	})
}

/**
 * @Description 字段值减少
 **/
func SetDemoDec(column string, value int) error {
	return db.MysqlDB(func(db *gorm.DB) error {
		return db.Model(&models.DemoModel{}).UpdateColumn(column, gorm.Expr("? - ?", column, value)).Error
	})
}

/**
 * @Description 字段值增加
 **/
func SetDemoInc(column string, value int) error {
	return db.MysqlDB(func(db *gorm.DB) error {
		return db.Model(&models.DemoModel{}).UpdateColumn(column, gorm.Expr("? + ?", column, value)).Error
	})
}

/**
 * @Description 通用获取
 **/
func GetDemoByWhere(where map[string]interface{}) (data models.DemoModel, err error) {
	whereSql, valueList, err := tools.WhereBuild(where)
	if err != nil {
		return
	}
	err = db.MysqlDB(func(db *gorm.DB) error {
		return db.Where(whereSql, valueList...).First(&data).Error
	})
	return
}
