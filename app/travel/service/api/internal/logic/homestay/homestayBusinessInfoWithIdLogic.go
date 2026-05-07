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

type HomestayBusinessInfoWithIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 根据房东信息ID查询：房东信息
func NewHomestayBusinessInfoWithIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayBusinessInfoWithIdLogic {
	return &HomestayBusinessInfoWithIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayBusinessInfoWithIdLogic) HomestayBusinessInfoWithId(req *types.HomestayBusinessInfoWithIdReq) (resp *types.HomestayBusinessInfoResp, err error) {
	id, _ := strconv.ParseInt(req.Id, 10, 64)
	homestayBusiness, err := l.svcCtx.HomestayBusinessModel.FindOne(l.ctx, id)
	if err != nil {
		return &types.HomestayBusinessInfoResp{}, xerr.NewErrCode(xerr.REUQEST_PARAM_ERROR)
	}
	if homestayBusiness == nil {
		return &types.HomestayBusinessInfoResp{}, nil
	}
	var homestayBusinessResp types.HomestayBusiness
	_ = copier.Copy(&homestayBusinessResp, homestayBusiness)
	homestayBusinessResp.Id = strconv.FormatInt(homestayBusiness.Id, 10)
	return &types.HomestayBusinessInfoResp{
		Code: xerr.OK,
		Msg:  "success",
		Info: homestayBusinessResp,
	}, nil
}
