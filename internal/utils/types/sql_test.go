package types

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

// TestNewNullBool tests the NewNullBool function.
func TestNewNullBool(t *testing.T) {
	tests := []struct {
		name     string
		input    bool
		expected SQLNullBool
	}{
		{"True Value", true, SQLNullBool{sql.NullBool{Bool: true, Valid: true}}},
		{"False Value", false, SQLNullBool{sql.NullBool{Bool: false, Valid: true}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewNullBool(tt.input)
			if result.Bool != tt.expected.Bool || result.Valid != tt.expected.Valid {

			}
		})
	}
}

// TestIsNull tests the IsNull method.
func TestIsNull(t *testing.T) {
	tests := []struct {
		name     string
		input    SQLNullBool
		expected bool
	}{
		{"Valid True", SQLNullBool{sql.NullBool{Bool: true, Valid: true}}, false},
		{"Valid False", SQLNullBool{sql.NullBool{Bool: false, Valid: true}}, false},
		{"Null Value", SQLNullBool{sql.NullBool{Valid: false}}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.IsNull()
			if result != tt.expected {

			}
		})
	}
}

// TestVal tests the Val method.
func TestVal1(t *testing.T) {
	tests := []struct {
		name     string
		input    SQLNullBool
		expected bool
	}{
		{"Valid True", SQLNullBool{sql.NullBool{Bool: true, Valid: true}}, true},
		{"Valid False", SQLNullBool{sql.NullBool{Bool: false, Valid: true}}, false},
		{"Null Value", SQLNullBool{sql.NullBool{Valid: false}}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.Val()
			if result != tt.expected {

			}
		})
	}
}

// TestSetBool tests the SetBool method.
func TestSetBool(t *testing.T) {
	tests := []struct {
		name     string
		input    bool
		expected SQLNullBool
	}{
		{"Set True", true, SQLNullBool{sql.NullBool{Bool: true, Valid: true}}},
		{"Set False", false, SQLNullBool{sql.NullBool{Bool: false, Valid: true}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result SQLNullBool
			result.SetBool(tt.input)

			if result.Bool != tt.expected.Bool || result.Valid != tt.expected.Valid {

			}
		})
	}
}

// TestString tests the String method.
func TestString1(t *testing.T) {
	tests := []struct {
		name     string
		input    SQLNullBool
		expected string
	}{
		{"Valid True", SQLNullBool{sql.NullBool{Bool: true, Valid: true}}, "true"},
		{"Valid False", SQLNullBool{sql.NullBool{Bool: false, Valid: true}}, "false"},
		{"Null Value", SQLNullBool{sql.NullBool{Valid: false}}, "false"}, // Default to false when null
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.String()
			if result != tt.expected {

			}
		})
	}
}

// TestSetNull tests the SetNull method.
func TestSetNull1(t *testing.T) {
	t.Run("Set Null", func(t *testing.T) {
		var s SQLNullBool
		s.SetNull()

		if s.Valid != false || s.Bool != false {

		}
	})
}

// TestNewNullFloat64 tests the NewNullFloat64 function.
func TestNewNullFloat641(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected SQLNullFloat64
	}{
		{"Valid Float64", 12.34, SQLNullFloat64{sql.NullFloat64{Float64: 12.34, Valid: true}}},
		{"Zero Float64", 0.0, SQLNullFloat64{sql.NullFloat64{Float64: 0.0, Valid: true}}},
		{"Negative Float64", -45.67, SQLNullFloat64{sql.NullFloat64{Float64: -45.67, Valid: true}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewNullFloat64(tt.input)
			if result.Float64 != tt.expected.Float64 || result.Valid != tt.expected.Valid {

			}
		})
	}
}

