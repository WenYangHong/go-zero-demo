package logic

import (
	"context"
	"fmt"
	model2 "go-zero-mall/app/order/service/model"
	"go-zero-mall/app/travel/service/model"
	"go-zero-mall/pkg/mq"
	"go-zero-mall/pkg/tool"
	"go-zero-mall/pkg/uniqueid"
	"go-zero-mall/pkg/xerr"
	"time"

	"go-zero-mall/app/order/service/rpc/internal/svc"
	"go-zero-mall/app/order/service/rpc/pb"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateHomestayOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateHomestayOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateHomestayOrderLogic {
	return &CreateHomestayOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 民宿下订单
func (l *CreateHomestayOrderLogic) CreateHomestayOrder(in *pb.CreateHomestayOrderReq) (*pb.CreateHomestayOrderResp, error) {
	// 民宿信息 - look使用 rpc 我这里使用 model，确实发现了问题
	homestay, _ := l.svcCtx.HomestayModel.FindOne(l.ctx, in.HomestayId)
	valid, err := l.orderParamsValid(in, homestay)

	if !valid || err != nil {
		return nil, err
	}

	// 准备订单入库
	order := new(model2.HomestayOrder)
	order.Sn = uniqueid.GenSn(uniqueid.SN_PREFIX_HOMESTAY_ORDER)
	order.UserId = in.UserId
	order.HomestayId = in.HomestayId
	order.Title = homestay.Title
	order.SubTitle = homestay.SubTitle
	order.Cover = homestay.Banner
	order.Info = homestay.Info
	order.PeopleNum = homestay.PeopleNum
	order.LivePeopleNum = in.LivePeopleNum // 实际入住人数
	order.RowType = 0
	order.LiveStartDate = time.Unix(in.LiveStartTime, 0)
	order.LiveEndDate = time.Unix(in.LiveEndTime, 0)
	// 入住天数
	liveDays := int64(order.LiveEndDate.Sub(order.LiveStartDate).Seconds() / 86400) // Stayed a few days in total
	if in.IsFood {
		order.NeedFood = model2.HomestayOrderNeedFoodYes
		// 计算食物费用
		order.FoodTotalPrice = homestay.FoodPrice * liveDays * in.LivePeopleNum
	}
	order.FoodPrice = homestay.FoodPrice
	order.FoodInfo = homestay.FoodInfo
	order.HomestayPrice = homestay.HomestayPrice
	order.HomestayTotalPrice = homestay.HomestayPrice * liveDays
	order.OrderTotalPrice = order.HomestayTotalPrice + order.FoodTotalPrice
	order.MarketHomestayPrice = homestay.MarketHomestayPrice
	order.HomestayBusinessId = homestay.HomestayBusinessId
	order.HomestayUserId = homestay.UserId
	order.TradeState = model2.HomestayOrderTradeStateWaitPay
	order.TradeCode = tool.Krand(8, tool.KC_RAND_KIND_ALL)
	order.Remark = in.Remark
	// 创建订单
	_, err = l.svcCtx.HomestayOrderModel.Insert(l.ctx, nil, order)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Order Database Exception order : %+v , err: %v", order, err)
	}
	// 创建延迟队列，30分钟后关闭未支付订单 - RabbitMQ + 死信队列实现
	mqErr := mq.PublishDelayOrder(order.Sn)
	if mqErr != nil {
		return nil, mqErr
	}
	return &pb.CreateHomestayOrderResp{
		Sn: order.Sn,
	}, nil
}

func (l *CreateHomestayOrderLogic) orderParamsValid(in *pb.CreateHomestayOrderReq, homestay *model.Homestay) (bool, error) {
	fmt.Println(" -- orderQuery")

	// 获取登录用户信息
	if in.UserId == 0 {
		// 用户未登录
		return false, xerr.NewErrCode(xerr.USER_NOT_FOUND_ERROR)
	}
	// 根据 homestayId 查询民宿信息
	if homestay == nil {
		// 找不到民宿信息
		return false, xerr.NewErrCode(xerr.HOMESTAY_DETAIL_FAIL)
	}
	if homestay.RowState != model.HomestayUpStatus {
		// 民宿下架
		return false, xerr.NewErrCode(xerr.HOMESTAY_DETAIL_FAIL)
	}
	// 判断入住时间是否合法
	if in.LiveStartTime >= in.LiveEndTime || in.LiveStartTime <= 0 {
		return false, xerr.NewErrCode(xerr.ORDER_LIVE_TIME_ERROR)
	}
	now := time.Now().Unix()
	if in.LiveStartTime < now || in.LiveEndTime < now {
		return false, xerr.NewErrCode(xerr.ORDER_LIVE_TIME_ERROR)
	}
	if in.LivePeopleNum <= 0 || in.LivePeopleNum > homestay.PeopleNum {
		// 人数不足 或 超过民宿可住人数上限
		return false, xerr.NewErrCode(xerr.ORDER_LIVE_PEOPLE_NUM_ERROR)
	}
	// 根据 homestayId + liveStartTime + liveEndTime + trade_state = [0,1] 查询该时间段内是否有订单
	hasOrder, hasOrderError := l.svcCtx.HomestayOrderModel.FindOneByHomestayIdLiveTimeTradeState(l.ctx, in.UserId, in.HomestayId, in.LiveStartTime, in.LiveEndTime, []int64{model2.HomestayOrderTradeStateWaitPay, model2.HomestayOrderTradeStateWaitUse})
	if hasOrderError != nil {
		return false, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Order Database Exception order : %+v , err: %v", in, hasOrderError)
	}
	if hasOrder {
		// 订单已存在
		return false, xerr.NewErrCode(xerr.ORDER_EXIST_ERROR)
	}
	return true, nil
}
