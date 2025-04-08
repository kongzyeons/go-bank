package queues

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewEventProducerMock(t *testing.T) {
	producer := NewEventProducerMock()

	producer.On("Produce", mock.Anything).Return(nil)

	err := producer.Produce(nil)
	assert.NoError(t, err)
}
