package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-mall/app/common/service/api/internal/logic"
	"go-zero-mall/app/common/service/api/internal/svc"
)

func UploadFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 最多 32 MB 不落盘，超出部分自动落盘
		if err := r.ParseMultipartForm(32 << 20); err != nil {
			httpx.Error(w, err)
			return
		}

		file, header, err := r.FormFile("file")
		if err != nil {
			httpx.Error(w, err)
			return
		}
		defer file.Close()

		// 通过读取前 512 字节检测 MIME 类型
		buf := make([]byte, 512)
		n, _ := file.Read(buf)
		mimeType := http.DetectContentType(buf[:n])
		fmt.Println("mimeType", mimeType)
		// if !isAllowedType(mimeType) {
		// 	httpx.Error(w, fmt.Errorf("不支持的文件类型: %s", mimeType))
		// 	return
		// }
		file.Seek(0, io.SeekStart)

		l := logic.NewUploadFileLogic(r.Context(), svcCtx)
		if svcCtx.Config.FileUploadType == "local" {
			resp, err := l.UploadFile(file, header)
			if err != nil {
				httpx.Error(w, err)
				return
			}
			httpx.OkJson(w, resp)
		} else {
			resp, err := l.UploadFileS3(file, header)
			if err != nil {
				httpx.Error(w, err)
				return
			}
			httpx.OkJson(w, resp)
		}

	}
}
