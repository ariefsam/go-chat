package mockdependency

import "github.com/stretchr/testify/mock"

type IDGenerator struct {
	mock.Mock
}

func (m *IDGenerator) Generate() string {
	args := m.Called()
	return args.String(0)
}

func (m *IDGenerator) GenerateNumberCode(length int) string {
	args := m.Called(length)
	return args.String(0)
}
