package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gapi/models"
	"github.com/lvxin0315/gapi/services"
	"math/rand"
)

//测试一下
func Index(c *gin.Context) {
	//建分类
	category1Model := &models.CategoryModel{
		Title: "分类1",
	}
	category2Model := &models.CategoryModel{
		Title: "分类2",
	}
	category3Model := &models.CategoryModel{
		Title: "分类3",
	}
	services.SaveCategory(category1Model)
	services.SaveCategory(category2Model)
	services.SaveCategory(category3Model)
	fmt.Println(category1Model.ID)
	fmt.Println(category2Model.ID)
	fmt.Println(category3Model.ID)
	//建商品
	//分类1
	for i := 0; i < 10; i++ {
		goodsModel := &models.GoodsModel{
			Name:       fmt.Sprintf("分类商品1-%d", i),
			Stock:      rand.Intn(999),
			Price:      rand.Float64(),
			CategoryId: category1Model.ID,
		}
		services.SaveGoods(goodsModel)
	}
	//分类2
	for i := 0; i < 20; i++ {
		goodsModel := &models.GoodsModel{
			Name:       fmt.Sprintf("分类商品2-%d", i),
			Stock:      rand.Intn(999),
			Price:      rand.Float64(),
			CategoryId: category2Model.ID,
		}
		services.SaveGoods(goodsModel)
	}
	//分类3
	for i := 0; i < 5; i++ {
		goodsModel := &models.GoodsModel{
			Name:       fmt.Sprintf("分类商品3-%d", i),
			Stock:      rand.Intn(999),
			Price:      rand.Float64(),
			CategoryId: category3Model.ID,
		}
		services.SaveGoods(goodsModel)
	}
}

//测试查询
func Index1(c *gin.Context) {
	data, err := services.GetGoods(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(data.Name)

	dataList, err := services.GetGoodsList(nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(dataList))

	dataList1, pagination, err := services.GetGoodsListPage(nil, 1, 10)
	if err != nil {
		panic(err)
	}
	fmt.Println(pagination.Total)
	fmt.Println(len(dataList1))
	for _, d := range dataList1 {
		fmt.Println(d.Name)
	}

	dataList2, pagination, err := services.GetGoodsListPage(nil, 2, 10)
	if err != nil {
		panic(err)
	}
	fmt.Println(pagination.Total)
	fmt.Println(len(dataList2))
	for _, d := range dataList2 {
		fmt.Println(d.Name)
	}
}
