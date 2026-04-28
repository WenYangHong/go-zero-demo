package middleware

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var validate = validator.New()

type ErrorResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ValidateRequest(req interface{}) error {
	return validate.Struct(req)
}

func ErrorHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	}
}

func WriteError(w http.ResponseWriter, r *http.Request, err error) {
	httpx.WriteJson(w, http.StatusBadRequest, ErrorResponse{
		Code: 400,
		Msg:  "参数错误",
		Data: nil,
	})
}