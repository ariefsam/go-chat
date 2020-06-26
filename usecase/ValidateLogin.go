package usecase

func (u *Usecase) ValidateLogin(phoneNumber string, deviceID string, verificationCode string) (isValid bool) {
	verificationLogin := u.LoginVerificationRepository.Get(phoneNumber, deviceID, u.Timer.CurrentTimestamp(), &verificationCode)
	if len(verificationLogin) > 0 {
		isValid = true
	}
	return
}
