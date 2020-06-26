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

	phoneNumber := "62852132373"
	deviceID := "xxxdevid"
	verificationCode := "778899"
	validateBefore := int64(15000)
	currentTime := int64(14000)
	mockTimer.On("CurrentTimestamp").Return(currentTime)

	t.Run("login verification found", func(t *testing.T) {
		expectedLoginVerification := entity.LoginVerification{
			ID:               "xxaa",
			PhoneNumber:      phoneNumber,
			DeviceID:         deviceID,
			ExpiredTimestamp: validateBefore,
			VerificationCode: verificationCode,
		}
		mockLoginVerification.On("Get", phoneNumber, deviceID, currentTime, &verificationCode).Return([]entity.LoginVerification{expectedLoginVerification}).Once()
		isValid := u.ValidateLogin(phoneNumber, deviceID, verificationCode)
		assert.True(t, isValid)
	})

	t.Run("login verification not found", func(t *testing.T) {
		mockLoginVerification.On("Get", phoneNumber, deviceID, currentTime, &verificationCode).Return([]entity.LoginVerification{}).Once()
		isValid := u.ValidateLogin(phoneNumber, deviceID, verificationCode)
		assert.False(t, isValid)
	})

}
