-- 参考项目: https://github.com/Mikaelemmmm/go-zero-looklook

-- 参考文档地址：https://github.com/Mikaelemmmm/go-zero-looklook/tree/main/doc/chinese

-- api接口更新
```goctl
goctl api go -api *.api -dir ../  -style=goZero
```

-- 创建model代码
```goctl
-- 在 deploy目录下执行 [使用本地模版声场model代码]
goctl template init --home ./goctl/1.8.5
goctl model mysql ddl -src="./sql/homestay.sql" -dir="../app/travel/service/model" --home="./goctl/1.8.5" -c
```