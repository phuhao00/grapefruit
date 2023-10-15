package err

type ErrorMessage struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e *ErrorMessage) SetMsg(errMsg string, code int) {
	e.Code = code
	e.Msg = errMsg
}

func (e *ErrorMessage) GetMsg() *ErrorMessage {
	return e
}

func DefaultSucceedErrMsg() ErrorMessage {
	return ErrorMessage{
		Code: 200,
		Msg:  "ok !",
	}
}

func DefaultFailedErrMsg() ErrorMessage {
	return ErrorMessage{
		Code: 0,
		Msg:  "failed !",
	}
}

func DefaultSuccessErr(msg string) ErrorMessage {
	return ErrorMessage{
		Code: 200,
		Msg:  msg,
	}
}

func DefaultFailedErr(msg string) ErrorMessage {
	return ErrorMessage{
		Code: 0,
		Msg:  msg,
	}
}

func ErrorMsg(msg string, code int) ErrorMessage {
	return ErrorMessage{
		Code: code,
		Msg:  msg,
	}
}
