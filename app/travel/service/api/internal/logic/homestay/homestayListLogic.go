package homestay

import (
	"context"
	"encoding/json"
	"fmt"
	"go-zero-mall/app/travel/service/api/internal/svc"
	"go-zero-mall/app/travel/service/api/internal/types"
	"go-zero-mall/app/travel/service/model"
	"go-zero-mall/pkg/tool"
	"go-zero-mall/pkg/xerr"
	"math"
	"strconv"

	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// homestay room list
func NewHomestayListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayListLogic {
	return &HomestayListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayListLogic) HomestayList(req *types.HomestayListReq) (resp *types.HomestayListResp, err error) {
	// // 民宿状态表
	// whereBuilder := l.svcCtx.HomestayActivityModel.SelectBuilder().Where(squirrel.Eq{
	// 	"row_type":   model.HomestayActivityPreferredType, // 优选民宿
	// 	"row_status": model.HomestayActivityUpStatus,      // 上架状态
	// })
	//
	// homestayActivityList, err := l.svcCtx.HomestayActivityModel.FindPageListByPage(l.ctx, whereBuilder, req.Page, req.PageSize, "create_time desc")
	// 根据人数筛选
	whereBuilder := l.svcCtx.HomestayModel.SelectBuilder().Where(squirrel.Eq{
		"row_state": model.HomestayUpStatus, // 上架状态
		"del_state": model.DelStateNotDone,  // 未删除
	})
	// 入住人数  房间可支持人数 大于 等于预计入住人数
	if req.PeopleNum > 0 {
		whereBuilder = whereBuilder.Where(squirrel.GtOrEq{"people_num": req.PeopleNum})
	}
	// 首页的地点搜索还支持不了 所以就暂时用民宿名称搜索吧
	if req.SearchKey != "" {
		whereBuilder = whereBuilder.Where(squirrel.Like{"title": req.SearchKey + "%"})
	}
	// 热门分类
	if req.HomestayCate != "" {
		// 不做二次验证了
		homestayCategoryWhereBuilder := l.svcCtx.HomestayCategoryModel.SelectBuilder().Where(squirrel.Eq{
			"category_id": req.HomestayCate,
		}).Column("homestay_id")
		homestayId, err := l.svcCtx.HomestayCategoryModel.FindPageListByPage(l.ctx, homestayCategoryWhereBuilder, req.Page, req.PageSize, "create_time desc")

		if err != nil || len(homestayId) == 0 {
			return &types.HomestayListResp{
				List: []types.Homestay{},
				CommonResp: types.CommonResp{
					Code: xerr.OK,
					Msg:  "success",
				},
			}, nil
		}
		// 取出所有的 homestay_id
		var homestayIdList []int64
		for _, v := range homestayId {
			homestayIdList = append(homestayIdList, v.HomestayId)
		}
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"id": homestayIdList,
		})
	}
	var respList []types.Homestay

	homestayList, err := l.svcCtx.HomestayModel.FindPageListByPage(l.ctx, whereBuilder, req.Page, req.PageSize, "create_time desc")
	if err != nil {
		return nil, err
	}
	if len(homestayList) == 0 {
		return &types.HomestayListResp{
			List: []types.Homestay{},
			CommonResp: types.CommonResp{
				Code: xerr.OK,
				Msg:  "success",
			},
		}, nil
	}
	// 查询评价的信息
	var homestayIdList []int64
	for _, v := range homestayList {
		homestayIdList = append(homestayIdList, v.Id)
	}
	commonMap, err := l.getCommonListById(homestayIdList)
	if err != nil {
		return nil, err
	}
	fmt.Println("commonMap", commonMap)
	for _, hs := range homestayList {
		var ty types.Homestay
		_ = copier.Copy(&ty, hs)
		img_info_list := []string{}
		if hs.ImgInfo != "" || string(hs.ImgInfo) != "" {
			err := json.Unmarshal([]byte(hs.ImgInfo), &img_info_list)
			if err != nil {
				img_info_list = []string{}
			}
		}
		ty.Id = strconv.FormatInt(hs.Id, 10)
		ty.HomestayBusinessId = hs.HomestayBusinessId
		ty.UserId = hs.UserId
		ty.FoodPrice = tool.ToFloat(hs.FoodPrice)
		ty.HomestayPrice = tool.ToFloat(hs.HomestayPrice)
		ty.MarketHomestayPrice = tool.ToFloat(hs.MarketHomestayPrice)
		ty.ImgInfo = img_info_list

		// 安全赋值评论数据
		if stat, ok := commonMap[hs.Id]; ok && stat != nil {
			ty.CommentCount = stat.Count                           // 假设 types.Homestay 中有这两个字段
			ty.CommentAvgScore = math.Round(stat.AvgScore*10) / 10 // 如果没有需要先添加
		} else {
			ty.CommentCount = 0
			ty.CommentAvgScore = 0.0
		}

		respList = append(respList, ty)
	}

	return &types.HomestayListResp{
		List: respList,
		CommonResp: types.CommonResp{
			Code: xerr.OK,
			Msg:  "success",
		},
	}, nil
}

// logic 包内定义（或 model 包内定义，但此处按 logic 使用方便）
type HomestayCommentStats struct {
	Count    int64   `db:"count"`     // 评价总数
	AvgScore float64 `db:"avg_score"` // 平均分
}

func (l *HomestayListLogic) getCommonListById(homestayIdList []int64) (map[int64]*HomestayCommentStats, error) {
	result := make(map[int64]*HomestayCommentStats)
	if len(homestayIdList) == 0 {
		return result, nil
	}

	statsList, err := l.svcCtx.HomestayCommentModel.FindStatsByHomestayIds(l.ctx, homestayIdList)
	if err != nil {
		logx.Errorf("FindStatsByHomestayIds error: %v, ids: %v", err, homestayIdList)
		return nil, err
	}

	for _, stat := range statsList {
		result[stat.HomestayId] = &HomestayCommentStats{
			Count:    stat.Count,
			AvgScore: stat.AvgScore,
		}
	}
	return result, nil
}
