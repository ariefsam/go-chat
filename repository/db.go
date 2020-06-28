package repository

import (
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func connect(connection interface{}) (db *gorm.DB, err error) {
	type Connection struct {
		Host         string
		Username     string
		Password     string
		DatabaseName string
	}
	var c Connection

	copier.Copy(&c, connection)

	db, err = gorm.Open("mysql", c.Username+":"+c.Password+"@tcp("+c.Host+":3306)/"+c.DatabaseName+"?charset=utf8&parseTime=True&loc=Local")
	db.LogMode(false)
	return
}
