// 文件: internal/s3/s3client.go
package s3

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"mime/multipart"
)

type Client struct {
	client     *s3.Client
	uploader   *manager.Uploader
	bucket     string
	cloudFront string
}

type Config struct {
	Region     string
	Bucket     string
	Endpoint   string // 可选，用于兼容 MinIO、OSS 等 S3 协议服务
	CloudFront string // 可选，CDN 域名
}

// 初始化 S3 客户端
func NewClient(cfg Config) (*Client, error) {
	awsCfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(cfg.Region),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	s3Client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		if cfg.Endpoint != "" {
			o.BaseEndpoint = aws.String(cfg.Endpoint)
			// 对于 MinIO 等非 AWS 服务，可能需要禁用 SSL 或使用路径模式
			// o.UsePathStyle = true
		}
	})

	return &Client{
		client:     s3Client,
		uploader:   manager.NewUploader(s3Client),
		bucket:     cfg.Bucket,
		cloudFront: cfg.CloudFront,
	}, nil
}

// 上传文件到 S3
func (c *Client) UploadFile(ctx context.Context, file multipart.File, size int64, key string) (string, error) {
	// 对于小文件，可以直接使用 PutObject；Uploader 对大小文件都适用。
	result, err := c.uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket:        aws.String(c.bucket),
		Key:           aws.String(key),
		Body:          file,
		ContentLength: aws.Int64(size),
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %w", err)
	}

	// 如果有 CloudFront 域名，则使用 CDN URL，否则使用 S3 的 URL
	if c.cloudFront != "" {
		return fmt.Sprintf("%s/%s", c.cloudFront, key), nil
	}
	return result.Location, nil
}
