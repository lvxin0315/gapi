package request

import (
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gapi/core/response"
)

type Request interface {
	Default()                  //设置默认值
	CustomVerification() error //自定义验证
}

//绑定数据，并验证
func ShouldBind(c *gin.Context, request Request) error {
	//设置默认值
	request.Default()
	if err := c.ShouldBind(request); err != nil {
		return err
	}
	//自定义验证
	if err := request.CustomVerification(); err != nil {
		return err
	}
	return nil
}

func ShouldBindAndResponse(c *gin.Context, request Request) error {
	err := ShouldBind(c, request)
	if err != nil {
		response.Fail(c, nil, err.Error())
	}
	return err
}
