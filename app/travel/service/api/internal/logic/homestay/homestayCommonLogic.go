package homestay

import (
	"context"

	"go-zero-mall/app/travel/service/api/internal/svc"
	"go-zero-mall/app/travel/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayCommonLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建评价
func NewHomestayCommonLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayCommonLogic {
	return &HomestayCommonLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayCommonLogic) HomestayCommon(req *types.HomestayCommonCreateReq) (resp *types.HomestayCommonCreateResp, err error) {
	// todo: add your logic here and delete this line

	return
}
