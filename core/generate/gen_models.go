package generate

import (
	"bytes"
	"fmt"
	"github.com/Shelnutt2/db2struct"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/lvxin0315/gapi/core/tools"
	"github.com/lvxin0315/gapi/db"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

/**
 * @Author lvxin0315@163.com
 * @Description mysql生成对应model
 * @Date 11:29 上午 2020/12/11
 **/
type GenModel struct {
	MariadbUser     string
	MariadbPassword string
	MariadbHost     string
	MariadbPort     int
	MariadbDatabase string
	ModelDir        string
}

var importCode = `

import "database/sql"
`

type mysqlTable struct {
	TableName string `gorm:"column:TABLE_NAME"`
}

var modelCodes []string

func (gen *GenModel) AutoModel() {
	logrus.Info("autoModel")
	var tables []mysqlTable
	err := db.MysqlDB(func(db *gorm.DB) error {
		err := db.Raw("SELECT * FROM `information_schema`.`TABLES` WHERE `TABLE_SCHEMA` LIKE ?", gen.MariadbDatabase).Scan(&tables).Error
		if err != nil {
			panic(err)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	logrus.Info("tables len:", len(tables))
	for _, t := range tables {
		gen.writeModelFile(t.TableName)
	}
	//models code
	logrus.Info("auto models code :")
	for _, code := range modelCodes {
		logrus.Info(code)
	}
}

func (gen *GenModel) writeModelFile(mariadbTable string) {
	columnDataTypes, columnsSorted, err := db2struct.GetColumnsFromMysqlTable(
		gen.MariadbUser,
		gen.MariadbPassword,
		gen.MariadbHost,
		gen.MariadbPort,
		gen.MariadbDatabase,
		mariadbTable)
	if err != nil {
		panic(err)
	}
	//表名 -> 首字母大写 + Model
	structName := tools.CamelString(mariadbTable) + "Model"
	modelCodes = append(modelCodes, fmt.Sprintf("models.%s{},", structName))
	structCode, err := db2struct.Generate(*columnDataTypes, columnsSorted, mariadbTable, structName, "models", true, true, false)
	if err != nil {
		logrus.Info("mariadbTable:", mariadbTable)
		panic(err)
	}
	//添加import内容
	structCode = gen.importSqlPackage(structCode)
	//写入文件
	err = ioutil.WriteFile(fmt.Sprintf("%s/%s_model.go", gen.ModelDir, mariadbTable), structCode, 0755)
	if err != nil {
		panic(err)
	}
}

func (gen *GenModel) importSqlPackage(structCode []byte) []byte {
	if bytes.Index(structCode, []byte("sql.")) < 0 {
		return structCode
	}
	index := len("package models")
	return append(structCode[:index], append([]byte(importCode), structCode[index:]...)...)
}
