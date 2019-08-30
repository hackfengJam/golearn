package services

// App内错误，内部传递时使用Code
type Error struct {
	Code string
}

func (e Error) Error() string {
	return e.Code
}

func NewError(code string) error {
	if code == "" {
		code = "forum.system.UnknownError"
	}
	return &Error{Code: code}
}

const (
	// 所有错误代码命名 系统-模块-错误码
	ErrRequestInvalid = "enigma.general.RequestInvalid" // 请求参数错误（客户端请求）

	ErrAuthRequired = "enigma.auth.AuthRequired" // 请先登录
)
