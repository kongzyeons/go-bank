package line

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLineAPIMock(t *testing.T) {
	api := NewLineAPIMock()

	expectedMessage := "Hello, LINE!"
	api.On("SendMessage", expectedMessage).Return(nil)

	err := api.SendMessage(expectedMessage)

	assert.NoError(t, err)
	api.AssertExpectations(t)
}
