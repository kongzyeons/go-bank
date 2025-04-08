package response

import (
	"testing"
)

func TestResponse(t *testing.T) {
	Ok[any](nil)
	BadRequest[any]("test")
	ValidationFailed[any](map[string]string{})
	Unauthorized[any]("test")
	Notfound[any]("test")
	InternalServerError[any](nil, "")
	ErrorWithCode[any](nil, "", 0)
}
