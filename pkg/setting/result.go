package setting

import "lyra-bot-web-crawlers/pkg/setting/e"

type Result[T any] struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Success bool   `json:"success"`
	Data    T      `json:"data"`
}

func SuccessOfData[T any](data T) Result[T] {
	return Result[T]{
		Status:  e.SUCCESS,
		Message: e.MsgFlags[e.SUCCESS],
		Success: true,
		Data:    data,
	}
}

func Success[T any]() Result[T] {
	return Result[T]{
		Status:  e.SUCCESS,
		Message: e.MsgFlags[e.SUCCESS],
		Success: true,
	}
}
