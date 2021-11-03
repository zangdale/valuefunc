package valuefunc

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"testing"
)

func TestTTTag(t *testing.T) {

	AddString("name", func(s string, isNil bool, args string) (msg string, err error) {
		if isNil {
			return args, fmt.Errorf("not nil")
		}

		fmt.Println("name", s, args)

		if s == "Hello" {
			return args, fmt.Errorf("not same %s", s)
		}
		return args, nil
	})

	AddInt("age", func(v int, isNil bool, args string) (msg string, err error) {
		if isNil {
			return args, fmt.Errorf("not nil")
		}
		fmt.Println("age", v, args)
		if v < 18 {
			return args, fmt.Errorf("not big %d", v)
		}
		return args, nil
	})

	// Add("value", func(v reflect.Value, params string) (msg string, err error) {
	// 	// 只能更改指针的值 SetInt
	// 	if v.Type().Kind() == reflect.Ptr {
	// 		if v.IsNil() {
	// 			return "", nil
	// 		}
	// 		vf := v.Elem()
	// 		log.Println("value: ", params)
	// 		if vf.Kind() == reflect.Int {
	// 			i, err := strconv.Atoi(params)
	// 			if err != nil {
	// 				return "", err
	// 			}
	// 			log.Println("value i: ", i, v.CanSet(), vf.CanSet())
	// 			if vf.CanSet() {
	// 				vf.SetInt(int64(i))
	// 			}
	// 		}
	// 	}

	// 	return "", nil
	// })

	SetValue("value", func(v reflect.Value, canSet bool, args string) (msg string, err error) {
		// 只能更改指针的值 SetInt

		log.Println("args: ", args, canSet)
		if canSet && v.Kind() == reflect.Int {
			i, err := strconv.Atoi(args)
			if err != nil {
				return "", err
			}
			v.SetInt(int64(i))
		}
		return "", nil
	})

	/////////////////////////////

	type Cart struct {
		Color int     `check:"color"`
		Tag   string  `check:"tag"`
		Name_ *string `check:"name"`
		Age_  *int    `check:"age"`
	}

	type Presion struct {
		Name   string  `check:"name"`
		Age    int     `check:"value:1000"`
		Name_  *string `check:"name"`
		Age_   *int    `check:"value:99"`
		f      func()
		Struct Cart `check:"-"`
		Cart   Cart
		CartP  *Cart
	}

	name_ := "Hell"
	age_ := 18
	p := Presion{Name: "H", Age: 19,
		Name_: &name_, Age_: &age_,
		f: func() {},
		Struct: Cart{
			Color: 10,
			Age_:  &age_,
		},
		Cart: Cart{
			Color: 10,
			Age_:  &age_,
		},
		CartP: &Cart{
			Color: 10,
			Tag:   "tttg",
		}}

	msg, err := Check(p)
	t.Log(msg)
	t.Log(err)

	fmt.Printf("\n%+v\n\n", *p.Age_)

}
