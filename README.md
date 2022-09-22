# ddd-template
六边形领域驱动项目架构模板，完善中

## 项目结构
- docs 文档
- configs 配置文件
- internal
    - adapters 适配器代码
    - application 应用层
      - mutation 写入操作
      - query 查询操作
    - config 配置
    - ports 端口
      - middleware http中间件 
    - schema api返回表单
- pkg 可提取公用包
## 运行
```shell
go run ./...
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
