# ddd-template
六边形领域驱动项目架构模板，完善中

## 项目结构
- cmd 
  - server 服务开启入口
- docs 文档
- internal
    - infra 基础设施
        - mock mock代码
    - app 应用层的接口定义
    - common 公用包
    - domain 领域设计
    - service 服务
      - delivery 运输
      - repository 存储库存储
      - usecase 用例
    - initialization 项目初始化

## 运行
```shell
go run ./cmd/server -conf conf ./configs/config.yaml
```
## 单元测试
```shell
sh run_unit_test.sh
```

## 生成mock代码
```shell
make -f ./makefile mockgen
```
## 感谢JetBrains的开源许可证书支持
<img src="https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.png?_gl=1*l2f4tq*_ga*MTE4NTc2NDE2MC4xNjU0MTM5MzQ0*_ga_9J976DJZ68*MTY1NDEzOTM0NC4xLjAuMTY1NDEzOTM0NC4w" alt="JetBrains" width="200">