package main

import (
	"bytes"
	"fmt"
	"github.com/Shelnutt2/db2struct"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/lvxin0315/gapi/db"
	"github.com/lvxin0315/gapi/etc"
	"github.com/lvxin0315/gapi/tools"
	"io/ioutil"
)

/**
 * @Author lvxin0315@163.com
 * @Description mysql生成对应model
 * @Date 11:29 上午 2020/12/11
 * @Param
 * @return
 **/

const ModelDir = "models"

var (
	mariadbUser     = "root"
	mariadbPassword = "root"
	mariadbHost     = "localhost"
	mariadbPort     = 3306
	mariadbDatabase = "sjyx"
)

var importCode = `

import "database/sql"
`

type mysqlTable struct {
	TableName string `gorm:"column:TABLE_NAME"`
}

var modelCodes []string

func main() {
	autoModel()
}

func autoModel() {
	fmt.Println("autoModel")
	var tables []mysqlTable
	err := db.MysqlDB(func(db *gorm.DB) error {
		fmt.Println("MysqlDB")
		err := db.Raw("SELECT * FROM `information_schema`.`TABLES` WHERE `TABLE_SCHEMA` LIKE ?", etc.MysqlDatabase).Scan(&tables).Error
		if err != nil {
			panic(err)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("tables len:", len(tables))
	for _, t := range tables {
		writeModelFile(t.TableName)
	}
	//models code
	fmt.Println("auto services code :")
	for _, code := range modelCodes {
		fmt.Println(code)
	}
}

func writeModelFile(mariadbTable string) {
	columnDataTypes, columnsSorted, err := db2struct.GetColumnsFromMysqlTable(mariadbUser, mariadbPassword, mariadbHost, mariadbPort, mariadbDatabase, mariadbTable)
	if err != nil {
		panic(err)
	}
	//表名 -> 首字母大写 + Model
	structName := tools.CamelString(mariadbTable) + "Model"
	modelCodes = append(modelCodes, fmt.Sprintf("models.%s{},", structName))
	structCode, err := db2struct.Generate(*columnDataTypes, columnsSorted, mariadbTable, structName, "models", true, true, false)
	if err != nil {
		fmt.Println("mariadbTable:", mariadbTable)
		panic(err)
	}
	//添加import内容
	structCode = importSqlPackage(structCode)
	//写入文件
	err = ioutil.WriteFile(fmt.Sprintf("%s/%s_model.go", ModelDir, mariadbTable), structCode, 0755)
	if err != nil {
		panic(err)
	}
}

func importSqlPackage(structCode []byte) []byte {
	if bytes.Index(structCode, []byte("sql.")) < 0 {
		return structCode
	}
	index := len("package models")
	return append(structCode[:index], append([]byte(importCode), structCode[index:]...)...)
}
