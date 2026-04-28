package user

import (
	"context"
	"go-zero-mall/app/usercenter/service/rpc/pb"
	"go-zero-mall/pkg/xerr"

	"go-zero-mall/app/usercenter/service/api/internal/svc"
	"go-zero-mall/app/usercenter/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// login
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	loginResp, err := l.svcCtx.UserCenterRpc.Login(l.ctx, &pb.LoginReq{
		AuthType: "",
		AuthKey:  req.Mobile,
		Password: req.Password,
	})
	// 如果登录失败的话，现在直接返回错误，应该根据code码对错误类型进行归纳

	if err != nil {
		return nil, xerr.CommonError()
	}
	return &types.LoginResp{
		AccessToken:  loginResp.AccessToken,
		AccessExpire: loginResp.AccessExpire,
		RefreshAfter: loginResp.RefreshAfter,
	}, err
}
