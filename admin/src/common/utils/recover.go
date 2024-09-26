package utils

import (
	"errors"
	"runtime/debug"

	"siteol.com/smart/src/common/log"
)

// RecoverWrap 公共的Recover函数
func RecoverWrap(inF func()) func() {
	return func() {
		defer func() {
			if r := recover(); r != nil {
				log.ErrorF("SYSTEM ACTION PANIC: %v, stack: %v", ToErr(r), string(debug.Stack()))
			}
		}()
		inF()
	}
}

// ToErr 将 any 类型, 转换为 error 类型
func ToErr(r any) error {
	var err error
	if r != nil {
		switch t := r.(type) {
		case string:
			err = errors.New(t)
		case error:
			err = t
		default:
			err = errors.New("Unknown error ")
		}
	}
	return err
}
