package usecase_test

import (
	"testing"

	"github.com/ariefsam/go-chat/entity"
	"github.com/ariefsam/go-chat/usecase"
	"github.com/ariefsam/go-chat/usecase/dependency/mockdependency"
	"github.com/stretchr/testify/assert"
)

func TestValidateLogin(t *testing.T) {
	var u usecase.Usecase
	var mockLoginVerification mockdependency.LoginVerificationRepository

	u.LoginVerificationRepository = &mockLoginVerification
	var mockTimer mockdependency.Timer
	u.Timer = &mockTimer
	var mockUserRepository mockdependency.UserRepository
	u.UserRepository = &mockUserRepository

	phoneNumber := "62852132373"
	deviceID := "xxxdevid"
	verificationCode := "778899"
	validateBefore := int64(15000)
	currentTime := int64(14000)
	userID := "001"
	mockTimer.On("CurrentTimestamp").Return(currentTime)
	expectedUser := entity.User{
		ID:              userID,
		PhoneNumber:     phoneNumber,
		IsPhoneVerified: true,
	}

	t.Run("login verification found", func(t *testing.T) {
		expectedLoginVerification := entity.LoginVerification{
			ID:               "xxaa",
			PhoneNumber:      phoneNumber,
			DeviceID:         deviceID,
			ExpiredTimestamp: validateBefore,
			VerificationCode: verificationCode,
			UserID:           userID,
		}
		mockLoginVerification.On("Get", phoneNumber, deviceID, currentTime, &verificationCode).Return([]entity.LoginVerification{expectedLoginVerification}).Once()

		filterUser := entity.FilterUser{
			UserID: &expectedLoginVerification.UserID,
		}

		mockUserRepository.On("Get", filterUser).Return([]entity.User{expectedUser})
		mockUserRepository.On("Save", expectedUser).Return(nil)

		isValid, user := u.ValidateLogin(phoneNumber, deviceID, verificationCode)
		assert.True(t, isValid)
		assert.Equal(t, expectedUser, user)
	})

	t.Run("login verification not found", func(t *testing.T) {
		mockLoginVerification.On("Get", phoneNumber, deviceID, currentTime, &verificationCode).Return([]entity.LoginVerification{}).Once()
		isValid, _ := u.ValidateLogin(phoneNumber, deviceID, verificationCode)
		assert.False(t, isValid)

	})

}
