package usecase

import "github.com/ariefsam/go-chat/entity"

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
	}
	return
}
