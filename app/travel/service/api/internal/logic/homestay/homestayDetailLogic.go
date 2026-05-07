package homestay

import (
	"context"
	"encoding/json"
	"strconv"

	"go-zero-mall/app/travel/service/api/internal/svc"
	"go-zero-mall/app/travel/service/api/internal/types"
	"go-zero-mall/pkg/xerr"

	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomestayDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayDetailLogic {
	return &HomestayDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayDetailLogic) HomestayDetail(req *types.HomestayDetailReq) (resp *types.HomestayDetailResp, err error) {
	id, err := strconv.ParseInt(req.Id, 10, 64)
	if err != nil {
		return nil, xerr.NewErrCode(xerr.REUQEST_PARAM_ERROR)
	}

	homestayDetail, err := l.svcCtx.HomestayModel.FindOne(l.ctx, id)
	if err != nil {
		return nil, err
	}

	var ty types.Homestay
	_ = copier.Copy(&ty, homestayDetail)
	ty.Id = strconv.FormatInt(homestayDetail.Id, 10)
	ty.HomestayBusinessId = homestayDetail.HomestayBusinessId
	ty.UserId = homestayDetail.UserId
	// 民宿图片
	img_info_list := []string{}
	if homestayDetail.ImgInfo != "" || string(homestayDetail.ImgInfo) != "" {
		err := json.Unmarshal([]byte(homestayDetail.ImgInfo), &img_info_list)
		if err != nil {
			img_info_list = []string{}
		}
	}
	ty.ImgInfo = img_info_list
	return &types.HomestayDetailResp{
		Homestay: ty,
		CommonResp: types.CommonResp{
			Code: xerr.OK,
			Msg:  "success",
		},
	}, nil
}
