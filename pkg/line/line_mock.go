package line

import "github.com/stretchr/testify/mock"

type lineAPIMock struct {
	mock.Mock
}

func NewLineAPIMock() *lineAPIMock {
	return &lineAPIMock{}
}

func (m *lineAPIMock) SendMessage(message string) error {
	args := m.Called(message)
	return args.Error(0)
}
