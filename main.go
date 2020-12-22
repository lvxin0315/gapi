package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/lvxin0315/gapi/controllers"
	"github.com/lvxin0315/gapi/controllers/product"
	"github.com/lvxin0315/gapi/db"
	_ "github.com/lvxin0315/gapi/docs"
	"github.com/lvxin0315/gapi/models"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/gen"
)

var ginEngine = gin.Default()

func main() {
	autoMigrate()
	// @title gapi
	// @version 1.0
	// @description gapi后端API接口文档
	// @contact.name API Support
	// @contact.url http://www.swagger.io/support
	// @license.name Apache 2.0
	// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
	// @host 127.0.0.1:8080
	// @BasePath
	ginRunner()

}

/**
 * @Author lvxin0315@163.com
 * @Description db迁移
 * @Date 2:11 下午 2020/12/8
 **/
func autoMigrate() {
	_ = db.DefaultSqliteDB(func(db *gorm.DB) error {
		db.LogMode(true)
		db.AutoMigrate(models.DemoModel{}, models.GoodsModel{}, models.CategoryModel{})
		return nil
	})
}

/**
 * @Author lvxin0315@163.com
 * @Description gin
 * @Date 5:01 下午 2020/12/8
 **/
func ginRunner() {
	addRouter()
	//生成swagger相关内容
	forSwaggerJson()
	//设置swagger路由
	swaggerRouter(&ginEngine.RouterGroup)
	//启动
	ginEngine.Run()
}

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

//路由
func addRouter() {
	ginEngine.GET("/", controllers.Index)
	ginEngine.GET("/index1", controllers.Index1)
	ginEngine.GET("/index2", controllers.Index2)
	ginEngine.GET("/product/v1.product.StoreProduct/index", product.Index)

	// api
	api := ginEngine.Group("/api")
	api.GET("category", controllers.Category)
}
