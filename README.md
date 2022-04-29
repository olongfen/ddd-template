# ddd-template
六边形领域驱动项目架构模板，完善中

## 项目结构
- api  项目接口protoc定义
- cmd 
  - server 服务开启入口
  - grpc_client rpc客户端
- configs 配置文件
- deployment 部署脚本
- docs 文档
- internal
    - adapters
        - mock mock代码
        - repository 数据存储接口实现
        - restful http接口实现
        - rpcx rpc接口实现
    - app 
        - service http&rpc服务接口实现
        - usecase 业务逻辑接口实现
    - common 公用包
    - domain 领域设计
- third_party protoc第三方依赖

## 运行
```shell
go run ./cmd/server -conf conf ./configs/config.yaml
```
## 单元测试
```shell
sh run_unit_test.sh
```
## 生成protoc代码
```shell
make -f ./makefile proto_build  
```

## 生成mock代码
```shell
make -f ./makefile mockgen
```
