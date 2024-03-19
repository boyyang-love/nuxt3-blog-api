package errorx

const defaultCode = 0

type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
type CodeErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CodeErrorWithStatus struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CodeErrorWithStatusResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewCodeError(code int, msg string) error {
	return &CodeError{Code: code, Msg: msg}
}
func NewDefaultError(msg string) error {
	return NewCodeError(defaultCode, msg)
}
func NewCodeErrorWithStatus(code int, msg string) error {
	return &CodeErrorWithStatus{Code: code, Msg: msg}
}

func NewDefaultErrorWithStatus(msg string) error {
	return NewCodeErrorWithStatus(defaultCode, msg)
}
func (e *CodeError) Error() string {
	return e.Msg
}
func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}
func (e *CodeErrorWithStatus) Error() string {
	return e.Msg
}
func (e *CodeErrorWithStatus) Data() *CodeErrorWithStatusResponse {
	return &CodeErrorWithStatusResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}
