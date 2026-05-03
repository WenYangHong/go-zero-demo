package homestay

import (
	"context"

	"go-zero-mall/app/travel/service/api/internal/svc"
	"go-zero-mall/app/travel/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayExtraDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// homestay room eatra detail
func NewHomestayExtraDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayExtraDetailLogic {
	return &HomestayExtraDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayExtraDetailLogic) HomestayExtraDetail(req *types.HomestayDetailReq) (resp *types.HomestayDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
