package env

import (
	"errors"
	"fmt"
	"reflect"
)

// Decode decodes environment variables into the given value.
func Decode(v interface{}) error {
	value := reflect.ValueOf(v)
	if value.Kind() != reflect.Ptr {
		return errors.New("env: Decode expects a pointer")
	}
	value = value.Elem()
	if value.Kind() != reflect.Struct {
		return errors.New("env: Decode expects a pointer to a struct")
	}
	return decode(value)
}

// decode decodes environment variables into the given value.
func decode(value reflect.Value) error {
	tp := value.Type()
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		if field.Kind() == reflect.Struct {
			if err := decode(field); err != nil {
				return err
			}
			continue
		}
		tf := tp.Field(i)
		if !tf.IsExported() {
			continue
		}
		tag := tf.Tag.Get("env")
		if tag == "-" {
			continue
		} else if tag == "" {
			tag = tp.Field(i).Name
		}
		def := Env(tf.Tag.Get("default"))
		if err := set(field, tag, def); err != nil {
			return err
		}
	}
	return nil
}

// set sets the given value from the environment variable.
func set(value reflect.Value, name string, def Env) error {
	env := Name(name)
	switch value.Kind() {
	case reflect.String:
		value.SetString(env.StringOrElse(def.String()))
	case reflect.Bool:
		value.SetBool(env.BoolOrElse(def.Bool()))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value.SetInt(env.Int64OrElse(def.Int64()))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		value.SetUint(env.Uint64OrElse(def.Uint64()))
	case reflect.Float32, reflect.Float64:
		value.SetFloat(env.Float64OrElse(def.Float64()))
	default:
		return fmt.Errorf("env: unsupported type %s", value.Kind())
	}
	return nil
}
