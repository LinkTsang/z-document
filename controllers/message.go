package controllers

type MessageCode int32
type MessageStr string

func h(code MessageCode, message string) map[string]interface{} {
	return map[string]interface{}{
		"code":    code,
		"message": message,
	}
}

const (
	CodeSuccess MessageCode = iota

	CodeBadRequest    = 40000
	CodeLoginRequired = 40001

	CodeFailed         = 50000
	CodeNotImplemented = 50001
)

const (
	MsgSuccess = "MSG_SUCCESS"

	MsgBadRequest    = "MSG_BAD_REQUEST"
	MsgLoginRequired = "MSG_LOGIN_REQUIRED"

	MsgFailed         = "MSG_FAILED"
	MsgNotImplemented = "MSG_NOT_IMPLEMENTED"
)

var (
	RespSuccess        = h(CodeSuccess, MsgSuccess)
	RespBadRequest     = h(CodeBadRequest, MsgBadRequest)
	RespLoginRequired  = h(CodeLoginRequired, MsgLoginRequired)
	RespFailed         = h(CodeFailed, MsgFailed)
	RespNotImplemented = h(CodeNotImplemented, MsgNotImplemented)
)

type Message struct {
	code    MessageCode
	message string
}
