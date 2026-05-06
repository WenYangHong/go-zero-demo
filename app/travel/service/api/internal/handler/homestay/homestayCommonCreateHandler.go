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
func HomestayCommonCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HomestayCommonCreateReq
		if err := httpx.Parse(r, &req); err != nil {
			xerr.HandleError(w, r, xerr.NewErrMsg(err.Error()))
			return
		}
		if err := validate.Struct(&req); err != nil {
			xerr.HandleError(w, r, xerr.NewErrCode(xerr.REUQEST_PARAM_ERROR))
			return
		}
		l := homestay.NewHomestayCommonCreateLogic(r.Context(), svcCtx)
		resp, err := l.HomestayCommonCreate(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
