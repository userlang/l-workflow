# workflow

#### 介绍

1、初衷: 目前很多业务系统都需要工作流框架，但是经过我的市场调研之后发现市面上的工作流框架及其臃肿，没有一款好用的，例如 Activiti,Camunda,Flowable这些框架动不动就需要部署 30-40来张表，里面复杂且很多东西没用，我只想要工作流相关的东西，他们却自带绘制及用户权限相关的东西，这种确实不是我想要的效果，
2、目标: 研发l-workflow框架可以做到 体积小又灵活且可配置，而且对外提供接口，
3、成果: l-workflow仅仅只有5张表，可以做到根据动态配置的工作流信息，审批人信息及步骤信息例如上一步下一步，通过接口调用方式 动态处理审批通过，拒绝，开始，重新发起，当前节点及数据，当前任务，历史数据查询等功能，非常好用

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
12. db sql脚本

#### 作者/参与贡献者
1.  郎俊楠


 