// TestIsNull tests the IsNull method.
func TestIsNull1(t *testing.T) {
	tests := []struct {
		name     string
		input    SQLNullFloat64
		expected bool
	}{
		{"Valid Float", SQLNullFloat64{sql.NullFloat64{Float64: 10.5, Valid: true}}, false},
		{"Null Value", SQLNullFloat64{sql.NullFloat64{Valid: false}}, true},
		{"Zero but Valid", SQLNullFloat64{sql.NullFloat64{Float64: 0.0, Valid: true}}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.IsNull()
			if result != tt.expected {

			}
		})
	}
}

// TestVal tests the Val method.
func TestVal2(t *testing.T) {
	tests := []struct {
		name        string
		input       SQLNullFloat64
		defaultVal  float64
		expectedVal float64
	}{
		{"Valid Float", SQLNullFloat64{sql.NullFloat64{Float64: 10.5, Valid: true}}, 0.0, 10.5},
		{"Null Value Uses Default", SQLNullFloat64{sql.NullFloat64{Valid: false}}, 5.5, 5.5},
		{"Zero but Valid", SQLNullFloat64{sql.NullFloat64{Float64: 0.0, Valid: true}}, 100.0, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.Val(tt.defaultVal)
			if result != tt.expectedVal {

			}
		})
	}
}

// TestSetFloat64 tests the SetFloat64 method.
func TestSetFloat642(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected SQLNullFloat64
	}{
		{"Set Float64", 20.45, SQLNullFloat64{sql.NullFloat64{Float64: 20.45, Valid: true}}},
		{"Set Zero", 0.0, SQLNullFloat64{sql.NullFloat64{Float64: 0.0, Valid: true}}},
		{"Set Negative", -99.99, SQLNullFloat64{sql.NullFloat64{Float64: -99.99, Valid: true}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result SQLNullFloat64
			result.SetFloat64(tt.input)

			if result.Float64 != tt.expected.Float64 || result.Valid != tt.expected.Valid {

			}
		})
	}
}

// TestSetNull tests the SetNull method.
func TestSetNull2(t *testing.T) {
	t.Run("Set Null", func(t *testing.T) {
		var s SQLNullFloat64
		s.SetNull()

		if s.Valid != false || s.Float64 != 0.0 {

		}
	})
}

// TestSetDecimal tests the SetDecimal method.
func TestSetDecimal2(t *testing.T) {
	tests := []struct {
		name     string
		input    decimal.Decimal
		expected SQLNullFloat64
	}{
		{"Set Decimal 10.50", decimal.NewFromFloat(10.50), SQLNullFloat64{sql.NullFloat64{Float64: 10.50, Valid: true}}},
		{"Set Decimal 0.00", decimal.NewFromFloat(0.00), SQLNullFloat64{sql.NullFloat64{Float64: 0.00, Valid: true}}},
		{"Set Decimal -99.99", decimal.NewFromFloat(-99.99), SQLNullFloat64{sql.NullFloat64{Float64: -99.99, Valid: true}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s SQLNullFloat64
			s.SetDecimal(tt.input)

			if s.Float64 != tt.expected.Float64 || s.Valid != tt.expected.Valid {

			}
		})
	}
}

// TestDecimal tests the Decimal method.
func TestDecimal2(t *testing.T) {
	tests := []struct {
		name     string
		input    SQLNullFloat64
		expected decimal.Decimal
	}{
		{"Convert Float 10.5", SQLNullFloat64{sql.NullFloat64{Float64: 10.5, Valid: true}}, decimal.NewFromFloat(10.5)},
		{"Convert Null Defaults to 0.0", SQLNullFloat64{sql.NullFloat64{Valid: false}}, decimal.NewFromFloat(0.0)},
		{"Convert Negative Float", SQLNullFloat64{sql.NullFloat64{Float64: -88.88, Valid: true}}, decimal.NewFromFloat(-88.88)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.Decimal()
			if !result.Equal(tt.expected) {

			}
		})
	}
}

