package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func (c *Chat) connect() (db *gorm.DB, err error) {
	db, err = gorm.Open("mysql", c.Username+":"+c.Password+"@tcp("+c.Host+":3306)/"+c.DatabaseName+"?charset=utf8&parseTime=True&loc=Local")
	db.LogMode(false)
	return
}
