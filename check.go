// Package valuefunc  检测结构体内部数据是否符合某条件
package valuefunc

import (
	"errors"
	"reflect"
	"strings"
	"sync"
)

var once sync.Once

func init() {
	once.Do(func() {
		DefaultCheckFunc = make(CheckFuncMap)
	})
}

// DefaultCheckFunc 提供一个默认
var DefaultCheckFunc CheckFuncMap

/////////////////////////////////////////////////////////////////////

const (
	// NoCheckTagValue tag 不检测的空标志
	NoCheckTagValue string = "-"

	// CheckTagSplice tag 绑定的检测函数和传入的参数之间的分隔符
	CheckTagSplice  string = ":"
	DefaultCheckTag string = "check"
)

var (
	// checkTag tag 的标签
	checkTag = DefaultCheckTag

	// checkTagSplice tag 的标分隔符
	checkTagSplice = CheckTagSplice
)

var (
	ErrNilCheckFuncMap   error = errors.New("check func map is nil")
	ErrNilCheckFunc      error = errors.New("check func is nil")
	ErrNilCheckFuncTag   error = errors.New("check func tag is empty")
	ErrExistCheckFuncTag error = errors.New("check func tag is already exist")
	ErrNilCheckInterface error = errors.New("check func input interface is nil")
	ErrValueType         error = errors.New("check func check value type wring")
)

//////////////////////////////////////////////////////////////////////////

// CheckFunc 检测 v 信息的方式
type CheckFunc func(v reflect.Value, args string) (msg string, err error)

// CheckFuncMap 对应 tag 绑定的检测函数的集合
type CheckFuncMap map[string]CheckFunc

// Add 添加一个对应 tag 的检测函数
func (cm CheckFuncMap) Add(tag string, f CheckFunc) error {
	if cm == nil {
		return ErrNilCheckFuncMap
	}
	if tag == "" {
		return ErrNilCheckFuncTag
	}
	if f == nil {
		return ErrNilCheckFunc
	}

	if _, ok := cm[tag]; ok {
		return ErrExistCheckFuncTag
	}
	cm[tag] = f

	return nil
}

func (cm CheckFuncMap) AddInt64(tag string, f func(v int64, isNil bool, args string) (msg string, err error)) error {
	return cm.Add(tag, addInt64(f))
}
func (cm CheckFuncMap) AddInt32(tag string, f func(v int32, isNil bool, args string) (msg string, err error)) error {
	return cm.Add(tag, addInt32(f))
}
func (cm CheckFuncMap) AddInt16(tag string, f func(v int16, isNil bool, args string) (msg string, err error)) error {
	return cm.Add(tag, addInt16(f))
}
func (cm CheckFuncMap) AddInt8(tag string, f func(v int8, isNil bool, args string) (msg string, err error)) error {
	return cm.Add(tag, addInt8(f))
}
func (cm CheckFuncMap) AddInt(tag string, f func(v int, isNil bool, args string) (msg string, err error)) error {
	return cm.Add(tag, addInt(f))
}
func (cm CheckFuncMap) AddString(tag string, f func(v string, isNil bool, args string) (msg string, err error)) error {
	return cm.Add(tag, addString(f))
}
func (cm CheckFuncMap) AddBool(tag string, f func(v bool, isNil bool, args string) (msg string, err error)) error {
	return cm.Add(tag, addBool(f))
}

func (cm CheckFuncMap) AddUint(tag string, f func(v uint, isNil bool, args string) (msg string, err error)) error {
	return cm.Add(tag, addUint(f))
}
func (cm CheckFuncMap) AddUint8(tag string, f func(v uint8, isNil bool, args string) (msg string, err error)) error {
	return cm.Add(tag, addUint8(f))
}
func (cm CheckFuncMap) AddUint16(tag string, f func(v uint16, isNil bool, args string) (msg string, err error)) error {
	return cm.Add(tag, addUint16(f))
}
func (cm CheckFuncMap) AddUint32(tag string, f func(v uint32, isNil bool, args string) (msg string, err error)) error {
	return cm.Add(tag, addUint32(f))
}
func (cm CheckFuncMap) AddUint64(tag string, f func(v uint64, isNil bool, args string) (msg string, err error)) error {
	return cm.Add(tag, addUint64(f))
}
func (cm CheckFuncMap) AddFloat32(tag string, f func(v float32, isNil bool, args string) (msg string, err error)) error {
	return cm.Add(tag, addFloat32(f))
}
func (cm CheckFuncMap) AddFloat64(tag string, f func(v float64, isNil bool, args string) (msg string, err error)) error {
	return cm.Add(tag, addFloat64(f))
}
func (cm CheckFuncMap) AddInterface(tag string, f func(v interface{}, isNil bool, args string) (msg string, err error)) error {
	return cm.Add(tag, addInterface(f))
}
func (cm CheckFuncMap) SetValue(tag string, f func(v reflect.Value, canSet bool, args string) (msg string, err error)) error {
	return cm.Add(tag, setValue(f))
}

// Check 通过 cm 检查 in 的信息
func (cm CheckFuncMap) Check(in interface{}) (string, error) {
	return check(in, cm)
}

//////////////////////////////////////////////////////////////////////////////////