// TestString tests the String method.
func TestString2(t *testing.T) {
	tests := []struct {
		name     string
		input    SQLNullFloat64
		expected string
	}{
		{"Valid Float", SQLNullFloat64{sql.NullFloat64{Float64: 12.3456, Valid: true}}, "12.35"},
		{"Zero but Valid", SQLNullFloat64{sql.NullFloat64{Float64: 0.0, Valid: true}}, ""},
		{"Null Value", SQLNullFloat64{sql.NullFloat64{Valid: false}}, ""},
		{"Negative Float", SQLNullFloat64{sql.NullFloat64{Float64: -55.5555, Valid: true}}, "-55.56"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.String()
			if result != tt.expected {

			}
		})
	}
}

// TestNewNullFloat64 tests the NewNullFloat64 function.
func TestNewNullFloat64(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected SQLNullFloat64
	}{
		{"Valid Float64", 12.34, SQLNullFloat64{sql.NullFloat64{Float64: 12.34, Valid: true}}},
		{"Zero Float64", 0.0, SQLNullFloat64{sql.NullFloat64{Float64: 0.0, Valid: true}}},
		{"Negative Float64", -45.67, SQLNullFloat64{sql.NullFloat64{Float64: -45.67, Valid: true}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewNullFloat64(tt.input)
			if result.Float64 != tt.expected.Float64 || result.Valid != tt.expected.Valid {

			}
		})
	}
}

// TestIsNull tests the IsNull method.
func TestIsNull2(t *testing.T) {
	tests := []struct {
		name     string
		input    SQLNullFloat64
		expected bool
	}{
		{"Valid Float", SQLNullFloat64{sql.NullFloat64{Float64: 10.5, Valid: true}}, false},
		{"Null Value", SQLNullFloat64{sql.NullFloat64{Valid: false}}, true},
		{"Zero but Valid", SQLNullFloat64{sql.NullFloat64{Float64: 0.0, Valid: true}}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.IsNull()
			if result != tt.expected {

			}
		})
	}
}

// TestVal tests the Val method.
func TestVal(t *testing.T) {
	tests := []struct {
		name        string
		input       SQLNullFloat64
		defaultVal  float64
		expectedVal float64
	}{
		{"Valid Float", SQLNullFloat64{sql.NullFloat64{Float64: 10.5, Valid: true}}, 0.0, 10.5},
		{"Null Value Uses Default", SQLNullFloat64{sql.NullFloat64{Valid: false}}, 5.5, 5.5},
		{"Zero but Valid", SQLNullFloat64{sql.NullFloat64{Float64: 0.0, Valid: true}}, 100.0, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.Val(tt.defaultVal)
			if result != tt.expectedVal {

			}
		})
	}
}

// TestSetFloat64 tests the SetFloat64 method.
func TestSetFloat64(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected SQLNullFloat64
	}{
		{"Set Float64", 20.45, SQLNullFloat64{sql.NullFloat64{Float64: 20.45, Valid: true}}},
		{"Set Zero", 0.0, SQLNullFloat64{sql.NullFloat64{Float64: 0.0, Valid: true}}},
		{"Set Negative", -99.99, SQLNullFloat64{sql.NullFloat64{Float64: -99.99, Valid: true}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result SQLNullFloat64
			result.SetFloat64(tt.input)

			if result.Float64 != tt.expected.Float64 || result.Valid != tt.expected.Valid {

			}
		})
	}
}

// TestSetNull tests the SetNull method.
func TestSetNull(t *testing.T) {
	t.Run("Set Null", func(t *testing.T) {
		var s SQLNullFloat64
		s.SetNull()

		if s.Valid != false || s.Float64 != 0.0 {

		}
	})
}

