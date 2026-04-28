package xerr

import (
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type ErrorResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type CodeError struct {
	errCode uint32
	errMsg  string
}

func (e *CodeError) GetErrCode() uint32 {
	return e.errCode
}

func (e *CodeError) GetErrMsg() string {
	return e.errMsg
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.errCode, e.errMsg)
}

func NewErrCodeMsg(errCode uint32, errMsg string) *CodeError {
	return &CodeError{errCode: errCode, errMsg: errMsg}
}

func NewErrCode(errCode uint32) *CodeError {
	return &CodeError{errCode: errCode, errMsg: MapErrMsg(errCode)}
}

func NewErrMsg(errMsg string) *CodeError {
	return &CodeError{errCode: SERVER_COMMON_ERROR, errMsg: errMsg}
}

func CommonError() error {
	return NewErrCode(SERVER_COMMON_ERROR)
}

func HandleError(w http.ResponseWriter, r *http.Request, err error) {
	fmt.Println("error:", err)

	if codeErr, ok := err.(*CodeError); ok {
		fmt.Println("error2:", codeErr)
		httpx.WriteJson(w, http.StatusBadRequest, ErrorResponse{
			Code: int(codeErr.errCode),
			Msg:  codeErr.errMsg,
			Data: nil,
		})
	} else {
		httpx.WriteJson(w, http.StatusInternalServerError, ErrorResponse{
			Code: int(SERVER_COMMON_ERROR),
			Msg:  "服务器错误",
			Data: nil,
		})
	}
}
