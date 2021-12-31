package config

import (
	"github.com/jinzhu/gorm"
)

func DBInit() *gorm.DB {

	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3307)/aset_ssh")

	if err != nil {
		panic("falied to connect to database")
	}
	db.AutoMigrate()

	return db
}
