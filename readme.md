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


-- 跳坑记录
```log
[ERROR] FindAll users failed: sql: Scan error on column index 8, name "delete_time": unsupported Scan, storing driver.Value type <nil> into type *time.Time, userIdList: [4]

**** users表中delete_time字段为null，Go struct 用的是 *time.Time，驱动不会自动把 NULL 转成这个类型，导致扫描失败
```