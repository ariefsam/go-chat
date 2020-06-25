package entity

type LoginVerification struct {
	UserID           string
	PhoneNumber      string
	DeviceID         string
	VerificationCode string
}
