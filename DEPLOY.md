## 港航科项目部署文档
### 环境依赖以及说明
- Docker 20以上版本
- 如果主机存在GPU显卡，请安装GPU显卡驱动和GPU的docker支持插件，基于显卡型号和类型区别，安装过程自行查看文档。
- 系统环境变量 GHK_DATA_DIR，如果不设置项目启动默认会从当前工作目录读取静态文件信息./ghk_data
- 项目静态总文件夹文件ghk_data，以下目录通过docker挂载到项目或者数据库的文件保存路径

    ```特别说明，ghk_data下的目录名称请勿修改```
    - ccb-bigdata-back/source 对应 ccb-bigdata-back项目的source目录
    - ccb-review-system 对应 ccb-review-system项目的static目录
    - ctxt-micro/static 对应ctxt-mirco的static目录
    - ghk_db 对应项目数据库db的/var/lib/postgresql/data目录
    - ghk_geoserver_data 对应 ccb-review-system项目依赖的geoserver的data_dir
    - information-analysis-micro 对应information-analysis-micro项目的static目录，也对应ml人工智能遥感影像监测的static目录
    - remote_sensing_information_analysis 对应remote_sensing_information_analysis项目下的static目录
    - satellite-data-portrait-back/static 对应satellite-data-portrait-back的static目录
    - spatial-data-micro/static 对应spatial-data-micro的static目录，也对应file-server的static目录
    - system-micro/uploads 对应system-micro的uploads目录
    - thematic-system 对应thematic-system的thematic-source目录
    - user-micro/source 对应user-micro的source目录

### 项目说明
- ccb-bigdata-back 对应遥感应用选项的宏观经济、绿色金融、大宗商品三个菜单
- ccb-review-system 对应遥感应用选项的资产管理菜单
- satellite-data-portrait-back 对应遥感应用选项的农业监测菜单
- information-analysis-micro 对应应用分析选项
- remote_sensing_information_analysis 对应信息分析挖掘选项
- thematic-system 对应专题产品选项
- spatial-data-micro 对应资源共享
- ckxt-micro 测控系统
- ml 人工智能遥感监测服务，属于项目内部依赖，需要GPU环境才能运行。

### 部署

#### 如果本地没有存在代码仓库
```shell
git clone ghk-data-api  git@e.coding.net:zkxrsz/ghk-data-api/ghk-data-api.git
```
#### 设置静态目录环境依赖，如果不设置，默认读取当前工作目录下的./ghk_data目录，例如
```shell
  export GHK_DATA_DIR=$HOME/data/ghk_data
```

#### 项目启动需要修改或者添加配置说明
- information-analysis-micro 项目指定配置文件是仓库下的information-analysis-micro/configs/config.yaml，目前无需修改便可以启动
- remote_sensing_information_analysis 项目指定配置文件是仓库下的remote_sensing_information_analysis/configs/config.yaml，目前无需修改便可以启动
- ccb-review-system 项目指定配置文件是仓库下的ccb-review-system/configs/test-global-config.yaml，目前无需修改便可以启动
- 

#### 进入项目目录，部署
```text
提示：若存在GPU环境，编辑ghk-data-api/docker-compose.yml文件,打开ml服务下面的deploy注释
编辑ghk-data-api/information-analysis-micro/script/ml/config/config.yaml的device为gpu
```

```shell
cd ghk-data-api
docker-compose up -d --build
```
- 导入数据库数据(如果在docker启动之前ghk_data目录下没有ghk_db的数据需要手动导入),sql数据会随项目文档一起打包
```shell
# db 密码 starwiz123
  psql -U postgres -d gis -h 127.0.0.1 -p 54399 -f gis.sql
  psql -U postgres -d belt_and_road -h 127.0.0.1 -p 54399 -f ccb_review.sql
  psql -U postgres -d ccb_bigdata -h 127.0.0.1 -p 54399 -f ccb_bigdata.sql
  psql -U postgres -d satellite_data_portrait -h 127.0.0.1 -p 54399 -f sate.sql
```
- 系统重启或者第一次部署需要注意事项
```shell
# 特别提示：有可能在启动的时候出现网关错误，正常情况下此时执行下面这个指令可以解决
docker-compose restart api-gateway
```
- 重启或者更新服务
  - remote_sensing_information_analysis
    ```shell
    # 重启
    docker-compose restart information_dig
    # 更新 
    docker-compose up -d --build information_dig 
    ``` 
  - information-analysis-micro
    ```shell
    # 重启
    docker-compose restart information_analysis
    # 更新 
    docker-compose up -d --build information_analysis 
    ```
  - api-gateway
    ```shell
    # 重启
    docker-compose restart api-gateway
    # 更新
    docker-compose up -d --build api-gateway
    ```
  - ckxt-micro
    ```shell
    # 重启
    docker-compose restart ckxt-micro
    # 更新
    docker-compose up -d --build ckxt-micro
    ```
  - satellite-data-portrait-back
    ```shell
    # 重启
    docker-compose restart satellite-data-portrait-back
    # 更新
    docker-compose up -d --build satellite-data-portrait-back
    ```
    
  - spatial-data-micro
    ```shell
    # 重启
    docker-compose restart spatial-data-micro
    # 更新
    docker-compose up -d --build spatial-data-micro
    ```
    
  - file-server
    ```shell
    # 重启
    docker-compose restart file-server
    # 更新
    docker-compose up -d --build file-server
    ```
    
  - user-micro
    ```shell
    # 重启
    docker-compose restart user-micro
    # 更新
    docker-compose up -d --build user-micro
    ```
    
  - system-mirco
    ```shell
    # 重启
    docker-compose restart system-micro
    # 更新
    docker-compose up -d --build system-micro
    ```
    
  - thematic-system
    ```shell
    # 重启
    docker-compose restart thematic-system
    # 更新
    docker-compose up -d --build thematic-system
    ```
    
  - ccb-review-system
    ```shell
    # 重启
    docker-compose restart ccb-review-system
    # 更新
    docker-compose up -d --build ccb-review-system
    ```
    
  - ccb-bigdata
    ```shell
    # 重启
    docker-compose restart ccb-bigdata
    # 更新
    docker-compose up -d --build ccb-bigdata
    ```
    
  - satellite-track
    ```shell
    # 重启
    docker-compose restart satellite-track
    # 更新
    docker-compose up -d --build satellite-track
    ```