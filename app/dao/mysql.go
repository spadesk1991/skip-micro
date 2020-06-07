package dao

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:123456@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	db.SetLogger(log.New(os.Stdout, "\r\n", 0))
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(50)
}

func GetDB() *gorm.DB {
	return db
}