// SetTag 设置标签的 tag
func SetTag(tag string) {
	if tag == "" {
		return
	}
	checkTag = tag
}

// SetTagSplice 设置标签的 分隔符
func SetTagSplice(splice string) {
	if splice == "" {
		return
	}
	checkTagSplice = splice
}

// Check 通过默认 DefaultCheckFunc 检查 in 的信息
func Check(in interface{}) (string, error) {
	return DefaultCheckFunc.Check(in)
}

// Add 添加一个对应 tag 的检测函数
func Add(tag string, f CheckFunc) error {
	return DefaultCheckFunc.Add(tag, f)
}

func AddInt64(tag string, f func(v int64, isNil bool, args string) (msg string, err error)) error {
	return DefaultCheckFunc.AddInt64(tag, f)
}
func AddInt32(tag string, f func(v int32, isNil bool, args string) (msg string, err error)) error {
	return DefaultCheckFunc.AddInt32(tag, f)
}
func AddInt16(tag string, f func(v int16, isNil bool, args string) (msg string, err error)) error {
	return DefaultCheckFunc.AddInt16(tag, f)
}
func AddInt8(tag string, f func(v int8, isNil bool, args string) (msg string, err error)) error {
	return DefaultCheckFunc.AddInt8(tag, f)
}
func AddInt(tag string, f func(v int, isNil bool, args string) (msg string, err error)) error {
	return DefaultCheckFunc.AddInt(tag, f)
}
func AddString(tag string, f func(v string, isNil bool, args string) (msg string, err error)) error {
	return DefaultCheckFunc.AddString(tag, f)
}
func AddBool(tag string, f func(v bool, isNil bool, args string) (msg string, err error)) error {
	return DefaultCheckFunc.AddBool(tag, f)
}

func AddUint(tag string, f func(v uint, isNil bool, args string) (msg string, err error)) error {
	return DefaultCheckFunc.AddUint(tag, f)
}
func AddUint8(tag string, f func(v uint8, isNil bool, args string) (msg string, err error)) error {
	return DefaultCheckFunc.AddUint8(tag, f)
}
func AddUint16(tag string, f func(v uint16, isNil bool, args string) (msg string, err error)) error {
	return DefaultCheckFunc.AddUint16(tag, f)
}
func AddUint32(tag string, f func(v uint32, isNil bool, args string) (msg string, err error)) error {
	return DefaultCheckFunc.AddUint32(tag, f)
}
func AddUint64(tag string, f func(v uint64, isNil bool, args string) (msg string, err error)) error {
	return DefaultCheckFunc.AddUint64(tag, f)
}
func AddFloat32(tag string, f func(v float32, isNil bool, args string) (msg string, err error)) error {
	return DefaultCheckFunc.AddFloat32(tag, f)
}
func AddFloat64(tag string, f func(v float64, isNil bool, args string) (msg string, err error)) error {
	return DefaultCheckFunc.AddFloat64(tag, f)
}
func AddInterface(tag string, f func(v interface{}, isNil bool, args string) (msg string, err error)) error {
	return DefaultCheckFunc.AddInterface(tag, f)
}
func SetValue(tag string, f func(v reflect.Value, canSet bool, args string) (msg string, err error)) error {
	return DefaultCheckFunc.Add(tag, setValue(f))
}

//////////////////////////////////////////////////////////////////////////////////////////

// check 通过 checkFuncMap 检查 inter 的信息
func check(inter interface{}, checkFuncMap CheckFuncMap) (string, error) {
	if inter == nil {
		return "", ErrNilCheckInterface
	}
	if checkFuncMap == nil {
		return "", ErrNilCheckFuncMap
	}

	sType := reflect.TypeOf(inter)
	sValue := reflect.ValueOf(inter)

	//fmt.Printf("%d\n", sType.NumField())

	for i := 0; i < sType.NumField(); i++ {
		stField := sType.Field(i)
		stValue := sValue.Field(i)

		checkFuncTag := stField.Tag.Get(checkTag)

		//fmt.Printf("%q -->  %q --> %q\n", stField.Name, stField.Type, checkFuncTag)

		if strings.HasPrefix(checkFuncTag, NoCheckTagValue) {
			continue
		}

		switch stField.Type.Kind() {
		case reflect.Func, reflect.Chan, reflect.UnsafePointer:
			continue
		case reflect.Struct:
			msg, err := check(stValue.Interface(), checkFuncMap)
			if err != nil {
				return msg, err
			}
			continue
		case reflect.Ptr:
			if stValue.IsNil() {
				continue
			}

			switch stValue.Elem().Type().Kind() {
			case reflect.Struct:
				msg, err := check(stValue.Elem().Interface(), checkFuncMap)
				if err != nil {
					return msg, err
				}
			}
		}

		// 正常的处理
		indexI := strings.Index(checkFuncTag, checkTagSplice)
		var args string
		var fName = checkFuncTag
		if indexI > 0 {
			fName = checkFuncTag[:indexI]
			args = checkFuncTag[indexI+1:]
		}
		if f, ok := checkFuncMap[fName]; ok {
			msg, err := f(stValue, args)
			if err != nil {
				return msg, err
			}
		}

	}

	return "", nil
}
