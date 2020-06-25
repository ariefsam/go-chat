package entity

type LoginVerification struct {
	ID               string
	UserID           string
	PhoneNumber      string
	DeviceID         string
	VerificationCode string
	ExpiredTimestamp int64
}