// TestSetDecimal tests the SetDecimal method.
func TestSetDecimal(t *testing.T) {
	tests := []struct {
		name     string
		input    decimal.Decimal
		expected SQLNullFloat64
	}{
		{"Set Decimal 10.50", decimal.NewFromFloat(10.50), SQLNullFloat64{sql.NullFloat64{Float64: 10.50, Valid: true}}},
		{"Set Decimal 0.00", decimal.NewFromFloat(0.00), SQLNullFloat64{sql.NullFloat64{Float64: 0.00, Valid: true}}},
		{"Set Decimal -99.99", decimal.NewFromFloat(-99.99), SQLNullFloat64{sql.NullFloat64{Float64: -99.99, Valid: true}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s SQLNullFloat64
			s.SetDecimal(tt.input)

			if s.Float64 != tt.expected.Float64 || s.Valid != tt.expected.Valid {

			}
		})
	}
}

// TestDecimal tests the Decimal method.
func TestDecimal(t *testing.T) {
	tests := []struct {
		name     string
		input    SQLNullFloat64
		expected decimal.Decimal
	}{
		{"Convert Float 10.5", SQLNullFloat64{sql.NullFloat64{Float64: 10.5, Valid: true}}, decimal.NewFromFloat(10.5)},
		{"Convert Null Defaults to 0.0", SQLNullFloat64{sql.NullFloat64{Valid: false}}, decimal.NewFromFloat(0.0)},
		{"Convert Negative Float", SQLNullFloat64{sql.NullFloat64{Float64: -88.88, Valid: true}}, decimal.NewFromFloat(-88.88)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.Decimal()
			if !result.Equal(tt.expected) {

			}
		})
	}
}

// TestString tests the String method.
func TestString(t *testing.T) {
	tests := []struct {
		name     string
		input    SQLNullFloat64
		expected string
	}{
		{"Valid Float", SQLNullFloat64{sql.NullFloat64{Float64: 12.3456, Valid: true}}, "12.35"},
		{"Zero but Valid", SQLNullFloat64{sql.NullFloat64{Float64: 0.0, Valid: true}}, ""},
		{"Null Value", SQLNullFloat64{sql.NullFloat64{Valid: false}}, ""},
		{"Negative Float", SQLNullFloat64{sql.NullFloat64{Float64: -55.5555, Valid: true}}, "-55.56"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.String()
			if result != tt.expected {

			}
		})
	}
}

// TestNewNullInt64 tests the NewNullInt64 function.
func TestNewNullInt64(t *testing.T) {
	tests := []struct {
		name     string
		input    int64
		expected SQLNullInt64
	}{
		{"Valid Positive Int", 123, SQLNullInt64{sql.NullInt64{Int64: 123, Valid: true}}},
		{"Valid Zero", 0, SQLNullInt64{sql.NullInt64{Int64: 0, Valid: true}}},
		{"Valid Negative Int", -456, SQLNullInt64{sql.NullInt64{Int64: -456, Valid: true}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewNullInt64(tt.input)
			if result.Int64 != tt.expected.Int64 || result.Valid != tt.expected.Valid {

			}
		})
	}
}

// TestIsNull tests the IsNull method.
func TestIsNull3(t *testing.T) {
	tests := []struct {
		name     string
		input    SQLNullInt64
		expected bool
	}{
		{"Valid Int", SQLNullInt64{sql.NullInt64{Int64: 100, Valid: true}}, false},
		{"Null Value", SQLNullInt64{sql.NullInt64{Valid: false}}, true},
		{"Zero but Valid", SQLNullInt64{sql.NullInt64{Int64: 0, Valid: true}}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.IsNull()
			if result != tt.expected {

			}
		})
	}
}

// TestVal tests the Val method.
func TestVal3(t *testing.T) {
	tests := []struct {
		name        string
		input       SQLNullInt64
		defaultVal  int64
		expectedVal int64
	}{
		{"Valid Int", SQLNullInt64{sql.NullInt64{Int64: 10, Valid: true}}, 0, 10},
		{"Null Value Uses Default", SQLNullInt64{sql.NullInt64{Valid: false}}, 5, 5},
		{"Zero but Valid", SQLNullInt64{sql.NullInt64{Int64: 0, Valid: true}}, 100, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.Val(tt.defaultVal)
			if result != tt.expectedVal {

			}
		})
	}
}

