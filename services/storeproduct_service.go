package services

import (
	"github.com/jinzhu/gorm"
	"github.com/lvxin0315/gapi/core/tools"
	"github.com/lvxin0315/gapi/db"
	"github.com/lvxin0315/gapi/models"
)

//分类查商品
func GetProductIdsBySid(sid int) ([]int, error) {
	var ids []int
	ebStoreProductCateModelList, err := GetEbStoreProductCateList(map[string]interface{}{
		"cate_id": sid,
	}, -1, "")
	if err != nil {
		return ids, err
	}
	for _, productCate := range ebStoreProductCateModelList {
		ids = append(ids, productCate.ProductID)
	}
	return ids, nil
}

func GetProductIdsByCid(cid int) ([]int, error) {
	var ids []int
	var sidList []int
	//获取二级分类
	ebCategoryModelList, err := GetEbStoreCategoryList(map[string]interface{}{
		"pid": cid,
	}, -1, "")
	if err != nil {
		return ids, err
	}
	for _, category := range ebCategoryModelList {
		sidList = append(sidList, category.ID)
	}
	ebStoreProductCateModelList, err := GetEbStoreProductCateList(map[string]interface{}{
		"cate_id in": sidList,
	}, -1, "")
	if err != nil {
		return ids, err
	}
	for _, productCate := range ebStoreProductCateModelList {
		ids = append(ids, productCate.ProductID)
	}
	return ids, nil
}

func GetSearchList(where map[string]interface{}, page uint, pageSize uint, order string) (dataList []models.EbStoreProductModel, pagination tools.Pagination, err error) {
	offset := 0
	if page > 1 {
		offset = int((page - 1) * pageSize)
	}
	whereSql, valueList, err := tools.WhereBuild(where)
	if err != nil {
		return
	}
	err = db.MysqlDB(func(db *gorm.DB) error {
		theDB := db.Model(&models.EbStoreProductModel{}).Where(whereSql, valueList...)
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
		return theDB.Offset(offset).
			Limit(pageSize).
			Order(order).
			Select("id,store_name,cate_id,image,IFNULL(sales, 0) + IFNULL(ficti, 0) as sales,price,stock,activity,ot_price,spec_type,recommend_image,unit_name").
			Find(&dataList).Error
	})
	return
}
