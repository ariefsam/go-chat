package mockdependency

import "github.com/stretchr/testify/mock"

type SMSSender struct {
	mock.Mock
}

func (m *SMSSender) Send(phoneNumber string, message string) (err error) {
	args := m.Called(phoneNumber, message)
	err = args.Error(0)
	return
}
