package homestay

import (
	"context"
	"github.com/jinzhu/copier"
	"go-zero-mall/pkg/xerr"
	"strconv"

	"go-zero-mall/app/travel/service/api/internal/svc"
	"go-zero-mall/app/travel/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayBusinessInfoWithHomeStayIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 根据民宿ID查询：房东信息
func NewHomestayBusinessInfoWithHomeStayIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayBusinessInfoWithHomeStayIdLogic {
	return &HomestayBusinessInfoWithHomeStayIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayBusinessInfoWithHomeStayIdLogic) HomestayBusinessInfoWithHomeStayId(req *types.HomestayBusinessInfoWithHomeStayIdReq) (resp *types.HomestayBusinessInfoResp, err error) {
	homestayId, _ := strconv.ParseInt(req.HomestayId, 10, 64)
	// 查询民宿获取房东 user_id
	homestay, err := l.svcCtx.HomestayModel.FindOne(l.ctx, homestayId)
	if homestay == nil {
		return nil, xerr.NewErrCode(xerr.USER_NOT_FOUND_ERROR)
	}
	if err != nil {
		return nil, err
	}
	// 获取userid
	userId := homestay.UserId
	homestayBusiness, err := l.svcCtx.HomestayBusinessModel.FindOneByUserId(l.ctx, userId)

	var homestayBusinessResp types.HomestayBusiness
	_ = copier.Copy(&homestayBusinessResp, homestayBusiness)
	homestayBusinessResp.Id = strconv.FormatInt(homestayBusiness.Id, 10)
	return &types.HomestayBusinessInfoResp{
		Code: xerr.OK,
		Msg:  "success",
		Info: homestayBusinessResp,
	}, nil
}
