package repository

import (
	"log"

	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
func connect(connection interface{}) (db *gorm.DB, err error) {
	type Connection struct {
		Host         string
		Username     string
		Password     string
		DatabaseName string
	}
	var c Connection

	copier.Copy(&c, connection)

	db, err = gorm.Open("mysql", c.Username+":"+c.Password+"@tcp("+c.Host+":3306)/"+c.DatabaseName+"?charset=utf8&parseTime=True&loc=Local&timeout=5s")
	db.LogMode(false)
	return
}
