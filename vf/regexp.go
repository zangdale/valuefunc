package vf

import (
	"errors"
	"reflect"
	"regexp"

	"github.com/zangdale/valuefunc"
)

var (
	ErrRegexpArgsIsNil    = errors.New("regexp arg is nil")
	ErrRegexpValueNoMatch = errors.New("not match regexp")
)

var RegexpValueFunc = valuefunc.New()

func init() {
	err := RegexpValueFunc.Add("regexp", func(v reflect.Value, args string) (msg string, err error) {
		if args == "" {
			return "", ErrRegexpArgsIsNil
		}
		compile := regexp.MustCompile(args)

		switch v.Type().Kind() {
		case reflect.String:
			if !compile.MatchString(v.String()) {
				return args, ErrRegexpValueNoMatch
			}
		default:
			// todo
		}

		return
	})
	if err != nil {
		panic(err)
	}
}
