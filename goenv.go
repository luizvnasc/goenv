package goenv

import (
	"os"
	"reflect"
	"strconv"
)

const tagName = "env"

// Unmarshal the environment variables to a struct
func Unmarshal(i interface{}) error {
	t := reflect.TypeOf(i).Elem()
	v := reflect.ValueOf(i).Elem()
	return unmarshal(t, v)
}

func unmarshal(t reflect.Type, v reflect.Value) (err error) {
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)
		switch field.Type().Kind() {
		case reflect.Struct:
			unmarshal(fieldType.Type, field)
		case reflect.String:
			setString(field, fieldType)
		case reflect.Int:
			setInt(field, fieldType)
		case reflect.Int16:
			setInt(field, fieldType)
		case reflect.Int32:
			setInt(field, fieldType)
		case reflect.Int64:
			setInt(field, fieldType)
		case reflect.Float32:
			setFloat(field, fieldType)
		case reflect.Float64:
			setFloat(field, fieldType)
		case reflect.Bool:
			setBool(field, fieldType)
		}
	}
	return
}

func setString(v reflect.Value, t reflect.StructField) {
	value := getEnvVariable(t)
	if v.CanSet() {
		v.SetString(value)
	}
}

func setInt(field reflect.Value, t reflect.StructField) {
	value := getEnvVariable(t)
	if v, err := strconv.Atoi(value); field.CanSet() && err == nil {
		field.SetInt(int64(v))
	}
}

func setFloat(field reflect.Value, t reflect.StructField) {
	value := getEnvVariable(t)
	if v, err := strconv.ParseFloat(value, 64); field.CanSet() && err == nil {
		field.SetFloat(float64(v))
	}
}

func setBool(field reflect.Value, t reflect.StructField) {
	value := getEnvVariable(t)
	if v, err := strconv.ParseBool(value); field.CanSet() && err == nil {
		field.SetBool(v)
	}
}

func getEnvVariable(s reflect.StructField) string {
	envVar := s.Tag.Get(tagName)
	return os.Getenv(envVar)
}
