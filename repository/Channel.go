package repository

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
