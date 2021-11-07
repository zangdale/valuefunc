package vf

import (
	"errors"
	"os"
	"reflect"

	"github.com/zangdale/valuefunc"
)

var (
	ErrEnvArgsIsNil = errors.New("no have env param")
)

var EnvValueFunc = valuefunc.New()

func init() {
	err := EnvValueFunc.SetValue("env", func(v reflect.Value, canSet bool, args string) (msg string, err error) {
		if args == "" {
			return "", ErrEnvArgsIsNil
		}

		if canSet {
			s := os.Getenv(args)
			switch v.Type().Kind() {
			case reflect.String:
				v.SetString(s)
			default:
				// todo
			}
		}
		return
	})

	if err != nil {
		panic(err)
	}

}
