package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gapi/core/request"
	"github.com/lvxin0315/gapi/core/response"
	"github.com/lvxin0315/gapi/services"
	"github.com/sirupsen/logrus"
	"strings"
)

// v1.store.StoreProductController/lst
func Products(c *gin.Context) {
	productsRequest := productsRequest{}
	if err := request.ShouldBindAndResponse(c, &productsRequest); err != nil {
		return
	}
	//排序
	order := "id DESC"
	if productsRequest.PriceOrder != "" {
		order = "price " + productsRequest.PriceOrder
	}
	if productsRequest.SalesOrder != "" {
		order = "sales " + productsRequest.SalesOrder
	}
	//组装条件
	where := make(map[string]interface{})
	where["is_show"] = 1
	where["is_del"] = 0
	if productsRequest.Ids != "" {
		where["id in"] = strings.Split(productsRequest.Ids, ",")
	}
	//商品名称
	if productsRequest.Keyword != "" {
		where["store_name like"] = "%" + productsRequest.Keyword + "%"
	}
	//一级分类处理
	if productsRequest.Cid != 0 {
		productIds, err := services.GetProductIdsByCid(productsRequest.Cid)
		if err != nil {
			response.Fail(c, err, err.Error())
			return
		}
		where["id in"] = productIds
	}
	//二级分类处理
	if productsRequest.Sid != 0 {
		productIds, err := services.GetProductIdsBySid(productsRequest.Sid)
		if err != nil {
			response.Fail(c, err, err.Error())
			return
		}
		where["id in"] = productIds
	}
	//查询
	dataList, pagination, err := services.GetSearchList(where, productsRequest.Page, productsRequest.Limit, order)
	if err != nil {
		response.Fail(c, err, err.Error())
		return
	}
	logrus.Info(pagination)
	response.Success(c, dataList)
}
