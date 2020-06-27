package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ariefsam/go-chat/usecase/dependency/mockdependency"

	"github.com/ariefsam/go-chat/entity"
	"github.com/ariefsam/go-chat/usecase"
)

func TestGetChannel(t *testing.T) {
	var u usecase.Usecase
	dummyChannelResults := []entity.Channel{
		entity.Channel{
			ID:      "c001",
			Name:    "Channel Testing",
			OwnerID: "userid0001",
		},
		entity.Channel{
			ID:      "c002",
			Name:    "Channel Testing 2",
			OwnerID: "userid0001",
		},
	}
	var mockChannelRepository mockdependency.ChannelRepository
	u.ChannelRepository = &mockChannelRepository

	searchName := "channel testing"
	filter := entity.FilterChannel{
		Name: &searchName,
	}
	mockChannelRepository.On("Get", filter).Return(dummyChannelResults)
	channels := u.GetChannel(filter)
	assert.Equal(t, dummyChannelResults, channels)
}
