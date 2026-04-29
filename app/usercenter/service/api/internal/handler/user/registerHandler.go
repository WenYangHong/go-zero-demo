package user

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-mall/app/usercenter/service/api/internal/logic/user"
	"go-zero-mall/app/usercenter/service/api/internal/svc"
	"go-zero-mall/app/usercenter/service/api/internal/types"
	"go-zero-mall/pkg/xerr"
)

var validate = validator.New()

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			xerr.HandleError(w, r, xerr.NewErrCode(xerr.REUQEST_PARAM_ERROR))
			return
		}

		if err := validate.Struct(&req); err != nil {
			xerr.HandleError(w, r, xerr.NewErrCode(xerr.REUQEST_PARAM_ERROR))
			return
		}

		l := user.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			xerr.HandleError(w, r, err)
			return
		}
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}
