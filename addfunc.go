package valuefunc

import (
	"reflect"
)

// addString  CheckFunc 转 string 函数
func addString(f func(v string, isNil bool, args string) (msg string, err error)) CheckFunc {
	return func(rv reflect.Value, args_ string) (string, error) {
		if f == nil {
			return "", ErrNilCheckFunc
		}

		switch rv.Type().Kind() {
		case reflect.String:
			return f(rv.String(), false, args_)
		case reflect.Ptr:
			if rv.IsNil() {
				return f("", true, args_)
			}
			if rv.Elem().Kind() == reflect.String {
				return f(rv.Elem().String(), false, args_)
			}
		}
		return "", ErrValueType
	}
}

// addBool  CheckFunc 转 bool 函数
func addBool(f func(v bool, isNil bool, args string) (msg string, err error)) CheckFunc {
	return func(rv reflect.Value, args_ string) (string, error) {
		if f == nil {
			return "", ErrNilCheckFunc
		}

		switch rv.Type().Kind() {
		case reflect.Bool:
			return f(rv.Bool(), false, args_)
		case reflect.Ptr:
			if rv.IsNil() {
				return f(false, true, args_)
			}
			if rv.Elem().Kind() == reflect.Bool {
				return f(rv.Elem().Bool(), false, args_)
			}
		}
		return "", ErrValueType
	}
}

// addInt  CheckFunc 转 int 函数
func addInt(f func(v int, isNil bool, args string) (msg string, err error)) CheckFunc {
	return func(rv reflect.Value, args_ string) (string, error) {
		if f == nil {
			return "", ErrNilCheckFunc
		}

		switch rv.Type().Kind() {
		case reflect.Int:
			return f(int(rv.Int()), false, args_)
		case reflect.Ptr:
			if rv.IsNil() {
				return f(0, true, args_)
			}
			if rv.Elem().Kind() == reflect.Int {
				return f(int(rv.Elem().Int()), false, args_)
			}
		}
		return "", ErrValueType
	}
}

// addInt8  CheckFunc 转 int8 函数
func addInt8(f func(v int8, isNil bool, args string) (msg string, err error)) CheckFunc {
	return func(rv reflect.Value, args_ string) (string, error) {
		if f == nil {
			return "", ErrNilCheckFunc
		}

		switch rv.Type().Kind() {
		case reflect.Int8:
			return f(int8(rv.Int()), false, args_)
		case reflect.Ptr:
			if rv.IsNil() {
				return f(0, true, args_)
			}
			if rv.Elem().Kind() == reflect.Int8 {
				return f(int8(rv.Elem().Int()), false, args_)
			}
		}
		return "", ErrValueType
	}
}

// addInt16  CheckFunc 转 int16 函数
func addInt16(f func(v int16, isNil bool, args string) (msg string, err error)) CheckFunc {
	return func(rv reflect.Value, args_ string) (string, error) {
		if f == nil {
			return "", ErrNilCheckFunc
		}

		switch rv.Type().Kind() {
		case reflect.Int16:
			return f(int16(rv.Int()), false, args_)
		case reflect.Ptr:
			if rv.IsNil() {
				return f(0, true, args_)
			}
			if rv.Elem().Kind() == reflect.Int16 {
				return f(int16(rv.Elem().Int()), false, args_)
			}
		}
		return "", ErrValueType
	}
}

// addBool  CheckFunc 转 bool 函数
func addInt32(f func(v int32, isNil bool, args string) (msg string, err error)) CheckFunc {
	return func(rv reflect.Value, args_ string) (string, error) {
		if f == nil {
			return "", ErrNilCheckFunc
		}

		switch rv.Type().Kind() {
		case reflect.Int32:
			return f(int32(rv.Int()), false, args_)
		case reflect.Ptr:
			if rv.IsNil() {
				return f(0, true, args_)
			}
			if rv.Elem().Kind() == reflect.Int32 {
				return f(int32(rv.Elem().Int()), false, args_)
			}
		}
		return "", ErrValueType
	}
}

// addInt64 函数类型转换 reflect.Value 转为 int64
func addInt64(f func(v int64, isNil bool, args string) (msg string, err error)) CheckFunc {
	return func(rv reflect.Value, args_ string) (string, error) {
		if f == nil {
			return "", ErrNilCheckFunc
		}

		switch rv.Type().Kind() {
		case reflect.Int:
			return f(int64(rv.Int()), false, args_)
		case reflect.Ptr:
			if rv.IsNil() {
				return f(0, true, args_)
			}
			if rv.Elem().Kind() == reflect.Int64 {
				return f(int64(rv.Elem().Int()), false, args_)
			}
		}
		return "", ErrValueType
	}
}

// addUint  CheckFunc 转 uint 函数
func addUint(f func(v uint, isNil bool, args string) (msg string, err error)) CheckFunc {
	return func(rv reflect.Value, args_ string) (string, error) {
		if f == nil {
			return "", ErrNilCheckFunc
		}

		switch rv.Type().Kind() {
		case reflect.Uint:
			return f(uint(rv.Uint()), false, args_)
		case reflect.Ptr:
			if rv.IsNil() {
				return f(0, true, args_)
			}
			if rv.Elem().Kind() == reflect.Uint {
				return f(uint(rv.Elem().Uint()), false, args_)
			}
		}
		return "", ErrValueType
	}
}

