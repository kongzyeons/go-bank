package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateReq(t *testing.T) {
	var req *interface{}
	ValidateReq(req)

	type Req struct {
		Required        string `json:"name" validate:"required"`
		Dd              string `json:"-"`
		Required_unless string `json:"required_unless" validate:"required_unless"`
		ObjectID        string `validate:"mongodb"`
		MinItems        int    `validate:"min=2"`
		AlphaNum        string `validate:"alphanum"`
		EqualTo         int    `validate:"eq=10"`
		NotEqualTo      int    `validate:"ne=0"`
		GreaterThan     int    `validate:"gt=3"`
		GreaterOrEqual  int    `validate:"gte=5"`
		LessThan        int    `validate:"lt=0"`
		LessOrEqual     int    `validate:"lte=-1"`
		OneOfField      string `validate:"oneof=option1 option2 option3"`
		EmailField      string `validate:"email"`
	}
	ValidateReq(&Req{})

}

// Sample struct for testing
type SampleStruct struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// Test GetField function
func TestGetField(t *testing.T) {
	sample := SampleStruct{}

	t.Run("Field Found", func(t *testing.T) {
		field, err := GetField("name", "json", sample)
		assert.Nil(t, err)
		assert.Equal(t, "Name", field.Name)
	})

	t.Run("Field Not Found", func(t *testing.T) {
		_, err := GetField("unknown", "json", sample)
		assert.EqualError(t, err, "not found filter")
	})

	t.Run("Not a Struct", func(t *testing.T) {
		_, err := GetField("name", "json", 123)
		assert.EqualError(t, err, "no type struct")
	})
}
