package model

type Output struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Success(msg string, data any) *Output {
	return &Output{
		Code: 1,
		Msg:  msg,
		Data: data,
	}
}

func Fail(msg string) *Output {
	return &Output{
		Code: -1,
		Msg:  msg,
	}
}
