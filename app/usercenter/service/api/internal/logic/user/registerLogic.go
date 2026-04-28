package user

import (
	"context"
	"go-zero-mall/app/usercenter/service/api/internal/svc"
	"go-zero-mall/app/usercenter/service/api/internal/types"
	"go-zero-mall/app/usercenter/service/rpc/pb"
	"go-zero-mall/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// register
func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	register, err := l.svcCtx.UserCenterRpc.Register(l.ctx, &pb.RegisterReq{
		Mobile:   req.Mobile,
		Password: req.Password,
	})
	if err != nil {
		return nil, xerr.CommonError()
	}
	return &types.RegisterResp{
		AccessToken:  register.AccessToken,
		AccessExpire: register.AccessExpire,
		RefreshAfter: register.RefreshAfter,
	}, nil
}
