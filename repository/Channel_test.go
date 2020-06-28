package repository_test

import (
	"testing"

	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/assert"

	"github.com/ariefsam/go-chat/configuration"
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
}
