package mockdependency

import (
	"github.com/ariefsam/go-chat/entity"
	"github.com/stretchr/testify/mock"
)

type LoginVerificationRepository struct {
	mock.Mock
}

func (m *LoginVerificationRepository) Get(phoneNumber string, deviceID string, validBefore int64, verificationCode *string) (loginVerifications []entity.LoginVerification) {
	m.Called(phoneNumber, deviceID, validBefore, verificationCode)
	return
}
