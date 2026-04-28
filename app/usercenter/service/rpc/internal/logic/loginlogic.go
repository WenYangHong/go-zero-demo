package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-mall/app/usercenter/service/rpc/internal/model"
	"go-zero-mall/app/usercenter/service/rpc/internal/svc"
	"go-zero-mall/app/usercenter/service/rpc/pb"
	"go-zero-mall/app/usercenter/service/rpc/usercenter"
	"go-zero-mall/pkg/tool"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	// 这里不做多方式登录的判断了，直接根据手机号进行登录
	user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.AuthKey)
	if err != nil || errors.Is(err, model.ErrNotFound) {
		return nil, err
	}
	if user == nil {
		return nil, err
	}
	// 再根据密码进行判断
	if tool.Md5ByString(in.Password) != user.Password {
		return nil, errors.New("密码错误")
	}
	// 获取jwt token
	generate := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	token, err := generate.GenerateToken(&usercenter.GenerateTokenReq{
		UserId: user.Id,
	})
	if err != nil {
		return nil, errors.New("登录失败")
	}
	return &usercenter.LoginResp{
		AccessToken:  token.AccessToken,
		AccessExpire: token.AccessExpire,
		RefreshAfter: token.RefreshAfter,
	}, nil
}
