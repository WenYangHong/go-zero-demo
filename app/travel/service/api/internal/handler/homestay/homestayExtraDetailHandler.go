package homestay

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-mall/app/travel/service/api/internal/logic/homestay"
	"go-zero-mall/app/travel/service/api/internal/middleware"
	"go-zero-mall/app/travel/service/api/internal/svc"
	"go-zero-mall/app/travel/service/api/internal/types"
)

func HomestayExtraDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HomestayDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			middleware.ErrorJson(w, r, http.StatusBadRequest, "参数错误")
			return
		}

		l := homestay.NewHomestayExtraDetailLogic(r.Context(), svcCtx)
		resp, err := l.HomestayExtraDetail(&req)
		if err != nil {
			middleware.ErrorJson(w, r, http.StatusInternalServerError, "请求失败")
		} else {
			middleware.OkJson(w, r, resp)
		}
	}
}
