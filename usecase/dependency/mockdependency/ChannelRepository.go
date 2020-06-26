package mockdependency

import (
	"github.com/ariefsam/go-chat/entity"
	"github.com/stretchr/testify/mock"
)

type ChannelRepository struct {
	mock.Mock
}

func (m *ChannelRepository) Get(filter entity.FilterChannel) (channels []entity.Channel) {
	args := m.Called(filter)
	return args.Get(0).([]entity.Channel)
}

func (m *ChannelRepository) Save(channel entity.Channel) error {
	args := m.Called(channel)
	return args.Error(0)
}
