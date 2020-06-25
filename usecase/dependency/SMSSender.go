package dependency

type SMSSender interface {
	Send(phoneNumber string, message string) (err error)
}
