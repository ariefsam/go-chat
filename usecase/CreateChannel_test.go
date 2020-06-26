package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ariefsam/go-chat/entity"
	"github.com/ariefsam/go-chat/usecase"
	"github.com/ariefsam/go-chat/usecase/dependency/mockdependency"
)

func TestCreateChannel(t *testing.T) {
	var u usecase.Usecase
	var mockIDGenerator mockdependency.IDGenerator
	u.IDGenerator = &mockIDGenerator
	var mockChannelRepository mockdependency.ChannelRepository
	u.ChannelRepository = &mockChannelRepository

	channel := entity.Channel{
		Name:    "Channel Testing",
		OwnerID: "userid0001",
	}

	expectedChannelID := "channelid0001"
	mockIDGenerator.On("Generate").Return(expectedChannelID).Once()
	expectedChannelToSave := entity.Channel{
		ID:      expectedChannelID,
		Name:    "Channel Testing",
		OwnerID: "userid0001",
	}

	mockChannelRepository.On("Save", expectedChannelToSave).Return(nil)
	savedChannel, err := u.CreateChannel(channel)
	mockChannelRepository.AssertCalled(t, "Save", expectedChannelToSave)
	assert.NoError(t, err)
	assert.Equal(t, expectedChannelToSave, savedChannel)

}
