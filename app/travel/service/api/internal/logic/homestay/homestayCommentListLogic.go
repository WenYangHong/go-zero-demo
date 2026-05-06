package homestay

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"

	"go-zero-mall/app/travel/service/api/internal/svc"
	"go-zero-mall/app/travel/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayCommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// homestay room common list
func NewHomestayCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayCommentListLogic {
	return &HomestayCommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayCommentListLogic) HomestayCommentList(req *types.HomestayCommentListReq) (resp *types.HomestayCommentListResp, err error) {
	whereBuilder := l.svcCtx.HomestayCommentModel.SelectBuilder().Where(squirrel.Eq{
		"homestay_id": req.Id,
		"del_state":   0,
	})
	commonList, err := l.svcCtx.HomestayCommentModel.FindPageListByPage(l.ctx, whereBuilder, req.Page, req.PageSize, "create_time desc")
	if len(commonList) == 0 {
		return &types.HomestayCommentListResp{
			List: []types.HomestayCommentResp{},
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
	var list []types.HomestayCommentResp
	for _, v := range commonList {
		var ty types.HomestayCommentResp
		_ = copier.Copy(&ty, v)
		ty.AvgStarResult = v.AvgStar
		ty.CleanlinessStarResult = v.CleanlinessStar
		ty.LocationStarResult = v.LocationStar
		ty.ServiceStarResult = v.ServiceStar
		ty.ValueStarResult = v.ValueStar
		ty.UserName = userIdNameMap[v.UserId]
		ty.CreateTime = v.CreateTime.Format("2006-01-02 15:04:05")
		list = append(list, ty)
	}
	return &types.HomestayCommentListResp{
		List: list,
	}, nil
}
