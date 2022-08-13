package utils

import (
	"errors"
	"gateway/pkg/logging"
)

// 包装错误
func UserError(err error) {
	if err != nil {
		err = errors.New("userService | err: " + err.Error())
		logging.Info(err)
		panic(err)
	}
}

func RecordError(err error) {
	if err != nil {
		err = errors.New("recordService | err: " + err.Error())
		logging.Info(err)
		panic(err)
	}
}
