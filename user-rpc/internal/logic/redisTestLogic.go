package logic

import (
	"context"
	"fmt"
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
	// 读取
	rdb := l.svcCtx.Redis
	add, _ := rdb.Zadd("leaderboard", 1500, "player:alice")
	_, _ = rdb.Zadd("leaderboard2", 1480, "player:alice")
	_, _ = rdb.Zadd("leaderboard3", 1440, "player:alice")
	err := rdb.Setex("session:abc123", "userId:42", 3600)

	fmt.Println("redis_test_add_key", add)
	// 临界区
	if err != nil {
		return &pb.CommonResp{
			Code:    200,
			Message: "缓存没了",
		}, nil
	}

	return &pb.CommonResp{
		Code:    200,
		Message: "成功",
	}, nil
}
