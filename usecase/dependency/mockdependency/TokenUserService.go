package mockdependency

import (
	"github.com/ariefsam/go-chat/entity"
	"github.com/stretchr/testify/mock"
)

type TokenUserService struct {
	mock.Mock
}

func (m *TokenUserService) Create(user entity.User) (token string) {
	args := m.Called(user)
	token = args.String(0)
	return
}
func (m *TokenUserService) Parse(token string) (isValid bool, user entity.User) {
	args := m.Called(token)
	isValid = args.Bool(0)
	user = args.Get(1).(entity.User)
	return
}