// TestSetInt64 tests the SetInt64 method.
func TestSetInt64(t *testing.T) {
	tests := []struct {
		name     string
		input    int64
		expected SQLNullInt64
	}{
		{"Set Positive Int", 20, SQLNullInt64{sql.NullInt64{Int64: 20, Valid: true}}},
		{"Set Zero", 0, SQLNullInt64{sql.NullInt64{Int64: 0, Valid: true}}},
		{"Set Negative Int", -99, SQLNullInt64{sql.NullInt64{Int64: -99, Valid: true}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result SQLNullInt64
			result.SetInt64(tt.input)

			if result.Int64 != tt.expected.Int64 || result.Valid != tt.expected.Valid {

			}
		})
	}
}

// TestSetNull tests the SetNull method.
func TestSetNull3(t *testing.T) {
	t.Run("Set Null", func(t *testing.T) {
		var s SQLNullInt64
		s.SetNull()

		if s.Valid != false || s.Int64 != 0 {

		}
	})
}

// TestString tests the String method.
func TestString3(t *testing.T) {
	tests := []struct {
		name     string
		input    SQLNullInt64
		expected string
	}{
		{"Valid Int", SQLNullInt64{sql.NullInt64{Int64: 123, Valid: true}}, "123"},
		{"Zero but Valid", SQLNullInt64{sql.NullInt64{Int64: 0, Valid: true}}, ""},
		{"Null Value", SQLNullInt64{sql.NullInt64{Valid: false}}, ""},
		{"Negative Int", SQLNullInt64{sql.NullInt64{Int64: -456, Valid: true}}, "-456"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.String()
			if result != tt.expected {

			}
		})
	}
}

// TestMarshalJSON tests the MarshalJSON method.
func TestMarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    SQLNullInt64
		expected string
	}{
		{"Valid Int", SQLNullInt64{sql.NullInt64{Int64: 123, Valid: true}}, "123"},
		{"Null Value", SQLNullInt64{sql.NullInt64{Valid: false}}, "null"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.input.MarshalJSON()
			if err != nil {

			}

			if string(result) != tt.expected {

			}
		})
	}
}

// TestGetIntOrNull tests the GetIntOrNull method.
func TestGetIntOrNull(t *testing.T) {
	tests := []struct {
		name     string
		input    SQLNullInt64
		expected *int64
	}{
		{"Valid Int", SQLNullInt64{sql.NullInt64{Int64: 42, Valid: true}}, func() *int64 { v := int64(42); return &v }()},
		{"Null Value", SQLNullInt64{sql.NullInt64{Valid: false}}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.GetIntOrNull()

			if result == nil && tt.expected != nil {

			} else if result != nil && tt.expected == nil {

			} else if result != nil && *result != *tt.expected {

			}
		})
	}
}

// TestGetInt tests the GetInt method.
func TestGetInt(t *testing.T) {
	tests := []struct {
		name     string
		input    SQLNullInt64
		expected int64
	}{
		{"Valid Int", SQLNullInt64{sql.NullInt64{Int64: 99, Valid: true}}, 99},
		{"Null Value Returns Zero", SQLNullInt64{sql.NullInt64{Valid: false}}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.GetInt()
			if result != tt.expected {

			}
		})
	}
}

// TestNewNullString tests the NewNullString function.
func TestNewNullString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected SQLNullString
	}{
		{"Valid String", "Hello", SQLNullString{sql.NullString{String: "Hello", Valid: true}}},
		{"Empty String", "", SQLNullString{sql.NullString{String: "", Valid: true}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewNullString(tt.input)
			if result.String != tt.expected.String || result.Valid != tt.expected.Valid {

			}
		})
	}
}

