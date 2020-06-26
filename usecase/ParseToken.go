package usecase

import "github.com/ariefsam/go-chat/entity"

func (u *Usecase) ParseToken(token string) (isValid bool, user entity.User) {
	isValid, user = u.TokenUserService.Parse(token)
	return
}
