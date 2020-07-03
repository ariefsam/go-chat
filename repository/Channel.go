package repository

import (
	"log"

	"github.com/ariefsam/go-chat/entity"
	"github.com/jinzhu/copier"
)

type Channel struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
}

type channelModel struct {
	MySQLID int    `gorm:"column:id;PRIMARY_KEY;AUTO_INCREMENT"`
	ID      string `gorm:"column:channel_id;unique_index"`
	OwnerID string
	Name    string
}

func (channelModel) TableName() string {
	return "channel"
}

func (c *Channel) Flush() (err error) {
	db, err := connect(c)
	if err != nil {
		return
	}
	err = db.Where("channel_id!=?", "").Delete(channelModel{}).Error

	return
}

func (c *Channel) AutoMigrate() {
	var cm channelModel

	db, err := connect(c)
	if err != nil {
		return
	}

	db.AutoMigrate(&cm)

	db.Close()
}

func (c *Channel) Save(channel entity.Channel) (err error) {
	db, err := connect(c)
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()
	var model, cm channelModel
	copier.Copy(&cm, &channel)
	db.Where("channel_id=?", channel.ID).Take(&model)

	if model.MySQLID == 0 {
		if err = db.Create(&cm).Error; err != nil {
			return
		}
	} else {
		cm.MySQLID = model.MySQLID
		if err = db.Model(&cm).Update(&cm).Error; err != nil {
			return
		}
	}

	return
}

func (c *Channel) Get(filter entity.FilterChannel) (channels []entity.Channel) {
	var cm []channelModel
	db, err := connect(c)
	if err != nil {
		return
	}
	defer db.Close()
	limit := filterLimit(filter.Limit)
	if filter.ID != nil {
		db = db.Where("channel_id=?", *filter.ID)
	}
	if filter.Name != nil {
		db = db.Where("name LIKE ?", "%"+*filter.Name+"%")
	}
	db.Limit(limit).Find(&cm)

	copier.Copy(&channels, &cm)
	return
}
