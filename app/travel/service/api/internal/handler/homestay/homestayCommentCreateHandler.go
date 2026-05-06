package homestay

import (
	"github.com/go-playground/validator/v10"
	"go-zero-mall/pkg/xerr"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-mall/app/travel/service/api/internal/logic/homestay"
	"go-zero-mall/app/travel/service/api/internal/svc"
	"go-zero-mall/app/travel/service/api/internal/types"
)

var validate = validator.New()

// 创建评价
func HomestayCommentCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HomestayCommentCreateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		if err := validate.Struct(&req); err != nil {
			xerr.HandleError(w, r, xerr.NewErrCode(xerr.REUQEST_PARAM_ERROR))
			return
		}
		l := homestay.NewHomestayCommentCreateLogic(r.Context(), svcCtx)
		resp, err := l.HomestayCommentCreate(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
