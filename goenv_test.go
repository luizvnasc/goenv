package goenv_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/luizvnasc/goenv"
)

type TestCase struct {
	title     string
	variables map[string]string
	input     Test
	want      Test
}

type Test struct {
	StrField     string  `env:"STRFIELD"`
	IntField     int     `env:"INTFIELD"`
	Int16Field   int16   `env:"INT16FIELD"`
	Int32Field   int32   `env:"INT32FIELD"`
	Int64Field   int64   `env:"INT64FIELD"`
	Float32Field float32 `env:"FLOAT32FIELD"`
	Float64Field float32 `env:"FLOAT64FIELD"`
	BoolField    bool    `env:"BOOLFIELD"`
	NestedField  struct {
		StrField     string  `env:"NSTRFIELD"`
		Int64Field   int64   `env:"NINT64FIELD"`
		Float32Field float32 `env:"NFLOAT32FIELD"`
	}
}

var testCases = []TestCase{
	TestCase{
		title:     "Unmarshal string field",
		input:     Test{},
		variables: map[string]string{"STRFIELD": "test"},
		want:      Test{StrField: "test"},
	},
	TestCase{
		title:     "Unmarshal int field",
		input:     Test{},
		variables: map[string]string{"INTFIELD": "1"},
		want:      Test{IntField: 1},
	},
	TestCase{
		title:     "Unmarshal int16 field",
		input:     Test{},
		variables: map[string]string{"INT16FIELD": "1"},
		want:      Test{Int16Field: 1},
	},
	TestCase{
		title:     "Unmarshal int32 field",
		input:     Test{},
		variables: map[string]string{"INT32FIELD": "1"},
		want:      Test{Int32Field: 1},
	},
	TestCase{
		title:     "Unmarshal int64 field",
		input:     Test{},
		variables: map[string]string{"INT64FIELD": "1"},
		want:      Test{Int64Field: 1},
	},
	TestCase{
		title:     "Unmarshal float32 field",
		input:     Test{},
		variables: map[string]string{"FLOAT32FIELD": "3.1415"},
		want:      Test{Float32Field: 3.1415},
	},
	TestCase{
		title:     "Unmarshal float64 field",
		input:     Test{},
		variables: map[string]string{"FLOAT64FIELD": "3.14159265358979323846"},
		want:      Test{Float64Field: 3.14159265358979323846},
	},
	TestCase{
		title:     "Unmarshal bool field",
		input:     Test{},
		variables: map[string]string{"BOOLFIELD": "true"},
		want:      Test{BoolField: true},
	},
	TestCase{
		title: "Unmarshal nested struct",
		input: Test{},
		variables: map[string]string{"BOOLFIELD": "true",
			"STRFIELD":      "test",
			"INTFIELD":      "1",
			"INT16FIELD":    "1",
			"INT32FIELD":    "1",
			"INT64FIELD":    "1",
			"FLOAT32FIELD":  "3.1415",
			"FLOAT64FIELD":  "3.14159265358979323846",
			"NFLOAT32FIELD": "3.1415",
			"NINT64FIELD":   "1",
			"NSTRFIELD":     "test",
		},
		want: Test{StrField: "test",
			IntField:     1,
			Int16Field:   1,
			Int32Field:   1,
			Int64Field:   1,
			BoolField:    true,
			Float32Field: 3.1415,
			Float64Field: 3.14159265358979323846,
			NestedField: struct {
				StrField     string  `env:"NSTRFIELD"`
				Int64Field   int64   `env:"NINT64FIELD"`
				Float32Field float32 `env:"NFLOAT32FIELD"`
			}{
				StrField:     "test",
				Int64Field:   1,
				Float32Field: 3.1415,
			},
		},
	},
}

func TestUnmarshal(t *testing.T) {

	for _, test := range testCases {
		t.Run(test.title, func(t *testing.T) {
			// Set the environment variables
			for k, v := range test.variables {
				os.Setenv(k, v)
			}
			if err := goenv.Unmarshal(&test.input); err != nil {
				t.Errorf("Error unmarshalling: %v", err)
			}
			if !reflect.DeepEqual(test.want, test.input) {
				t.Errorf("Error unmarshalling, got %v but want %v", test.input, test.want)
			}
			for k := range test.variables {
				os.Setenv(k, "")
			}
		})
	}
}
