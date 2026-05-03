package homestay

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-mall/app/travel/service/api/internal/logic/homestay"
	"go-zero-mall/app/travel/service/api/internal/middleware"
	"go-zero-mall/app/travel/service/api/internal/svc"
	"go-zero-mall/app/travel/service/api/internal/types"
)

func HomestayCommonListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HomestayCommonListReq
		if err := httpx.Parse(r, &req); err != nil {
			middleware.ErrorJson(w, r, http.StatusBadRequest, "参数错误")
			return
		}

		l := homestay.NewHomestayCommonListLogic(r.Context(), svcCtx)
		resp, err := l.HomestayCommonList(&req)
		if err != nil {
			middleware.ErrorJson(w, r, http.StatusInternalServerError, "请求失败")
		} else {
			middleware.OkJson(w, r, resp)
		}
	}
}
