package mockdependency

import (
	"github.com/ariefsam/go-chat/entity"
	"github.com/stretchr/testify/mock"
)

type LoginVerificationRepository struct {
	mock.Mock
}

func (m *LoginVerificationRepository) Get(phoneNumber string, deviceID string, validBefore int64, verificationCode *string) (loginVerifications []entity.LoginVerification) {
	args := m.Called(phoneNumber, deviceID, validBefore, verificationCode)
	loginVerifications = args.Get(0).([]entity.LoginVerification)
	return
}

func (m *LoginVerificationRepository) Save(loginVerification entity.LoginVerification) (err error) {
	args := m.Called(loginVerification)
	err = args.Error(0)
	return
}
