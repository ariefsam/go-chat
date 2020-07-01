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

	t.Run("Test search by phone number", func(t *testing.T) {
		user = entity.User{
			ID:          "id002",
			Name:        "xx2",
			PhoneNumber: "12623232",
		}
		err = userRepository.Save(user)
		filter = entity.FilterUser{
			PhoneNumber: &user.PhoneNumber,
		}
		getUsers := userRepository.Get(filter)
		assert.Equal(t, 1, len(getUsers))
		if len(getUsers) == 1 {
			assert.Equal(t, user, getUsers[0])
		}

	})

	t.Run("Error if same phone number saved from different id", func(t *testing.T) {
		user = entity.User{
			ID:          "id003",
			Name:        "xx2",
			PhoneNumber: "12623232",
		}
		err = userRepository.Save(user)
		assert.Error(t, err)

	})
}
