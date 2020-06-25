package mockdependency

import (
	"github.com/ariefsam/go-chat/entity"
	"github.com/stretchr/testify/mock"
)

type UserRepository struct {
	mock.Mock
}

func (m *UserRepository) Get(filter entity.FilterUser) (users []entity.User) {
	args := m.Called(filter)
	users = args.Get(0).([]entity.User)
	return
}

func (m *UserRepository) Save(user entity.User) (err error) {
	args := m.Called(user)
	err = args.Error(0)
	return
}
