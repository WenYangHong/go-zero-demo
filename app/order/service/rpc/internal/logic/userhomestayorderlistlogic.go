package logic

import (
	"context"
	"go-zero-mall/pkg/xerr"

	"go-zero-mall/app/order/service/rpc/internal/svc"
	"go-zero-mall/app/order/service/rpc/pb"

	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserHomestayOrderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserHomestayOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserHomestayOrderListLogic {
	return &UserHomestayOrderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户民宿订单
func (l *UserHomestayOrderListLogic) UserHomestayOrderList(in *pb.UserHomestayOrderListReq) (*pb.UserHomestayOrderListResp, error) {
	// 获取登录用户信息
	if in.UserId == 0 {
		// 用户未登录
		return nil, xerr.NewErrCode(xerr.USER_NOT_FOUND_ERROR)
	}
	whereBuilder := l.svcCtx.HomestayOrderModel.SelectBuilder().
		Where(squirrel.Eq{"user_id": in.UserId})

	if in.TraderState != -99 {
		// -99 表示全部订单
		whereBuilder = whereBuilder.Where(squirrel.Eq{"trade_state": in.TraderState})
	}
	list, err := l.svcCtx.HomestayOrderModel.FindPageListByPage(l.ctx, whereBuilder, in.Page, in.PageSize, "create_time desc")
	if err != nil {
		return nil, err
	}
	var returnList []*pb.HomestayOrder
	for _, homestayOrder := range list {
		var item pb.HomestayOrder
		err := copier.Copy(&item, homestayOrder)
		if err != nil {
			return nil, err
		}
		returnList = append(returnList, &item)
	}
	return &pb.UserHomestayOrderListResp{
		List: returnList,
	}, nil
}
