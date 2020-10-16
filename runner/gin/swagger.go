package gin

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lvxin0315/gapi/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/gen"
)

//生产swagger对应的json文件
func forSwaggerJson() {
	gen.New().Build(&gen.Config{
		SearchDir:   "./",
		OutputDir:   "./docs",
		MainAPIFile: "main.go",
	})
}

//设置swagger访问路由
func swaggerRouter(r *gin.RouterGroup) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
