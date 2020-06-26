package dependency

import "github.com/ariefsam/go-chat/entity"

type TokenUserService interface {
	Create(user entity.User) (token string)
	Parse(token string) (isValid bool, user entity.User)
}
