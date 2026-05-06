package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	UploadDir string
	S3Config  struct {
		Region     string
		Bucket     string
		Endpoint   string
		CloudFront string
	}
	FileUploadType string
}
