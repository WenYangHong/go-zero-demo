package logic

import (
	"context"
	"go-zero-mall/pkg/xerr"

	"go-zero-mall/app/order/service/rpc/internal/svc"
	"go-zero-mall/app/order/service/rpc/pb"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayOrderDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomestayOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayOrderDetailLogic {
	return &HomestayOrderDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 民宿订单详情
func (l *HomestayOrderDetailLogic) HomestayOrderDetail(in *pb.HomestayOrderDetailReq) (*pb.HomestayOrderDetailResp, error) {
	detail, err := l.svcCtx.HomestayOrderModel.FindOneBySn(l.ctx, in.Sn)
	if err != nil {
		return nil, err
	}
	if detail == nil {
		// 订单不存在
		return nil, xerr.NewErrCode(xerr.ORDER_NOT_EXIST_ERROR)
	}

	// 判断登录人 和 订单人是否是同一个
	if detail.UserId != in.UserId {
		return nil, xerr.NewErrCode(xerr.ORDER_USER_NOT_MATCH_ERROR)
	}
	var returnDetail pb.HomestayOrder
	err = copier.Copy(&returnDetail, detail)
	if err != nil {
		return nil, xerr.NewErrCode(xerr.ORDER_EXIST_ERROR)
	}
	return &pb.HomestayOrderDetailResp{
		HomestayOrder: &returnDetail,
	}, nil
}
