package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	dbConnect, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/go-books?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = dbConnect
}

func GetDB() *gorm.DB {
	return db
}
