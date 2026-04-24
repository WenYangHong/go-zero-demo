package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-demo/user-rpc/internal/model"

	"go-zero-demo/user-rpc/internal/svc"
	"go-zero-demo/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateNewUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateNewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateNewUserLogic {
	return &CreateNewUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateNewUserLogic) CreateNewUser(in *pb.CreateNewUserReq) (*pb.CommonResp, error) {
	// todo: add your logic here and delete this line
	logx.Info("CreateNewUser:", in)
	// 非事务型
	// result, err := l.svcCtx.UserModel.Insert(l.ctx, &model.User{
	// 	Username: in.Username,
	// 	Password: in.Password,
	// 	Mobile:   in.Mobile,
	// })

	// 事务
	err := l.svcCtx.UserModel.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		result, err := l.svcCtx.UserModel.Insert(ctx, &model.User{
			Username: in.Username,
			Password: in.Password,
			Mobile:   in.Mobile,
		})
		if err != nil {
			return err // 自动回滚
		}
		rows, err := result.RowsAffected()
		if rows != 1 {
			return err
		}
		return nil // 提交
	})

	if err != nil {
		logx.Errorf("RowsAffected error: %v", err)
		return &pb.CommonResp{Code: 500, Message: "创建用户失败"}, nil
	}

	return &pb.CommonResp{Code: 200, Message: "创建用户成功"}, nil
}
