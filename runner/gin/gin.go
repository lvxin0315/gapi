package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gapi/controller/goods"
)

func Run() {
	r := gin.Default()
	//生成swagger相关内容
	forSwaggerJson()
	//设置swagger路由
	swaggerRouter(&r.RouterGroup)
	//加载路由
	addRouter(r)
	//启动
	r.Run()
}

func addRouter(r *gin.Engine) {
	r.GET("/goods/:id", goods.GetOne)
}
