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

type UserHomestayOrderDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户订单明细
func NewUserHomestayOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserHomestayOrderDetailLogic {
	return &UserHomestayOrderDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserHomestayOrderDetailLogic) UserHomestayOrderDetail(req *types.UserHomestayOrderDetailReq) (resp *types.UserHomestayOrderDetailResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	detail, err := l.svcCtx.OrderRpc.HomestayOrderDetail(l.ctx, &pb.HomestayOrderDetailReq{
		UserId: userId,
		Sn:     req.Sn,
	})
	if err != nil {
		return nil, err
	}
	var returnDetail types.UserHomestayOrderDetailResp

	err = copier.Copy(&returnDetail, detail.HomestayOrder)
	if err != nil {
		return nil, err
	}
	return &returnDetail, nil
}
