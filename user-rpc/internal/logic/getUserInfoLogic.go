package logic

import (
	"context"
	"errors"
	"fmt"
	"go-zero-demo/user-rpc/internal/model"

	"go-zero-demo/user-rpc/internal/svc"
	"go-zero-demo/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
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
	// todo: add your logic here and delete this line

	if in.GetId() == 0 {
		return nil, errors.New("不可以缺少ID")
	}
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)

	count, _ := l.svcCtx.UserModel.CountActiveUsers(l.ctx)
	fmt.Println("自定义SQL：count", count)

	if errors.Is(err, model.ErrNotFound) {
		return nil, errors.New("未查询到用户信息")
	}
	return &pb.GetUserInfoResp{
		Id:       user.Id,
		Nickname: user.Username,
	}, nil
}
