package main

import (
	"github.com/jinzhu/gorm"
	"github.com/lvxin0315/gapi/db"
	"github.com/lvxin0315/gapi/models"
	"github.com/lvxin0315/gapi/runner/gin"
)

func main() {
	autoMigrate()
	// @title Docker监控服务
	// @version 1.0
	// @description docker监控服务后端API接口文档
	// @contact.name API Support
	// @contact.url http://www.swagger.io/support
	// @license.name Apache 2.0
	// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
	// @host 127.0.0.1:9009
	// @BasePath
	gin.Run()
}

/**
 * @Author lvxin0315@163.com
 * @Description db迁移
 * @Date 2:11 下午 2020/12/8
 * @Param
 * @return
 **/
func autoMigrate() {
	_ = db.DefaultSqliteDB(func(db *gorm.DB) error {
		db.AutoMigrate(models.DemoModel{})
		return nil
	})
}
