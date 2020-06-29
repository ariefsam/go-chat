package repository_test

import (
	"testing"

	"github.com/ariefsam/go-chat/configuration"
	"github.com/jinzhu/copier"

	"github.com/stretchr/testify/assert"

	"github.com/ariefsam/go-chat/entity"
	"github.com/ariefsam/go-chat/repository"
)

func TestUser(t *testing.T) {
	var userRepository repository.User
	copier.Copy(&userRepository, &configuration.Config.MySQL)
	userRepository.AutoMigrate()
	userRepository.Flush()
	user := entity.User{
		ID:          "id001",
		Name:        "xx1",
		PhoneNumber: "623232",
	}
	err := userRepository.Save(user)
	assert.NoError(t, err)

	filter := entity.FilterUser{
		UserID: &user.ID,
	}
	users := userRepository.Get(filter)
	assert.True(t, len(users) == 1)
	if len(users) == 1 {
		assert.Equal(t, user, users[0])
	}
}
