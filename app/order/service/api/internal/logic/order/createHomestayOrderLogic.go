package order

import (
	"context"
	"strconv"

	"go-zero-mall/app/order/service/api/internal/svc"
	"go-zero-mall/app/order/service/api/internal/types"
	"go-zero-mall/app/order/service/rpc/order"
	"go-zero-mall/pkg/ctxdata"
	"go-zero-mall/pkg/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateHomestayOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建民宿订单
func NewCreateHomestayOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateHomestayOrderLogic {
	return &CreateHomestayOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateHomestayOrderLogic) CreateHomestayOrder(req *types.CreateHomestayOrderReq) (resp *types.HomestayOrderInfoResp, err error) {
	homestayId, _ := strconv.ParseInt(req.HomestayId, 10, 64)
	userId := ctxdata.GetUidFromCtx(l.ctx)
	isFood := false
	if req.IsFood == 1 {
		isFood = true
	}

	homestayOrder, err := l.svcCtx.OrderRpc.CreateHomestayOrder(l.ctx, &order.CreateHomestayOrderReq{
		HomestayId:    homestayId,
		IsFood:        isFood,
		LiveEndTime:   req.LiveEndTime,
		LiveStartTime: req.LiveStartTime,
		UserId:        userId,
		LivePeopleNum: req.PeopleNum,
		Remark:        req.Remark,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ORDER_CREATE_FAIL), "rpc CreateHomestayOrder err: %v", err)
	}
	return &types.HomestayOrderInfoResp{
		Sn: homestayOrder.Sn,
	}, nil
}
