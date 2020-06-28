package repository

import (
	"github.com/ariefsam/go-chat/entity"
	"github.com/jinzhu/copier"
)

type LoginVerification struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
}

type loginVerificationModel struct {
	MySQLID          int    `gorm:"column:id;PRIMARY_KEY;AUTO_INCREMENT"`
	ID               string `gorm:"column:login_verification_id;unique_index"`
	UserID           string
	PhoneNumber      string
	DeviceID         string
	VerificationCode string
	ExpiredTimestamp int64
}

func (loginVerificationModel) TableName() string {
	return "login_verification"
}

func (c *LoginVerification) Flush() (err error) {
	db, err := connect(c)
	if err != nil {
		return
	}
	err = db.Where("login_verification_id!=?", "").Delete(loginVerificationModel{}).Error

	return
}

func (c *LoginVerification) AutoMigrate() {

	var lvModel loginVerificationModel

	db, err := connect(c)
	if err != nil {
		return
	}

	err = db.AutoMigrate(&lvModel).Error
	if err != nil {
		return
	}

	db.Close()
}

func (r *LoginVerification) Save(lv entity.LoginVerification) (err error) {
	db, err := connect(r)
	if err != nil {
		return
	}
	defer db.Close()
	var model, lvModel loginVerificationModel
	copier.Copy(&lvModel, &lv)
	db.Where("login_verification_id=?", lv.ID).Take(&model)

	if model.MySQLID == 0 {
		if err = db.Create(&lvModel).Error; err != nil {
			return
		}
	} else {
		lvModel.MySQLID = model.MySQLID
		if err = db.Model(&lvModel).Update(&lvModel).Error; err != nil {
			return
		}
	}
	return
}

func (r *LoginVerification) Get(phoneNumber string, deviceID string, validBefore int64, verificationCode *string) (loginVerifications []entity.LoginVerification) {
	db, err := connect(r)
	if err != nil {
		return
	}
	defer db.Close()

	db = db.Where("phone_number=? AND device_id=? AND expired_timestamp>?", phoneNumber, deviceID, validBefore)

	if verificationCode != nil {
		db = db.Where("verification_code=?", *verificationCode)
	}
	var lvs []loginVerificationModel
	db.Find(&lvs)
	copier.Copy(&loginVerifications, &lvs)
	return
}
