package mockdependency

import (
	"github.com/ariefsam/go-chat/entity"
	"github.com/stretchr/testify/mock"
)

type ChatRepository struct {
	mock.Mock
}

func (m *ChatRepository) Get(filter entity.FilterChat) (chats []entity.Chat) {
	args := m.Called(filter)
	chats = args.Get(0).([]entity.Chat)
	return
}

func (m *ChatRepository) Save(chat entity.Chat) (err error) {
	args := m.Called(chat)
	err = args.Error(0)
	return
}
