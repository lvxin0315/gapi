package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gapi/core/request"
	"github.com/lvxin0315/gapi/core/response"
	"github.com/lvxin0315/gapi/services"
)

// v1.store.CategoryController/category
func Category(c *gin.Context) {
	categoryRequest := categoryRequest{}
	if err := request.ShouldBindAndResponse(c, &categoryRequest); err != nil {
		return
	}
	if categoryRequest.Limit <= 0 {
		dataList, err := services.GetAllEbStoreCategoryWithChildren()
		if err != nil {
			response.Fail(c, nil, err.Error())
			return
		}
		response.Success(c, dataList)
	} else {
		dataList, err := services.GetAllEbStoreCategoryByIndex(map[string]interface{}{
			"pid":  categoryRequest.Pid,
			"name": categoryRequest.Name,
		}, categoryRequest.Limit)
		if err != nil {
			response.Fail(c, nil, err.Error())
			return
		}
		response.Success(c, dataList)
	}
}
