package homestay

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-mall/app/travel/service/model"
	"go-zero-mall/pkg/ctxdata"
	"go-zero-mall/pkg/xerr"
	"strconv"

	"go-zero-mall/app/travel/service/api/internal/svc"
	"go-zero-mall/app/travel/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayCommonCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建评价
func NewHomestayCommonCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayCommonCreateLogic {
	return &HomestayCommonCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayCommonCreateLogic) HomestayCommonCreate(req *types.HomestayCommonCreateReq) (resp *types.HomestayCommonCreateResp, err error) {
	// 用户ID
	userId := ctxdata.GetUidFromCtx(l.ctx)
	if userId == 0 {
		return nil, xerr.NewErrCode(xerr.USER_NOT_FOUND_ERROR)
	}
	// 先查询一下民宿的情况
	homestayId, err := strconv.ParseInt(req.HomestayId, 10, 64)
	homestay, err := l.svcCtx.HomestayModel.FindOne(l.ctx, homestayId)
	if err != nil {
		return nil, xerr.NewErrCode(xerr.REUQEST_PARAM_ERROR)
	}
	if homestay == nil {
		return nil, xerr.NewErrCode(xerr.REUQEST_PARAM_ERROR)
	}
	// 计算平均分数
	homestayCommon := new(model.HomestayComment)
	homestayCommon.AvgStar = (req.ServiceStar + req.LocationStar + req.CleanlinessStar + req.ValueStar) / 4
	homestayCommon.ServiceStar = req.ServiceStar
	homestayCommon.LocationStar = req.LocationStar
	homestayCommon.CleanlinessStar = req.CleanlinessStar
	homestayCommon.ValueStar = req.ValueStar
	homestayCommon.Content = req.Content
	homestayCommon.UserId = userId
	homestayCommon.HomestayId = homestayId
	err = l.svcCtx.HomestayCommentModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
		l.svcCtx.HomestayCommentModel.Insert(l.ctx, session, homestayCommon)
		return nil
	})
	if err != nil {
		return nil, xerr.NewErrCode(xerr.COMMENT_CREATE_FAIL)
	}
	return &types.HomestayCommonCreateResp{
		Code: xerr.OK,
		Msg:  "评价创建成功",
	}, nil
}
