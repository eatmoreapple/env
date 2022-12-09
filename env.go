package env

import (
	"os"
	"strconv"
)

// Env is a wrapper for os.Getenv
type Env string

// String returns the value of the environment variable
func (e Env) String() string {
	return string(e)
}

// Int returns the value of the environment variable as an int
func (e Env) Int() int {
	if v, err := strconv.Atoi(e.String()); err == nil {
		return v
	}
	return 0
}

// MustInt returns the value of the environment variable as an int or panics
func (e Env) MustInt() int {
	if v, err := strconv.Atoi(e.String()); err == nil {
		return v
	}
	panic("env: invalid int")
}

// Int8 returns the value of the environment variable as an int8
func (e Env) Int8() int8 {
	if v, err := strconv.ParseInt(e.String(), 10, 8); err == nil {
		return int8(v)
	}
	return 0
}

// MustInt8 returns the value of the environment variable as an int8 or panics
func (e Env) MustInt8() int8 {
	if v, err := strconv.ParseInt(e.String(), 10, 8); err == nil {
		return int8(v)
	}
	panic("env: invalid int8")
}

// Int16 returns the value of the environment variable as an int16
func (e Env) Int16() int16 {
	if v, err := strconv.ParseInt(e.String(), 10, 16); err == nil {
		return int16(v)
	}
	return 0
}

// MustInt16 returns the value of the environment variable as an int16 or panics
func (e Env) MustInt16() int16 {
	if v, err := strconv.ParseInt(e.String(), 10, 16); err == nil {
		return int16(v)
	}
	panic("env: invalid int16")
}

// Int32 returns the value of the environment variable as an int32
func (e Env) Int32() int32 {
	if v, err := strconv.ParseInt(e.String(), 10, 32); err == nil {
		return int32(v)
	}
	return 0
}

// MustInt32 returns the value of the environment variable as an int32 or panics
func (e Env) MustInt32() int32 {
	if v, err := strconv.ParseInt(e.String(), 10, 32); err == nil {
		return int32(v)
	}
	panic("env: invalid int32")
}

// Int64 returns the value of the environment variable as an int64
func (e Env) Int64() int64 {
	if v, err := strconv.ParseInt(e.String(), 10, 64); err == nil {
		return v
	}
	return 0
}

// MustInt64 returns the value of the environment variable as an int64 or panics
func (e Env) MustInt64() int64 {
	if v, err := strconv.ParseInt(e.String(), 10, 64); err == nil {
		return v
	}
	panic("env: invalid int64")
}

// Uint returns the value of the environment variable as an uint
func (e Env) Uint() uint {
	if v, err := strconv.ParseUint(e.String(), 10, 64); err == nil {
		return uint(v)
	}
	return 0
}

// MustUint returns the value of the environment variable as an uint or panics
func (e Env) MustUint() uint {
	if v, err := strconv.ParseUint(e.String(), 10, 64); err == nil {
		return uint(v)
	}
	panic("env: invalid uint")
}

// Uint8 returns the value of the environment variable as an uint8
func (e Env) Uint8() uint8 {
	if v, err := strconv.ParseUint(e.String(), 10, 8); err == nil {
		return uint8(v)
	}
	return 0
}

// MustUint8 returns the value of the environment variable as an uint8 or panics
func (e Env) MustUint8() uint8 {
	if v, err := strconv.ParseUint(e.String(), 10, 8); err == nil {
		return uint8(v)
	}
	panic("env: invalid uint8")
}

// Uint16 returns the value of the environment variable as an uint16
func (e Env) Uint16() uint16 {
	if v, err := strconv.ParseUint(e.String(), 10, 16); err == nil {
		return uint16(v)
	}
	return 0
}

// MustUint16 returns the value of the environment variable as an uint16 or panics
func (e Env) MustUint16() uint16 {
	if v, err := strconv.ParseUint(e.String(), 10, 16); err == nil {
		return uint16(v)
	}
	panic("env: invalid uint16")
}

// Uint32 returns the value of the environment variable as an uint32
func (e Env) Uint32() uint32 {
	if v, err := strconv.ParseUint(e.String(), 10, 32); err == nil {
		return uint32(v)
	}
	return 0
}

// MustUint32 returns the value of the environment variable as an uint32 or panics
func (e Env) MustUint32() uint32 {
	if v, err := strconv.ParseUint(e.String(), 10, 32); err == nil {
		return uint32(v)
	}
	panic("env: invalid uint32")
}

// Uint64 returns the value of the environment variable as an uint64
func (e Env) Uint64() uint64 {
	if v, err := strconv.ParseUint(e.String(), 10, 64); err == nil {
		return v
	}
	return 0
}

