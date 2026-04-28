package user

import (
	"context"
	"github.com/jinzhu/copier"
	"go-zero-mall/app/usercenter/service/rpc/pb"
	"go-zero-mall/pkg/ctxdata"

	"go-zero-mall/app/usercenter/service/api/internal/svc"
	"go-zero-mall/app/usercenter/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// get user info
func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)

	user, err := l.svcCtx.UserCenterRpc.GetUserInfo(l.ctx, &pb.GetUserInfoReq{
		Id: userId,
	})

	if user == nil {
		return nil, err
	}

	var respUser types.User
	_ = copier.Copy(&respUser, user.User)

	return &types.UserInfoResp{
		UserInfo: respUser,
	}, nil
}
