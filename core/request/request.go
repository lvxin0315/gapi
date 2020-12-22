package request

import (
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gapi/core/response"
	"reflect"
)

type ReqWithDefault interface {
	Default() //设置默认值
}

type ReqWithCustomVerification interface {
	CustomVerification() error //自定义验证
}

//绑定数据，并验证
func ShouldBind(c *gin.Context, request interface{}) error {
	//设置默认值
	hasDefault(request)

	//获取参数
	if err := c.ShouldBind(request); err != nil {
		return err
	}

	//自定义验证
	if err := hasCustomVerification(request); err != nil {
		return err
	}

	return nil
}

func ShouldBindAndResponse(c *gin.Context, request interface{}) error {
	err := ShouldBind(c, request)
	if err != nil {
		response.Fail(c, nil, err.Error())
	}
	return err
}

func hasDefault(request interface{}) {
	val := reflect.ValueOf(request)
	if fun := val.MethodByName("Default"); fun.IsValid() {
		request.(ReqWithDefault).Default()
	}
}

func hasCustomVerification(request interface{}) error {
	val := reflect.ValueOf(request)
	if fun := val.MethodByName("CustomVerification"); fun.IsValid() {
		return request.(ReqWithCustomVerification).CustomVerification()
	}
	return nil
}
