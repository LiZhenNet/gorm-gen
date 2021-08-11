## Gorm Generator

### Overview
> 根据表结构生成 golang struct 
1. 生成符合 gorm 的 struct
2. 字段添加 gorm&json tag
> 生成对应的 Dal 简化增删改查
1. 生成 Dal 辅助结构体，避免手写gorm字符串参数导致的问题，参数类型安全
2. 生成 common dal  提供增删改查方法

### [Example](./example)
[use case](./example/mian.go)   
[model struct](./example/internal/model/project_model.go)  
[dal struct](./example/internal/dal/project_model_common_dal.go)  

### How to use
#### install 
```
go get github.com/lizhennet/gorm-gen/cmd
```

#### Use
1. 创建 generator 需要的配置文件 [示例](./example/config/gorm-gen.yml)
2. 执行生成 struct  

只生成model
```
gorm-gen model  tableName   
```
生成model&Dal
```
gorm-gen model  tableName --dal 
```
只生成Dal  
```
gorm-gen dal  tableName  
```
查看帮助
```
gorm-gen -h 
```