package svc

import (
	"fmt"
	s3 "go-zero-mall/app/common/service/api/internal/S3"
	"go-zero-mall/app/common/service/api/internal/config"
)

type ServiceContext struct {
	Config config.Config
	S3     *s3.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	fmt.Println("文件存储方式：", c.FileUploadType)
	if c.FileUploadType == "s3" {
		// 初始化 S3 客户端
		s3Client, err := s3.NewClient(s3.Config{
			Region:     c.S3Config.Region,
			Bucket:     c.S3Config.Bucket,
			Endpoint:   c.S3Config.Endpoint,   // 可选
			CloudFront: c.S3Config.CloudFront, // 可选
		})
		if err != nil {
			panic(fmt.Sprintf("failed to init s3 client: %v", err))
		}

		return &ServiceContext{
			Config: c,
			S3:     s3Client,
		}
	} else {
		return &ServiceContext{
			Config: c,
		}
	}

}
