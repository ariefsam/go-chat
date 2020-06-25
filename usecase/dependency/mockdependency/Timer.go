package mockdependency

import "github.com/stretchr/testify/mock"

type Timer struct {
	mock.Mock
}

func (m *Timer) CurrentTimestamp() int64 {
	args := m.Called()
	return args.Get(0).(int64)
}
