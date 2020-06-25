package usecase

import (
	"github.com/ariefsam/go-chat/entity"
)

func (u *Usecase) LoginBySMS(phoneNumber string, deviceID string) (err error) {
	var user entity.User
	filter := entity.FilterUser{
		PhoneNumber: &phoneNumber,
	}
	checkUsers := u.UserRepository.Get(filter)

	if len(checkUsers) == 0 {
		user = entity.User{
			ID:          u.IDGenerator.Generate(),
			PhoneNumber: phoneNumber,
		}
		err = u.UserRepository.Save(user)
		if err != nil {
			return
		}
	} else {
		user = checkUsers[0]
	}
	lv := entity.LoginVerification{
		DeviceID:         deviceID,
		ID:               u.IDGenerator.Generate(),
		UserID:           user.ID,
		ExpiredTimestamp: u.Timer.CurrentTimestamp() + 300,
		VerificationCode: u.IDGenerator.GenerateNumberCode(6),
		PhoneNumber:      phoneNumber,
	}
	message := lv.VerificationCode + " is your code verification."
	err = u.SMSSender.Send(phoneNumber, message)
	if err != nil {
		return
	}

	err = u.LoginVerificationRepository.Save(lv)
	if err != nil {
		return
	}

	return
}
