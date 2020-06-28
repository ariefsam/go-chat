package repository_test

import (
	"testing"

	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/assert"

	"github.com/ariefsam/go-chat/configuration"
	"github.com/ariefsam/go-chat/entity"
	"github.com/ariefsam/go-chat/repository"
)

func TestChannelRepository(t *testing.T) {
	var channelRepository repository.Channel
	copier.Copy(&channelRepository, configuration.Config.MySQL)
	channelRepository.AutoMigrate()

	t.Run("Flush data", func(t *testing.T) {
		err := channelRepository.Flush()
		assert.NoError(t, err)
	})
	expectedChannel := entity.Channel{
		ID:      "id001",
		Name:    "Channel Testing",
		OwnerID: "userid0001",
	}
	t.Run("Save Normal Data", func(t *testing.T) {

		err := channelRepository.Save(expectedChannel)
		assert.NoError(t, err)
	})

	t.Run("Get Normal Data", func(t *testing.T) {
		filter := entity.FilterChannel{ID: &expectedChannel.ID}
		channels := channelRepository.Get(filter)
		assert.Equal(t, len(channels), 1)
		if len(channels) == 1 {
			assert.Equal(t, expectedChannel, channels[0])
		}
	})

	t.Run("Save Edited Data", func(t *testing.T) {
		expectedChannel.Name = "Name edited"
		err := channelRepository.Save(expectedChannel)
		assert.NoError(t, err)
	})

	t.Run("Get Edited Data", func(t *testing.T) {
		filter := entity.FilterChannel{ID: &expectedChannel.ID}
		channels := channelRepository.Get(filter)
		assert.Equal(t, len(channels), 1)
		if len(channels) == 1 {
			assert.Equal(t, expectedChannel, channels[0])
		}
	})

	expectedChannel2 := entity.Channel{
		ID:      "id002",
		Name:    "Channel Ada",
		OwnerID: "userid0001",
	}

	t.Run("Save Normal Data", func(t *testing.T) {
		err := channelRepository.Save(expectedChannel2)
		assert.NoError(t, err)
	})

	expectedChannel3 := entity.Channel{
		ID:      "id003",
		Name:    "Channel Ada Two",
		OwnerID: "userid0001",
	}

	t.Run("Save Normal Data", func(t *testing.T) {
		err := channelRepository.Save(expectedChannel3)
		assert.NoError(t, err)
	})

	t.Run("Find by name", func(t *testing.T) {
		search := "channel ada"
		filter := entity.FilterChannel{Name: &search}
		channels := channelRepository.Get(filter)
		expectedChannels := []entity.Channel{
			expectedChannel2,
			expectedChannel3,
		}
		assert.Equal(t, expectedChannels, channels)
	})
}
