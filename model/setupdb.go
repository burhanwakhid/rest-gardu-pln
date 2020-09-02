package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SetUpModels() *gorm.DB {
	db, err := gorm.Open("mysql", "root:burhan@(localhost)/gardupln_db?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("gagal koneksi database")
	}

	db.AutoMigrate(&GarduM{})
	return db
}
