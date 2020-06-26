package usecase

import "github.com/ariefsam/go-chat/entity"

func (u *Usecase) GenerateToken(user entity.User) (token string) {
	token = u.TokenUserService.Create(user)
	return
}
