# ddd-template
六边形领域驱动项目架构模板，完善中

## 项目结构
- docs 文档
- configs 配置文件
- graphql graphql文件
- internal
    - adapters 适配器代码
        - mock mock代码
        - repository 存储器
        - store 缓存
    - application 应用层
      - mutation 写入操作
      - query 查询操作
      - transform 实体转成返回对象
      - schema api返回表单
    - domain 领域层 
    - ports 端口
      - controller 控制器
          - handler 接口处理
          - middleware http中间件
      - graph graphQL入口
    - rely 项目启动依赖配置和系统变量
    - service 服务入口
## 运行
```shell
go run . --config configs/config.yaml
```

## docker 部署
```shell
docker-compose up -d 
```

## docker 更新
```shell
    docker-compose restart app
```

## 单元测试
```shell
sh run_unit_test.sh
```

## 生成mock代码
```shell
make -f ./makefile mockgen
```


## 生成EC256
```shell
openssl ecparam -genkey -name prime256v1 -noout -out access_token_key.pem
openssl ec -in access_token_key.pem -pubout -out access_token_pub.pem
openssl ecparam -genkey -name prime256v1 -noout -out refresh_token_key.pem
openssl ec -in refresh_token_key.pem -pubout -out refresh_token_pub.pem
```

## 感谢JetBrains的开源许可证书支持
<img src="https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.png?_gl=1*l2f4tq*_ga*MTE4NTc2NDE2MC4xNjU0MTM5MzQ0*_ga_9J976DJZ68*MTY1NDEzOTM0NC4xLjAuMTY1NDEzOTM0NC4w" alt="JetBrains" width="200">
