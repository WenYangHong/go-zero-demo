package logic

import (
	"context"
	"github.com/pkg/errors"
	"go-zero-mall/app/usercenter/service/rpc/internal/model"
	"go-zero-mall/app/usercenter/service/rpc/usercenter"
	"go-zero-mall/pkg/xerr"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-mall/app/usercenter/service/rpc/internal/svc"
	"go-zero-mall/app/usercenter/service/rpc/pb"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_NOT_FOUND_ERROR), "userId: %d", in.Id)
	}

	if user == nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_NOT_FOUND_ERROR), "userId: %d", in.Id)
	}
	var respUser usercenter.User
	_ = copier.Copy(&respUser, user) // 和java BeanUtils.do2bo4List 一样的效果
	return &pb.GetUserInfoResp{
		User: &respUser,
	}, nil
}
