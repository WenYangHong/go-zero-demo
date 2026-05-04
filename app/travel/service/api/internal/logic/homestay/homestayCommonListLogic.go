package homestay

import (
	"context"
	"fmt"
	"go-zero-mall/app/travel/service/api/internal/svc"
	"go-zero-mall/app/travel/service/api/internal/types"
	"strconv"

	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayCommonListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// homestay room common list
func NewHomestayCommonListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayCommonListLogic {
	return &HomestayCommonListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayCommonListLogic) HomestayCommonList(req *types.HomestayCommonListReq) (resp *types.HomestayCommonListResp, err error) {
	whereBuilder := l.svcCtx.HomestayCommentModel.SelectBuilder().Where(squirrel.Eq{
		"homestay_id": req.Id,
		"del_state":   0,
	})
	commonList, err := l.svcCtx.HomestayCommentModel.FindPageListByPage(l.ctx, whereBuilder, req.Page, req.PageSize, "create_time desc")
	if len(commonList) == 0 {
		return &types.HomestayCommonListResp{
			List: []types.HomestayCommonResp{},
		}, nil
	}
	// 取出所有的userid
	var userIdList []int64
	for _, v := range commonList {
		userIdList = append(userIdList, v.UserId)
	}

	userBuilder := l.svcCtx.UsersModel.SelectBuilder().Where(squirrel.Eq{
		"id": userIdList,
	})
	userInfo, err := l.svcCtx.UsersModel.FindAll(l.ctx, userBuilder, "id desc")
	if err != nil {
		fmt.Printf("[ERROR] FindAll users failed: %v, userIdList: %v\n", err, userIdList)
	}
	fmt.Printf("[DEBUG] FindAll users result: count=%d, userIdList=%v\n", len(userInfo), userIdList)

	userIdNameMap := make(map[int64]string)
	for _, v := range userInfo {
		userIdNameMap[v.Id] = v.Nickname
	}
	fmt.Println("userIdNameMap", userIdNameMap)
	var list []types.HomestayCommonResp
	for _, v := range commonList {
		var ty types.HomestayCommonResp
		_ = copier.Copy(&ty, v)
		ty.StarResult, _ = strconv.ParseInt(v.Star, 10, 32)
		ty.UserName = userIdNameMap[v.UserId]
		ty.CreateTime = v.CreateTime.Format("2006-01-02 15:04:05")
		list = append(list, ty)
	}
	return &types.HomestayCommonListResp{
		List: list,
	}, nil
}
