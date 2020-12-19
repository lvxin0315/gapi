package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/lvxin0315/gapi/core/generate"
	"github.com/lvxin0315/gapi/etc"
)

/**
 * @Author lvxin0315@163.com
 * @Description mysql生成对应model
 * @Date 11:29 上午 2020/12/11
 * @Param
 * @return
 **/

const ModelDir = "models"

func main() {
	genModel := generate.GenModel{
		MariadbUser:     etc.MysqlUser,
		MariadbPassword: etc.MysqlPassword,
		MariadbHost:     etc.MysqlHost,
		MariadbPort:     etc.MysqlPort,
		MariadbDatabase: etc.MysqlDatabase,
		ModelDir:        ModelDir,
	}
	genModel.AutoModel()
}
