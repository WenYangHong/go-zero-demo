package logic

import (
	"context"

	"go-zero-demo/user-rpc/internal/svc"
	"go-zero-demo/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RedisTestLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRedisTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RedisTestLogic {
	return &RedisTestLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RedisTestLogic) RedisTest(in *pb.RedisTestReq) (*pb.CommonResp, error) {
	// todo: add your logic here and delete this line
	l.svcCtx.Redis.Setex("session:abc123", "userId:42", 3600)

	// 读取
	val, err := l.svcCtx.Redis.Get("session:abc123")

	if err != nil {
		return &pb.CommonResp{
			Code:    200,
			Message: "缓存没了",
		}, nil
	}
	return &pb.CommonResp{
		Code:    200,
		Message: val,
	}, nil
}
