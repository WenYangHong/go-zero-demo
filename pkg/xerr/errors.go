package xerr

import (
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"regexp"
	"strconv"
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

var codeErrRegexp = regexp.MustCompile(`ErrCode:(\d+)`)

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

func CommonErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	httpx.WriteJson(w, http.StatusBadRequest, ErrorResponse{
		Code: 400,
		Msg:  "参数错误",
		Data: nil,
	})
}

func CommonError() error {
	return NewErrCode(SERVER_COMMON_ERROR)
}

func ParseCodeErr(err error) *CodeError {
	if err == nil {
		return nil
	}

	if codeErr, ok := err.(*CodeError); ok {
		return codeErr
	}

	matches := codeErrRegexp.FindStringSubmatch(err.Error())
	if len(matches) != 2 {
		return nil
	}

	errCode, parseErr := strconv.ParseUint(matches[1], 10, 32)
	if parseErr != nil || !IsCodeErr(uint32(errCode)) {
		return nil
	}

	return NewErrCode(uint32(errCode))
}

func FromRpcError(err error, defaultCode uint32, codeMapping map[uint32]uint32) error {
	if codeErr := ParseCodeErr(err); codeErr != nil {
		if mappedCode, ok := codeMapping[codeErr.GetErrCode()]; ok {
			return NewErrCode(mappedCode)
		}
	}

	return NewErrCode(defaultCode)
}

func HandleError(w http.ResponseWriter, r *http.Request, err error) {
	if codeErr, ok := err.(*CodeError); ok {
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
