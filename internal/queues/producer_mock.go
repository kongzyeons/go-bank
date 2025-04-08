package queues

import (
	"github.com/kongzyeons/go-bank/internal/models/events"
	"github.com/stretchr/testify/mock"
)

type eventProducerMock struct {
	mock.Mock
}

func NewEventProducerMock() *eventProducerMock {
	return &eventProducerMock{}
}

func (m *eventProducerMock) Produce(event events.Event) error {
	args := m.Called(event)
	return args.Error(0)
}
