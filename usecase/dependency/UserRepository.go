package dependency

import "github.com/ariefsam/go-chat/entity"

type UserRepository interface {
	Get(filter entity.FilterUser) []entity.User
	Save(user entity.User) (err error)
}
