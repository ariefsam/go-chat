package repository

import (
	"log"

	"github.com/ariefsam/go-chat/entity"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
)

type User struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
}

type userModel struct {
	MySQLID         int    `gorm:"column:id;PRIMARY_KEY;AUTO_INCREMENT"`
	ID              string `gorm:"column:user_id;unique_index"`
	Username        string
	Name            string
	PhoneNumber     string `gorm:"column:phone_number;unique_index"`
	IsPhoneVerified bool
}

func (userModel) TableName() string {
	return "user"
}
func (u *User) Save(user entity.User) (err error) {
	var cm, temp userModel
	copier.Copy(&cm, &user)

	db, err := u.connect()
	defer db.Close()
	if err != nil {
		log.Println(err)
		return
	}

	db.Where("chat_id=?", user.ID).Take(&temp)

	if temp.MySQLID == 0 {
		if err = db.Create(&cm).Error; err != nil {
			return
		}
	} else {
		cm.MySQLID = temp.MySQLID
		if err = db.Model(&cm).Update(&cm).Error; err != nil {
			return
		}
	}

	return
}

func (u *User) Get(filter entity.FilterUser) (listUsers []entity.User) {
	db, err := u.connect()
	if err != nil {
		return
	}
	defer db.Close()
	var limit int
	if filter.Limit == nil {
		limit = 10
	} else if limit > 10000 {
		limit = 10000
	} else {
		limit = *filter.Limit
	}

	if filter.UserID != nil {
		db = db.Where("user_id=?", *filter.UserID)
	}

	if filter.PhoneNumber != nil {
		db = db.Where("phone_number=?", *filter.PhoneNumber)
	}
	var users []userModel

	db.Order("id desc").Limit(limit).Find(&users)

	reverseSlice(users)
	copier.Copy(&listUsers, &users)

	return
}

func (u *User) Flush() (err error) {
	db, err := connect(u)
	if err != nil {
		return
	}
	err = db.Where("user_id!=?", "").Delete(userModel{}).Error

	return
}

func (u *User) AutoMigrate() {
	var um userModel

	db, err := connect(u)
	if err != nil {
		return
	}
	defer db.Close()

	db.AutoMigrate(&um)

}

func (u *User) connect() (db *gorm.DB, err error) {
	db, err = connect(u)
	return

}
