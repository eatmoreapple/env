package env

import (
	"errors"
	"os"
	"reflect"
)

// Encode encodes the struct v into the environment
func Encode(v interface{}) error {
	value := reflect.ValueOf(v)
	for value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	if value.Kind() != reflect.Struct {
		return errors.New("env: Encode expects a struct")
	}
	return encode(value)
}

func encode(value reflect.Value) error {
	tp := value.Type()
	for i := 0; i < value.NumField(); i++ {
		field := tp.Field(i)
		if field.Type.Kind() == reflect.Struct {
			if err := encode(value.Field(i)); err != nil {
				return err
			}
			continue
		}
		tf := tp.Field(i)
		env := tf.Tag.Get("env")
		if env == "" || env == "-" {
			continue
		}
		if err := os.Setenv(env, value.Field(i).String()); err != nil {
			return err
		}
	}
	return nil
}
