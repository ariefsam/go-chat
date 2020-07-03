package usecase

import "github.com/ariefsam/go-chat/entity"

func (u *Usecase) ValidateLogin(phoneNumber string, deviceID string, verificationCode string) (isValid bool, user entity.User) {
	verificationLogin := u.LoginVerificationRepository.Get(phoneNumber, deviceID, u.Timer.CurrentTimestamp(), &verificationCode)
	if len(verificationLogin) > 0 {
		isValid = true
		filter := entity.FilterUser{
			UserID: &verificationLogin[0].UserID,
		}
		users := u.UserRepository.Get(filter)
		if len(users) > 0 {
			users[0].IsPhoneVerified = true
			u.UserRepository.Save(users[0])
			user = users[0]
		}
	}
	return
}