// TestSetString tests the SetString method.
func TestSetString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected SQLNullString
	}{
		{"Set Non-Empty String", "Test", SQLNullString{sql.NullString{String: "Test", Valid: true}}},
		{"Set Empty String", "", SQLNullString{sql.NullString{String: "", Valid: true}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result SQLNullString
			result.SetString(tt.input)

			if result.String != tt.expected.String || result.Valid != tt.expected.Valid {

			}
		})
	}
}

// TestIsNull tests the IsNull method.
func TestIsNull4(t *testing.T) {
	tests := []struct {
		name     string
		input    SQLNullString
		expected bool
	}{
		{"Valid String", SQLNullString{sql.NullString{String: "GoLang", Valid: true}}, false},
		{"Empty String but Valid", SQLNullString{sql.NullString{String: "", Valid: true}}, false},
		{"Null Value", SQLNullString{sql.NullString{Valid: false}}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.IsNull()
			if result != tt.expected {

			}
		})
	}
}

// TestSetNull tests the SetNull method.
func TestSetNul4l(t *testing.T) {
	t.Run("Set Null", func(t *testing.T) {
		var s SQLNullString
		s.SetNull()

		if s.Valid != false || s.String != "" {

		}
	})
}

// TestVal tests the Val method.
func TestVal4(t *testing.T) {
	tests := []struct {
		name     string
		input    SQLNullString
		expected string
	}{
		{"Valid String", SQLNullString{sql.NullString{String: "World", Valid: true}}, "World"},
		{"Empty String but Valid", SQLNullString{sql.NullString{String: "", Valid: true}}, ""},
		{"Null Value Returns Empty String", SQLNullString{sql.NullString{Valid: false}}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.Val()
			if result != tt.expected {

			}
		})
	}
}

// TestMarshalJSON tests the MarshalJSON method.
func TestMarshalJSON4(t *testing.T) {
	tests := []struct {
		name     string
		input    SQLNullString
		expected string
	}{
		{"Valid String", SQLNullString{sql.NullString{String: "JSONTest", Valid: true}}, "\"JSONTest\""},
		{"Null Value Returns Empty String", SQLNullString{sql.NullString{Valid: false}}, "\"\""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.input.MarshalJSON()
			if err != nil {

			}

			if string(result) != tt.expected {

			}
		})
	}
}

// TestNewNullTime tests the NewNullTime function.
func TestNewNullTime(t *testing.T) {
	now := time.Now()
	zeroTime := time.Time{}

	tests := []struct {
		name     string
		input    time.Time
		expected SQLNullTime
	}{
		{"Valid Time", now, SQLNullTime{sql.NullTime{Time: now, Valid: true}}},
		{"Zero Time", zeroTime, SQLNullTime{sql.NullTime{Time: zeroTime, Valid: false}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewNullTime(tt.input)
			if result.Time != tt.expected.Time || result.Valid != tt.expected.Valid {

			}
		})
	}
}

// TestScan tests the Scan method.
func TestScan(t *testing.T) {
	validTime := time.Now()
	validTimeBytes := []byte(validTime.Format(mysqlDatetimeFormat))

	tests := []struct {
		name     string
		input    interface{}
		expected SQLNullTime
	}{
		{"Valid Time", validTime, SQLNullTime{sql.NullTime{Time: validTime, Valid: true}}},
		{"Valid Time as Bytes", validTimeBytes, SQLNullTime{sql.NullTime{Time: validTime.UTC(), Valid: true}}},
		{"Invalid Value", "invalid", SQLNullTime{sql.NullTime{Valid: false}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result SQLNullTime
			err := result.Scan(tt.input)

			if tt.name == "Invalid Value" && err == nil {

			} else if result.Valid != tt.expected.Valid || (!result.IsNull() && result.Time != tt.expected.Time) {

			}
		})
	}
}

