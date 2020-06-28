package repository

import (
	"log"

	"github.com/ariefsam/go-chat/entity"
	"github.com/jinzhu/copier"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

}

type Chat struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
}

type chatModel struct {
	MySQLID    int    `gorm:"column:id;PRIMARY_KEY;AUTO_INCREMENT"`
	ID         string `gorm:"column:chat_id;unique_index"`
	SenderID   string
	ReceiverID string
	ChatType   string
	Message    string `gorm:"type:text"`
	Timestamp  int64
}

func (chatModel) TableName() string {
	return "chat"
}

func (c *Chat) Save(chat entity.Chat) (err error) {

	var cm, temp chatModel
	copier.Copy(&cm, &chat)

	db, err := connect(c)
	defer db.Close()
	if err != nil {
		return
	}

	db.Where("chat_id=?", chat.ID).Take(&temp)

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

func (c *Chat) Get(filter entity.FilterChat) (chats []entity.Chat) {
	db, err := connect(c)
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
	if filter.BeforeID != nil {
		var chat chatModel

		db.Where("chat_id=?", *filter.BeforeID).First(&chat)
		if chat.MySQLID != 0 {
			db = db.Where("id<?", chat.MySQLID)
		}
	}

	if filter.ReceiverID != nil {
		db = db.Where("receiver_id=?", *filter.ReceiverID)
	}

	var chatModels []chatModel

	db.Order("id desc").Limit(limit).Find(&chatModels)

	reverseSlice(chatModels)
	copier.Copy(&chats, &chatModels)

	return
}

func (c *Chat) Flush() (err error) {
	db, err := connect(c)
	if err != nil {
		return
	}
	err = db.Where("chat_id!=?", "").Delete(chatModel{}).Error

	return
}

func (c *Chat) AutoMigrate() {
	var chatModel chatModel

	db, err := connect(c)
	if err != nil {
		return
	}

	db.AutoMigrate(&chatModel)

	db.Close()
}
