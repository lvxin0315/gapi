package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/lvxin0315/gapi/etc"
)

type FuncWithDB func(db *gorm.DB) error

func SqliteDB(dbName string, fs ...FuncWithDB) error {
	db, err := gorm.Open("sqlite3", fmt.Sprintf("sqlite/%s.db", dbName))
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

func DefaultSqliteDB(fs ...FuncWithDB) error {
	return SqliteDB(etc.SqliteDatabaseName, fs...)
}
