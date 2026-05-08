package order

import (
	"go-zero-mall/pkg/xerr"
	"net/http"

	"go-zero-mall/app/order/service/api/internal/logic/order"
	"go-zero-mall/app/order/service/api/internal/svc"
	"go-zero-mall/app/order/service/api/internal/types"

	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var validate = validator.New()

// 创建民宿订单
func CreateHomestayOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateHomestayOrderReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		if err := validate.Struct(&req); err != nil {
			xerr.HandleError(w, r, xerr.NewErrCode(xerr.REUQEST_PARAM_ERROR))
			return
		}
		l := order.NewCreateHomestayOrderLogic(r.Context(), svcCtx)
		resp, err := l.CreateHomestayOrder(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
