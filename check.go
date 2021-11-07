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
		DefaultValueFunc = New()
	})
}

// DefaultValueFunc 提供一个默认
var DefaultValueFunc *ValueFunc

/////////////////////////////////////////////////////////////////////

const (
	// NoCheckTagValue tag 不检测的空标志
	NoCheckTagValue string = "-"

	// CheckTagSplice tag 绑定的检测函数和传入的参数之间的分隔符
	DefaultCheckTagSplice string = ":"
	DefaultCheckTag       string = "check"
)

var (
	// checkTag tag 的标签
	checkTag = DefaultCheckTag

	// checkTagSplice tag 的标分隔符
	checkTagSplice = DefaultCheckTagSplice
)

// SetCheckTag 设置标签的 tag
func SetCheckTag(tag string) {
	if tag == "" {
		return
	}
	checkTag = tag
}

// SetCheckTagSplice 设置标签的 分隔符
func SetCheckTagSplice(splice string) {
	if splice == "" {
		return
	}
	checkTagSplice = splice
}

var (
	ErrNilValueFunc      error = errors.New("check value func is nil")
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

type ValueFunc struct {
	CheckFuncMap
	sync.RWMutex
	checkTag       string
	checkTagSplice string
}

func New() *ValueFunc {
	return &ValueFunc{
		RWMutex:        sync.RWMutex{},
		CheckFuncMap:   make(CheckFuncMap),
		checkTag:       checkTag,
		checkTagSplice: checkTagSplice,
	}
}

func (vf *ValueFunc) SetCheckFuncMap(cfs CheckFuncMap) {
	vf.Lock()
	defer vf.Unlock()

	if cfs == nil {
		return
	}
	vf.CheckFuncMap = cfs
}

func (vf *ValueFunc) GetCheckFuncMap(tag string) CheckFuncMap {
	vf.Lock()
	defer vf.Unlock()

	return vf.CheckFuncMap
}

func (vf *ValueFunc) SetCheckTag(tag string) {
	if tag == "" {
		return
	}
	vf.checkTag = tag
}

// SetCheckTagSplice 设置标签的 分隔符
func (vf *ValueFunc) SetCheckTagSplice(splice string) {
	if splice == "" {
		return
	}
	vf.checkTagSplice = splice
}

// Add 添加一个对应 tag 的检测函数
func (vf *ValueFunc) Add(tag string, f CheckFunc) error {
	if vf == nil {
		return ErrNilValueFunc
	}
	vf.Lock()
	defer vf.Unlock()

	if vf.CheckFuncMap == nil {
		return ErrNilCheckFuncMap
	}

	if tag == "" {
		return ErrNilCheckFuncTag
	}
	if f == nil {
		return ErrNilCheckFunc
	}

	if _, ok := vf.CheckFuncMap[tag]; ok {
		return ErrExistCheckFuncTag
	}
	vf.CheckFuncMap[tag] = f

	return nil
}

/////////////////////////////////////////

func (vf *ValueFunc) AddInt64(tag string, f func(v int64, isNil bool, args string) (msg string, err error)) error {
	return vf.Add(tag, addInt64(f))
}
func (vf *ValueFunc) AddInt32(tag string, f func(v int32, isNil bool, args string) (msg string, err error)) error {
	return vf.Add(tag, addInt32(f))
}
func (vf *ValueFunc) AddInt16(tag string, f func(v int16, isNil bool, args string) (msg string, err error)) error {
	return vf.Add(tag, addInt16(f))
}
func (vf *ValueFunc) AddInt8(tag string, f func(v int8, isNil bool, args string) (msg string, err error)) error {
	return vf.Add(tag, addInt8(f))
}
func (vf *ValueFunc) AddInt(tag string, f func(v int, isNil bool, args string) (msg string, err error)) error {
	return vf.Add(tag, addInt(f))
}
func (vf *ValueFunc) AddString(tag string, f func(v string, isNil bool, args string) (msg string, err error)) error {
	return vf.Add(tag, addString(f))
}
func (vf *ValueFunc) AddBool(tag string, f func(v bool, isNil bool, args string) (msg string, err error)) error {
	return vf.Add(tag, addBool(f))
}

func (vf *ValueFunc) AddUint(tag string, f func(v uint, isNil bool, args string) (msg string, err error)) error {
	return vf.Add(tag, addUint(f))
}
func (vf *ValueFunc) AddUint8(tag string, f func(v uint8, isNil bool, args string) (msg string, err error)) error {
	return vf.Add(tag, addUint8(f))
}
func (vf *ValueFunc) AddUint16(tag string, f func(v uint16, isNil bool, args string) (msg string, err error)) error {
	return vf.Add(tag, addUint16(f))
}
func (vf *ValueFunc) AddUint32(tag string, f func(v uint32, isNil bool, args string) (msg string, err error)) error {
	return vf.Add(tag, addUint32(f))
}
func (vf *ValueFunc) AddUint64(tag string, f func(v uint64, isNil bool, args string) (msg string, err error)) error {
	return vf.Add(tag, addUint64(f))
}
func (vf *ValueFunc) AddFloat32(tag string, f func(v float32, isNil bool, args string) (msg string, err error)) error {
	return vf.Add(tag, addFloat32(f))
}
func (vf *ValueFunc) AddFloat64(tag string, f func(v float64, isNil bool, args string) (msg string, err error)) error {
	return vf.Add(tag, addFloat64(f))
}
func (vf *ValueFunc) AddInterface(tag string, f func(v interface{}, isNil bool, args string) (msg string, err error)) error {
	return vf.Add(tag, addInterface(f))
}
func (vf *ValueFunc) SetValue(tag string, f func(v reflect.Value, canSet bool, args string) (msg string, err error)) error {
	return vf.Add(tag, setValue(f))
}

// Check 通过 cm 检查 in 的信息
func (vf *ValueFunc) Check(in interface{}) (string, error) {
	vf.Lock()
	defer vf.Unlock()
	return check(in, vf)
}

//////////////////////////////////////////////////////////////////////////////////

// Check 通过默认 DefaultValueFunc 检查 in 的信息
func Check(in interface{}) (string, error) {
	return DefaultValueFunc.Check(in)
}

// Add 添加一个对应 tag 的检测函数
func Add(tag string, f CheckFunc) error {
	return DefaultValueFunc.Add(tag, f)
}

func AddInt64(tag string, f func(v int64, isNil bool, args string) (msg string, err error)) error {
	return DefaultValueFunc.AddInt64(tag, f)
}
func AddInt32(tag string, f func(v int32, isNil bool, args string) (msg string, err error)) error {
	return DefaultValueFunc.AddInt32(tag, f)
}
func AddInt16(tag string, f func(v int16, isNil bool, args string) (msg string, err error)) error {
	return DefaultValueFunc.AddInt16(tag, f)
}
func AddInt8(tag string, f func(v int8, isNil bool, args string) (msg string, err error)) error {
	return DefaultValueFunc.AddInt8(tag, f)
}
func AddInt(tag string, f func(v int, isNil bool, args string) (msg string, err error)) error {
	return DefaultValueFunc.AddInt(tag, f)
}
func AddString(tag string, f func(v string, isNil bool, args string) (msg string, err error)) error {
	return DefaultValueFunc.AddString(tag, f)
}
func AddBool(tag string, f func(v bool, isNil bool, args string) (msg string, err error)) error {
	return DefaultValueFunc.AddBool(tag, f)
}

func AddUint(tag string, f func(v uint, isNil bool, args string) (msg string, err error)) error {
	return DefaultValueFunc.AddUint(tag, f)
}
func AddUint8(tag string, f func(v uint8, isNil bool, args string) (msg string, err error)) error {
	return DefaultValueFunc.AddUint8(tag, f)
}
func AddUint16(tag string, f func(v uint16, isNil bool, args string) (msg string, err error)) error {
	return DefaultValueFunc.AddUint16(tag, f)
}
func AddUint32(tag string, f func(v uint32, isNil bool, args string) (msg string, err error)) error {
	return DefaultValueFunc.AddUint32(tag, f)
}
func AddUint64(tag string, f func(v uint64, isNil bool, args string) (msg string, err error)) error {
	return DefaultValueFunc.AddUint64(tag, f)
}
func AddFloat32(tag string, f func(v float32, isNil bool, args string) (msg string, err error)) error {
	return DefaultValueFunc.AddFloat32(tag, f)
}
func AddFloat64(tag string, f func(v float64, isNil bool, args string) (msg string, err error)) error {
	return DefaultValueFunc.AddFloat64(tag, f)
}
func AddInterface(tag string, f func(v interface{}, isNil bool, args string) (msg string, err error)) error {
	return DefaultValueFunc.AddInterface(tag, f)
}
func SetValue(tag string, f func(v reflect.Value, canSet bool, args string) (msg string, err error)) error {
	return DefaultValueFunc.Add(tag, setValue(f))
}

//////////////////////////////////////////////////////////////////////////////////////////

// check 通过 vf.CheckFuncMap 检查 inter 的信息
func check(inter interface{}, vf *ValueFunc) (string, error) {
	if inter == nil {
		return "", ErrNilCheckInterface
	}
	if vf == nil {
		return "", ErrNilCheckFunc
	}

	if vf.checkTag == "" {
		vf.checkTag = checkTag
	}
	if vf.checkTagSplice == "" {
		vf.checkTagSplice = checkTagSplice
	}

	if vf.CheckFuncMap == nil {
		return "", ErrNilCheckFuncMap
	}

	sType := reflect.TypeOf(inter)
	sValue := reflect.ValueOf(inter)

	//fmt.Printf("%d\n", sType.NumField())

	for i := 0; i < sType.NumField(); i++ {
		stField := sType.Field(i)
		stValue := sValue.Field(i)

		checkFuncTag := stField.Tag.Get(vf.checkTag)

		//fmt.Printf("%q -->  %q --> %q\n", stField.Name, stField.Type, checkFuncTag)

		if strings.HasPrefix(checkFuncTag, NoCheckTagValue) {
			continue
		}

		switch stField.Type.Kind() {
		case reflect.Func, reflect.Chan, reflect.UnsafePointer:
			continue
		case reflect.Struct:
			msg, err := check(stValue.Interface(), vf)
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
				msg, err := check(stValue.Elem().Interface(), vf)
				if err != nil {
					return msg, err
				}
			}
		}

		// 正常的处理
		indexI := strings.Index(checkFuncTag, vf.checkTagSplice)
		var args string
		var fName = checkFuncTag
		if indexI > 0 {
			fName = checkFuncTag[:indexI]
			args = checkFuncTag[indexI+1:]
		}
		if f, ok := vf.CheckFuncMap[fName]; ok {
			msg, err := f(stValue, args)
			if err != nil {
				return msg, err
			}
		}

	}

	return "", nil
}
