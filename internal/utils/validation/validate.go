package validation

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateReq[T any](req *T) map[string]string {

	valMap := make(map[string]string)
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	if req == nil {
		return valMap
	}
	if err := validate.Struct(req); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			// Iterate over validation errors
			for _, validationError := range validationErrors {
				// Extract the field name causing the error
				fieldName := validationError.Field()
				switch validationError.Tag() {
				case "required":
					valMap[fieldName] = fmt.Sprintf("%s is requied", fieldName)
				case "required_unless":
					valMap[fieldName] = fmt.Sprintf("%s is requied", fieldName)
				case "required_if_with_min":
					valMap[fieldName] = fmt.Sprintf("%s must be at least one", fieldName)
				case "required_if_with_min_sum":
					valMap[fieldName] = fmt.Sprintf("%s must be at least one", fieldName)
				case "mongodb":
					valMap[fieldName] = fmt.Sprintf("%s must be objectID format", fieldName)
				case "date_format":
					valMap[fieldName] = fmt.Sprintf("%s date format must be %s", fieldName, validationError.Param())
				case "max":
					valMap[fieldName] = fmt.Sprintf("%s can be contain only %s items", fieldName, validationError.Param())
				case "min":
					valMap[fieldName] = fmt.Sprintf("%s must be at least %s items", fieldName, validationError.Param())
				case "alphanum":
					valMap[fieldName] = fmt.Sprintf("%s must be letters and numbers", fieldName)
				case "eq":
					valMap[fieldName] = fmt.Sprintf("%s must be equal to %s", fieldName, validationError.Param())
				case "ne":
					valMap[fieldName] = fmt.Sprintf("%s must be not equal to %s", fieldName, validationError.Param())
				case "gt":
					valMap[fieldName] = fmt.Sprintf("%s must be greater than to %s", fieldName, validationError.Param())
				case "gte":
					valMap[fieldName] = fmt.Sprintf("%s must be greater than or equal to %s", fieldName, validationError.Param())
				case "lt":
					valMap[fieldName] = fmt.Sprintf("%s must be less than to %s", fieldName, validationError.Param())
				case "lte":
					valMap[fieldName] = fmt.Sprintf("%s must be less than or equal to %s", fieldName, validationError.Param())
				case "oneof":
					valMap[fieldName] = fmt.Sprintf("%s must be %s", fieldName, validationError.Param())
				case "email":
					valMap[fieldName] = fmt.Sprintf("%s must be email format", fieldName)
				case "required_string":
					valMap[fieldName] = fmt.Sprintf("%s is requied", fieldName)
				case "array_max":
					valMap[fieldName] = fmt.Sprintf("%s can be contain only %s items", fieldName, validationError.Param())
				case "string_max":
					valMap[fieldName] = fmt.Sprintf("%s can be contain only %s characters", fieldName, validationError.Param())
				default:
					valMap[fieldName] = validationError.Error()
				}
			}
		}
	}
	return valMap
}

func GetField(tag, key string, s interface{}) (field reflect.StructField, err error) {
	rt := reflect.TypeOf(s)
	if rt.Kind() != reflect.Struct {
		return field, errors.New("no type struct")
	}
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		v := strings.Split(field.Tag.Get(key), ",")[0] // use split to ignore tag "options" like omitempty, etc.
		if v == tag {
			return field, nil
		}
	}
	return field, errors.New("not found filter")
}
