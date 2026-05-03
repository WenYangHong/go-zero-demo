package middleware

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
	body       bytes.Buffer
	written    bool
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	if !w.written {
		w.ResponseWriter.WriteHeader(statusCode)
	}
}

func (w *responseWriter) Write(b []byte) (int, error) {
	return w.body.Write(b)
}

func (w *responseWriter) flush() {
	if w.body.Len() > 0 {
		w.written = true
		w.ResponseWriter.Write(w.body.Bytes())
	}
}

func ResponseMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		defer func() {
			if err := recover(); err != nil {
				logx.Errorf("panic recovered: %v", err)
				httpx.WriteJson(w, http.StatusInternalServerError, Response{
					Code: 500,
					Msg:  "服务器内部错误",
					Data: nil,
				})
			}
		}()

		next(rw, r)

		if rw.body.Len() == 0 {
			return
		}

		var tmp map[string]interface{}
		if json.Unmarshal(rw.body.Bytes(), &tmp) == nil && tmp["code"] != nil {
			rw.flush()
			return
		}

		if rw.statusCode >= 400 {
			var result map[string]interface{}
			json.Unmarshal(rw.body.Bytes(), &result)
			msg, _ := result["msg"].(string)
			if msg == "" {
				msg = "请求失败"
			}
			httpx.WriteJson(w, rw.statusCode, Response{
				Code: rw.statusCode,
				Msg:  msg,
				Data: nil,
			})
		} else {
			var data interface{}
			json.Unmarshal(rw.body.Bytes(), &data)
			httpx.WriteJson(w, http.StatusOK, Response{
				Code: 200,
				Msg:  "success",
				Data: data,
			})
		}
	}
}

func OkJson(w http.ResponseWriter, r *http.Request, data interface{}) {
	httpx.WriteJson(w, http.StatusOK, Response{
		Code: 200,
		Msg:  "success",
		Data: data,
	})
}

func ErrorJson(w http.ResponseWriter, r *http.Request, code int, msg string) {
	httpx.WriteJson(w, code, Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}
