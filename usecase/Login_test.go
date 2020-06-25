package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ariefsam/go-chat/entity"
	"github.com/ariefsam/go-chat/usecase/dependency/mockdependency"

	"github.com/ariefsam/go-chat/usecase"
)

var mockUserRepository mockdependency.UserRepository
var mockIDGenerator mockdependency.IDGenerator
var mockSMSSender mockdependency.SMSSender

func TestLogin(t *testing.T) {
	var u usecase.Usecase
	mockUserRepository = mockdependency.UserRepository{}
	u.UserRepository = &mockUserRepository
	u.IDGenerator = &mockIDGenerator
	u.SMSSender = &mockSMSSender

	phoneNumber := "62852123456"
	deviceID := "xxdeviceid"
	t.Run("New user", func(t *testing.T) {

		filter := entity.FilterUser{
			PhoneNumber: &phoneNumber,
		}
		mockUserRepository.On("Get", filter).Return([]entity.User{})

		generatedID := "id001"
		mockIDGenerator.On("Generate").Return(generatedID)

		expectedUserToSave := entity.User{
			ID:          generatedID,
			PhoneNumber: phoneNumber,
		}

		mockUserRepository.On("Save", expectedUserToSave).Return(nil)

		// mockSMSSender.On("Send", phoneNumber, message).Return(nil)

		err := u.LoginBySMS(phoneNumber, deviceID)
		mockUserRepository.AssertCalled(t, "Get", filter)
		mockUserRepository.AssertCalled(t, "Save", expectedUserToSave)
		assert.NoError(t, err)
	})
}