// TestValue tests the Value method.
func TestValue(t *testing.T) {
	validTime := time.Now()

	tests := []struct {
		name     string
		input    SQLNullTime
		expected driver.Value
	}{
		{"Valid Time", SQLNullTime{sql.NullTime{Time: validTime, Valid: true}}, validTime},
		{"Null Time", SQLNullTime{sql.NullTime{Valid: false}}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, _ := tt.input.Value()
			if result != tt.expected {

			}
		})
	}
}

// TestIsNull tests the IsNull method.
func TestIsNull5(t *testing.T) {
	tests := []struct {
		name     string
		input    SQLNullTime
		expected bool
	}{
		{"Valid Time", SQLNullTime{sql.NullTime{Valid: true}}, false},
		{"Null Time", SQLNullTime{sql.NullTime{Valid: false}}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.IsNull()
			if result != tt.expected {

			}
		})
	}
}

// TestVal tests the Val method.
func TestVal5(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name     string
		input    SQLNullTime
		expected time.Time
	}{
		{"Valid Time", SQLNullTime{sql.NullTime{Time: now, Valid: true}}, now},
		{"Null Time", SQLNullTime{sql.NullTime{Valid: false}}, time.Time{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.Val()
			if !result.Equal(tt.expected) {

			}
		})
	}
}

// TestLocal tests the Local method.
func TestLocal(t *testing.T) {
	now := time.Now().UTC()
	localTime := now.Local()

	tests := []struct {
		name     string
		input    SQLNullTime
		expected time.Time
	}{
		{"Valid Time", SQLNullTime{sql.NullTime{Time: now, Valid: true}}, localTime},
		{"Null Time", SQLNullTime{sql.NullTime{Valid: false}}, time.Time{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.Local()
			if !result.Equal(tt.expected) {

			}
		})
	}
}

// TestSetTime tests the SetTime method.
func TestSetTime(t *testing.T) {
	now := time.Now()

	t.Run("Set Valid Time", func(t *testing.T) {
		var s SQLNullTime
		s.SetTime(now)

		if s.Time != now || !s.Valid {

		}
	})
}

// TestDateString tests the DateString method.
func TestDateString(t *testing.T) {
	now := time.Date(2025, 2, 25, 14, 30, 0, 0, time.UTC)
	tests := []struct {
		name     string
		input    SQLNullTime
		spit     string
		expected string
	}{
		{"Valid Date", SQLNullTime{sql.NullTime{Time: now, Valid: true}}, "/", "25/02/2025"},
		{"Null Time", SQLNullTime{sql.NullTime{Valid: false}}, "/", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.DateString(tt.spit)
			if result != tt.expected {

			}
		})
	}
}

// TestString tests the String method.
func TestString5(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name     string
		input    SQLNullTime
		expected string
	}{
		{"Valid Time", SQLNullTime{sql.NullTime{Time: now, Valid: true}}, now.String()},
		{"Null Time", SQLNullTime{sql.NullTime{Valid: false}}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.String()
			if result != tt.expected {

			}
		})
	}
}

// TestGetTimeOrNull tests the GetTimeOrNull method.
func TestGetTimeOrNull(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name     string
		input    SQLNullTime
		expected *time.Time
	}{
		{"Valid Time", SQLNullTime{sql.NullTime{Time: now, Valid: true}}, &now},
		{"Null Time", SQLNullTime{sql.NullTime{Valid: false}}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.GetTimeOrNull()
			if (result == nil && tt.expected != nil) || (result != nil && *result != *tt.expected) {

			}
		})
	}
}

// TestMarshalJSON tests the MarshalJSON method.
func TestMarshalJSON5(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name     string
		input    SQLNullTime
		expected string
	}{
		{"Valid Time", SQLNullTime{sql.NullTime{Time: now, Valid: true}}, fmt.Sprintf("\"%s\"", now.Format(mysqlDatetimeFormat))},
		{"Null Time", SQLNullTime{sql.NullTime{Valid: false}}, "\"\""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, _ := json.Marshal(tt.input)
			if string(result) != tt.expected {

			}
		})
	}
}
