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

type HomestayCommentCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建评价
func NewHomestayCommentCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayCommentCreateLogic {
	return &HomestayCommentCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayCommentCreateLogic) HomestayCommentCreate(req *types.HomestayCommentCreateReq) (resp *types.HomestayCommentCreateResp, err error) {
	// todo: 2个问题 => 1. content 长度限制问题 2.同一个用户+同一间民宿多次评价
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
	homestayComment := new(model.HomestayComment)
	homestayComment.AvgStar = (req.ServiceStar + req.LocationStar + req.CleanlinessStar + req.ValueStar) / 4
	homestayComment.ServiceStar = req.ServiceStar
	homestayComment.LocationStar = req.LocationStar
	homestayComment.CleanlinessStar = req.CleanlinessStar
	homestayComment.ValueStar = req.ValueStar
	homestayComment.Content = req.Content
	homestayComment.UserId = userId
	homestayComment.HomestayId = homestayId
	err = l.svcCtx.HomestayCommentModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
		l.svcCtx.HomestayCommentModel.Insert(l.ctx, session, homestayComment)
		return nil
	})
	if err != nil {
		return nil, xerr.NewErrCode(xerr.COMMENT_CREATE_FAIL)
	}
	return &types.HomestayCommentCreateResp{
		Code: xerr.OK,
		Msg:  "评价创建成功",
	}, nil
}
