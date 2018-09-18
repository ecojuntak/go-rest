package model

import (
	"github.com/jinzhu/gorm"
)

func Migrate() error {
	db, err := gorm.Open("mysql", "root:eco123@/gorest?charset=utf8&parseTime=True&loc=Local")
	db.AutoMigrate(&User{})

	return err
}