// MustUint64 returns the value of the environment variable as an uint64 or panics
func (e Env) MustUint64() uint64 {
	if v, err := strconv.ParseUint(e.String(), 10, 64); err == nil {
		return v
	}
	panic("env: invalid uint64")
}

// Float64 returns the value of the environment variable as a float64
func (e Env) Float64() float64 {
	if v, err := strconv.ParseFloat(e.String(), 64); err == nil {
		return v
	}
	return 0
}

// MustFloat64 returns the value of the environment variable as a float64 or panics
func (e Env) MustFloat64() float64 {
	if v, err := strconv.ParseFloat(e.String(), 64); err == nil {
		return v
	}
	panic("env: invalid float64")
}

// Float32 returns the value of the environment variable as a float32
func (e Env) Float32() float32 {
	if v, err := strconv.ParseFloat(e.String(), 32); err == nil {
		return float32(v)
	}
	return 0
}

// MustFloat32 returns the value of the environment variable as a float32 or panics
func (e Env) MustFloat32() float32 {
	if v, err := strconv.ParseFloat(e.String(), 32); err == nil {
		return float32(v)
	}
	panic("env: invalid float32")
}

// Bool returns the value of the environment variable as a bool
func (e Env) Bool() bool {
	if v, err := strconv.ParseBool(e.String()); err == nil {
		return v
	}
	return false
}

// MustBool returns the value of the environment variable as a bool or panics
func (e Env) MustBool() bool {
	if v, err := strconv.ParseBool(e.String()); err == nil {
		return v
	}
	panic("env: invalid bool")
}

// IsEmpty returns true if the environment variable is empty
func (e Env) IsEmpty() bool {
	return e.String() == ""
}

// StringOrElse returns the value of the environment variable or the default value
func (e Env) StringOrElse(v string) string {
	if !e.IsEmpty() {
		return e.String()
	}
	return v
}

// IntOrElse returns the value of the environment variable as an int or the default value
func (e Env) IntOrElse(v int) int {
	if !e.IsEmpty() {
		return e.Int()
	}
	return v
}

// Int8OrElse returns the value of the environment variable as an int8 or the default value
func (e Env) Int8OrElse(v int8) int8 {
	if !e.IsEmpty() {
		return e.Int8()
	}
	return v
}

// Int16OrElse returns the value of the environment variable as an int16 or the default value
func (e Env) Int16OrElse(v int16) int16 {
	if !e.IsEmpty() {
		return e.Int16()
	}
	return v
}

// Int32OrElse returns the value of the environment variable as an int32 or the default value
func (e Env) Int32OrElse(v int32) int32 {
	if !e.IsEmpty() {
		return e.Int32()
	}
	return v
}

// Int64OrElse returns the value of the environment variable as an int64 or the default value
func (e Env) Int64OrElse(v int64) int64 {
	if !e.IsEmpty() {
		return e.Int64()
	}
	return v
}

// UintOrElse returns the value of the environment variable as an uint or the default value
func (e Env) UintOrElse(v uint) uint {
	if !e.IsEmpty() {
		return e.Uint()
	}
	return v
}

// Uint8OrElse returns the value of the environment variable as an uint8 or the default value
func (e Env) Uint8OrElse(v uint8) uint8 {
	if !e.IsEmpty() {
		return e.Uint8()
	}
	return v
}

// Uint16OrElse returns the value of the environment variable as an uint16 or the default value
func (e Env) Uint16OrElse(v uint16) uint16 {
	if !e.IsEmpty() {
		return e.Uint16()
	}
	return v
}

// Uint32OrElse returns the value of the environment variable as an uint32 or the default value
func (e Env) Uint32OrElse(v uint32) uint32 {
	if !e.IsEmpty() {
		return e.Uint32()
	}
	return v
}

// Uint64OrElse returns the value of the environment variable as an uint64 or the default value
func (e Env) Uint64OrElse(v uint64) uint64 {
	if !e.IsEmpty() {
		return e.Uint64()
	}
	return v
}

// Float64OrElse returns the value of the environment variable as a float64 or the default value
func (e Env) Float64OrElse(v float64) float64 {
	if !e.IsEmpty() {
		return e.Float64()
	}
	return v
}

// Float32OrElse returns the value of the environment variable as a float32 or the default value
func (e Env) Float32OrElse(v float32) float32 {
	if !e.IsEmpty() {
		return e.Float32()
	}
	return v
}

// BoolOrElse returns the value of the environment variable as a bool or the default value
func (e Env) BoolOrElse(v bool) bool {
	if !e.IsEmpty() {
		return e.Bool()
	}
	return v
}

// Name returns the name of the environment variable
func Name(key string) Env {
	return Env(os.Getenv(key))
}