// addUint8  CheckFunc 转 uint8 函数
func addUint8(f func(v uint8, isNil bool, args string) (msg string, err error)) CheckFunc {
	return func(rv reflect.Value, args_ string) (string, error) {
		if f == nil {
			return "", ErrNilCheckFunc
		}

		switch rv.Type().Kind() {
		case reflect.Uint8:
			return f(uint8(rv.Uint()), false, args_)
		case reflect.Ptr:
			if rv.IsNil() {
				return f(0, true, args_)
			}
			if rv.Elem().Kind() == reflect.Uint8 {
				return f(uint8(rv.Elem().Uint()), false, args_)
			}
		}
		return "", ErrValueType
	}
}

// addUint16  CheckFunc 转 uint16 函数
func addUint16(f func(v uint16, isNil bool, args string) (msg string, err error)) CheckFunc {
	return func(rv reflect.Value, args_ string) (string, error) {
		if f == nil {
			return "", ErrNilCheckFunc
		}

		switch rv.Type().Kind() {
		case reflect.Uint16:
			return f(uint16(rv.Int()), false, args_)
		case reflect.Ptr:
			if rv.IsNil() {
				return f(0, true, args_)
			}
			if rv.Elem().Kind() == reflect.Uint16 {
				return f(uint16(rv.Elem().Uint()), false, args_)
			}
		}
		return "", ErrValueType
	}
}

// addUint32  CheckFunc 转 uint32 函数
func addUint32(f func(v uint32, isNil bool, args string) (msg string, err error)) CheckFunc {
	return func(rv reflect.Value, args_ string) (string, error) {
		if f == nil {
			return "", ErrNilCheckFunc
		}

		switch rv.Type().Kind() {
		case reflect.Uint32:
			return f(uint32(rv.Uint()), false, args_)
		case reflect.Ptr:
			if rv.IsNil() {
				return f(0, true, args_)
			}
			if rv.Elem().Kind() == reflect.Uint32 {
				return f(uint32(rv.Elem().Uint()), false, args_)
			}
		}
		return "", ErrValueType
	}
}

// addUint64  CheckFunc 转 bool 函数
func addUint64(f func(v uint64, isNil bool, args string) (msg string, err error)) CheckFunc {
	return func(rv reflect.Value, args_ string) (string, error) {
		if f == nil {
			return "", ErrNilCheckFunc
		}

		switch rv.Type().Kind() {
		case reflect.Uint64:
			return f(uint64(rv.Uint()), false, args_)
		case reflect.Ptr:
			if rv.IsNil() {
				return f(0, true, args_)
			}
			if rv.Elem().Kind() == reflect.Uint64 {
				return f(uint64(rv.Elem().Int()), false, args_)
			}
		}
		return "", ErrValueType
	}
}

// addFloat32  CheckFunc 转 float32 函数
func addFloat32(f func(v float32, isNil bool, args string) (msg string, err error)) CheckFunc {
	return func(rv reflect.Value, args_ string) (string, error) {
		if f == nil {
			return "", ErrNilCheckFunc
		}

		switch rv.Type().Kind() {
		case reflect.Float32:
			return f(float32(rv.Float()), false, args_)
		case reflect.Ptr:
			if rv.IsNil() {
				return f(0, true, args_)
			}
			if rv.Elem().Kind() == reflect.Float32 {
				return f(float32(rv.Elem().Int()), false, args_)
			}
		}
		return "", ErrValueType
	}
}

// addFloat64  CheckFunc 转 float64 函数
func addFloat64(f func(v float64, isNil bool, args string) (msg string, err error)) CheckFunc {
	return func(rv reflect.Value, args_ string) (string, error) {
		if f == nil {
			return "", ErrNilCheckFunc
		}

		switch rv.Type().Kind() {
		case reflect.Float64:
			return f(float64(rv.Float()), false, args_)
		case reflect.Ptr:
			if rv.IsNil() {
				return f(0, true, args_)
			}
			if rv.Elem().Kind() == reflect.Float64 {
				return f(float64(rv.Elem().Float()), false, args_)
			}
		}
		return "", ErrValueType
	}
}

// addInterface  CheckFunc 转 interface 函数
func addInterface(f func(v interface{}, isNil bool, args string) (msg string, err error)) CheckFunc {
	return func(rv reflect.Value, args_ string) (string, error) {
		if f == nil {
			return "", ErrNilCheckFunc
		}
		switch rv.Type().Kind() {
		case reflect.Ptr:
			if rv.IsNil() {
				return f(0, true, args_)
			}
			return f(rv.Elem().Interface(), false, args_)
		}
		return f(rv.Interface(), false, args_)
	}
}

// setValue
func setValue(f func(v reflect.Value, canSet bool, args string) (msg string, err error)) CheckFunc {
	return func(rv reflect.Value, args_ string) (string, error) {
		if f == nil {
			return "", ErrNilCheckFunc
		}
		if rv.Kind() == reflect.Ptr {
			if rv.IsNil() {
				return f(rv.Elem(), rv.Elem().CanSet(), args_)
			}
			return f(rv.Elem(), rv.Elem().CanSet(), args_)
		}
		return f(rv, rv.CanSet(), args_)
	}
}
