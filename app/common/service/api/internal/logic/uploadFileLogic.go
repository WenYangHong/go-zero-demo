package logic

import (
	"context"
	"fmt"
	"go-zero-mall/pkg/xerr"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"go-zero-mall/app/common/service/api/internal/svc"
	"go-zero-mall/app/common/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadFileLogic {
	return &UploadFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadFileLogic) UploadFile(file multipart.File, header *multipart.FileHeader) (*types.UploadResp, error) {
	// 校验大小
	const maxSize = 10 << 20 // 10 MB
	if header.Size > maxSize {
		return nil, xerr.NewErrMsg("文件超过 10 MB 限制")
	}

	// 安全处理文件名，防路径穿越
	safeFilename := fmt.Sprintf("%d_%s", time.Now().UnixNano(),
		filepath.Base(filepath.Clean(header.Filename)))

	dst, err := os.Create(filepath.Join(l.svcCtx.Config.UploadDir, safeFilename))
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	// size, err := io.Copy(dst, file)
	// if err != nil {
	// 	return nil, err
	// }

	return &types.UploadResp{
		URL: "/files/" + safeFilename,
	}, nil
}

func (l *UploadFileLogic) UploadFileS3(file multipart.File, header *multipart.FileHeader) (*types.UploadResp, error) {
	fmt.Println("uploadToS3")
	return &types.UploadResp{
		URL: "测试代码",
	}, nil
	// // 安全处理文件名，防路径穿越
	// safeFilename := fmt.Sprintf("%d_%s", time.Now().UnixNano(),
	// 	filepath.Base(filepath.Clean(header.Filename)))
	// url, err := l.svcCtx.S3.UploadFile(l.ctx, file, header.Size, safeFilename)
	// if err != nil {
	// 	logx.Errorf("upload file to S3 failed: %v", err)
	// 	return &types.UploadResp{
	// 		URL: "",
	// 	}, errors.New("failed to upload file")
	// }
	// return &types.UploadResp{
	// 	URL: url, // 这里还不是正确的
	// }, nil
}
