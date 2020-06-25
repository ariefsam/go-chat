package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ariefsam/go-chat/entity"
	"github.com/ariefsam/go-chat/usecase/dependency/mockdependency"

	"github.com/ariefsam/go-chat/usecase"
)

func TestLogin(t *testing.T) {
	var u usecase.Usecase
	var mockUserRepository mockdependency.UserRepository
	var mockIDGenerator mockdependency.IDGenerator
	var mockSMSSender mockdependency.SMSSender
	var mockTimer mockdependency.Timer
	var mockLoginVerificationRepository mockdependency.LoginVerificationRepository

	u.UserRepository = &mockUserRepository
	u.IDGenerator = &mockIDGenerator
	u.SMSSender = &mockSMSSender
	u.LoginVerificationRepository = &mockLoginVerificationRepository
	u.Timer = &mockTimer

	phoneNumber := "62852123456"
	deviceID := "xxdeviceid"
	verificationCode := "543322"
	t.Run("New user", func(t *testing.T) {
		filter := entity.FilterUser{
			PhoneNumber: &phoneNumber,
		}
		mockUserRepository.On("Get", filter).Return([]entity.User{})

		generatedUserID := "id001"
		generatedLoginVerificationID := "lp001"
		mockIDGenerator.On("Generate").Return(generatedUserID).Once()
		mockIDGenerator.On("Generate").Return(generatedLoginVerificationID).Once()
		mockIDGenerator.On("GenerateNumberCode", 6).Return(verificationCode)

		expectedMessage := verificationCode + " is your code verification."
		expectedUserToSave := entity.User{
			ID:          generatedUserID,
			PhoneNumber: phoneNumber,
		}

		mockUserRepository.On("Save", expectedUserToSave).Return(nil)

		mockSMSSender.On("Send", phoneNumber, expectedMessage).Return(nil)

		var expectedCurrentTimestamp int64
		expectedCurrentTimestamp = 100
		mockTimer.On("CurrentTimestamp").Return(expectedCurrentTimestamp)
		expectedExpiredTimestamp := expectedCurrentTimestamp + 300

		expectedLoginVerificationToSave := entity.LoginVerification{
			DeviceID:         deviceID,
			PhoneNumber:      phoneNumber,
			UserID:           generatedUserID,
			ID:               generatedLoginVerificationID,
			ExpiredTimestamp: expectedExpiredTimestamp,
			VerificationCode: verificationCode,
		}
		mockLoginVerificationRepository.On("Save", expectedLoginVerificationToSave).Return(nil)

		err := u.LoginBySMS(phoneNumber, deviceID)
		mockUserRepository.AssertCalled(t, "Get", filter)
		mockUserRepository.AssertCalled(t, "Save", expectedUserToSave)
		mockLoginVerificationRepository.AssertCalled(t, "Save", expectedLoginVerificationToSave)
		mockSMSSender.AssertCalled(t, "Send", phoneNumber, expectedMessage)

		assert.NoError(t, err)
	})
}
