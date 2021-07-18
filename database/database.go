package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"google.com/banking-app/helpers"
)

var Db *gorm.DB

func DbConnection() {

	dsn := "root:1234@tcp(127.0.0.1:3306)/users?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open("mysql", dsn)
	if err != nil {
		er := &helpers.Error{ErrorMessage: "unable to connect to db", Err: err}
		er.HandleErr()
	}
	DB.DB().SetMaxIdleConns(20)
	DB.DB().SetMaxOpenConns(200)

	Db = DB

}
