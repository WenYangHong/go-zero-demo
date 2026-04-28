package user

import (
	"go-zero-mall/pkg/xerr"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-mall/app/usercenter/service/api/internal/logic/user"
	"go-zero-mall/app/usercenter/service/api/internal/svc"
	"go-zero-mall/app/usercenter/service/api/internal/types"
)

// login
func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			xerr.HandleError(w, r, err)
			return
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
