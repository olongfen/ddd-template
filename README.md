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
    - infra 基础设施
        - mock mock代码
    - app 
        - controller
           - restful http接口实现
           - rpcx rpc接口实现
      - repository 数据存储接口实现
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
## 感谢JetBrains的开源许可证书支持
<img src="https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.png?_gl=1*l2f4tq*_ga*MTE4NTc2NDE2MC4xNjU0MTM5MzQ0*_ga_9J976DJZ68*MTY1NDEzOTM0NC4xLjAuMTY1NDEzOTM0NC4w" alt="JetBrains" width="200">