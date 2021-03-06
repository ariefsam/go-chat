package dependency

import "github.com/ariefsam/go-chat/entity"

type LoginVerificationRepository interface {
	Get(phoneNumber string, deviceID string, validBefore int64, verificationCode *string) (loginVerifications []entity.LoginVerification)
	Save(loginVerification entity.LoginVerification) (err error)
}
