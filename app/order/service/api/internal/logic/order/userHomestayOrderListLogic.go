package order

import (
	"context"
	"go-zero-mall/app/order/service/rpc/pb"
	"go-zero-mall/pkg/ctxdata"

	"go-zero-mall/app/order/service/api/internal/svc"
	"go-zero-mall/app/order/service/api/internal/types"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserHomestayOrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户订单列表
func NewUserHomestayOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserHomestayOrderListLogic {
	return &UserHomestayOrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserHomestayOrderListLogic) UserHomestayOrderList(req *types.UserHomestayOrderListReq) (resp *types.UserHomestayOrderListResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	list, err := l.svcCtx.OrderRpc.UserHomestayOrderList(l.ctx, &pb.UserHomestayOrderListReq{
		UserId:      userId,
		Page:        req.Page,
		PageSize:    req.PageSize,
		TraderState: req.TradeState,
		LastId:      req.LastId,
	})
	if err != nil {
		return nil, err
	}
	var returnList []types.UserHomestayOrderListView
	for _, v := range list.List {
		var item types.UserHomestayOrderListView
		err := copier.Copy(&item, v)
		if err != nil {
			return nil, err
		}
		returnList = append(returnList, item)
	}
	return &types.UserHomestayOrderListResp{
		List: returnList,
	}, nil
}
