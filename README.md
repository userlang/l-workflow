# workflow

#### 介绍
目前市场上的工作流都是超级臃肿，Activiti，Camunda，Flowable 动不动就30来张表， 看起来就像是在 秀 什么
对于很多业务系统里，根本不需要额外的功能，只提供工作流即可，l-workflow做到了，而且仅仅只用了，5张表 就可以实现根据用户配置工作流并使用，
配置工作流程描述 并配置好上一步下一步即可调用API调用，简单又灵活，简直YYDS

#### 软件架构
golang
gin
gorm
mysql
redis
nacos
gopkg

#### 安装教程

1. 本项目使用[Go Mod](https://github.com/golang/go/wiki/Modules)管理依赖。
2.  省略 补充
3.  项目运行后启动在 808o 端口上

#### 目录结构

1. cache 内部缓存
2. common 工具
3. config 配置
4. controllers 接口层
5. log 日志
6. logic 逻辑层
7. models 实体层
8. nacos 注册中心配置中心
9. redis 外部缓存
10. repository 数据操作层
11. route 路由层

#### 作者/参与贡献者
1.  郎俊楠


 
