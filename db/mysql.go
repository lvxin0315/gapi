package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lvxin0315/gapi/etc"
)

func MysqlDB(fs ...FuncWithDB) error {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		etc.MysqlUser,
		etc.MysqlPassword,
		etc.MysqlHost,
		etc.MysqlPort,
		etc.MysqlDatabase))
	db.LogMode(etc.MysqlLogMode)
	defer db.Close()
	if err != nil {
		return err
	}
	for _, f := range fs {
		err = f(db)
		if err != nil {
			return err
		}
	}
	return nil
}
