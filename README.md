# Orginone 

> Org-in-one 英文发音[ˈɔːdʒɪnʌn]，中文:奥集能

## 什么是 Orginone
  
 
 

上架应用应以组织为颗粒度设计多租户，用户从属于组织，组织有内设机构和岗位，用户可以加入多组织，组织之间可以组成集团。

传统SaaS应用提供服务，组织或用户需要登陆服务地址，注册账号，提供相关用户认证信息。Orginone要求面向组织和用户聚合服务。SaaS应用必须注册到平台，组织用户登陆平台，即可获取、分配和使用SaaS服务。

## Orginone 的理念
  
### 精一

**精一**语出《尚书·大禹谟》：“人心惟危，道心惟微，惟精惟一，允执厥中” 。

### 核心关键词

**精益、专注、中立**

Orginone以组织用户为中心，利用云原生技术提升应用和数据的管理能力，让组织更专注核心业务的代码和数据，以All-in-one方式降低组织中用户对工作的心智负担。

保持平台开放和中立，通过更细颗粒度的划分服务边界，明晰权责利关系，激发市场创新活力。聚焦能力集成，快速迭代，精益求精，为组织用户业务数字化改革和价值创造提供持续的推动力。

## Orginone 的目标

> 成为面向组织用户的云原生操作系统。

为组织提供一体化操作界面，集中管理组织架构和用户体系，满足组织用户统一管理软件、业务数据和云资源等核心数字资产的需求，面向用户聚合各类云服务。利用云原生技术架构原则和设计模式，发挥平台优势，解开组织、应用、数据和资源的耦合关系，以应用市场上架和获取云服务方式，转变传统软件交付和分发模式，聚焦组织用户的核心价值，培育云服务市场生态，改进组织间大规模协作机制，优化资源配置，推动数字化改革，提高社会整体运转效率。


### 专注能力集成

Orginone 聚焦组织核心业务，围绕用户做聚合服务。

基于云原生技术架构原则和设计模式，充分利用以Kubernetes为核心的开源生态优势，通过OAM等开放标准，向下封装基础设施资源，屏蔽底层架构的差异性，向上支撑多种工作负载和分布式架构。提供应用全生命周期管理、运维管理、配置范围和扩展和管理、以及语言无关的编程框架，一起构成崭新的应用与云之间的操作界面。

将云应用中的非业务代码部分进行最大化的剥离，从而让云设施接管应用中原有的大量非功能特性（如弹性、韧性、安全、可观测性、灰度等），使业务不再有非功能性业务中断困扰的同时，具备轻量、敏捷、高度自动化的优势。满足大数据、人工智能、物联网、区块链等各种类型场景应用需求。

Orginone主要聚焦以下四个方面能力的集成：

1. 应用管理能力    

提供应用开发、部署、迭代、运维、运营能力。利用各类公共组件，敏捷开发环境，快速搭建应用。提供从代码到镜像到部署的流水线支持，通过持续集成和持续部署，进行快速迭代。提供的组件、接口、服务和应用的各种颗粒度运维和运营能力。利用市场机制，提供应用上架、发布、下架、工单等运营工具满足基本商业化运营需求。

2. 组织管理能力
   
提供灵活的组织管理。根据业务场景自由创建组织树。能够实现组织树继承、复制、迁移、拆分、合并等操作，满足各类跨地区、跨层级、跨部门的针对组织树的应用分发、流程设计。能够对接各类基于组织树的用户验证和鉴权，消息分发，数据操作等复杂组织管理需求。

3. 数据管理能力   

提供数据分析、建模能力,专注业务层数据服务能力。利用消息机制，实现元数据、字典等数据标准的维护、发布等管理，在应用设计和建模阶段介入对数据的管理，实现数据规范的前对齐。拥有完善的数据治理能力，能够平滑对接各类开放域或私有域的基础数据平台，公共数据服务平台。实现对数据资产细颗粒度的保护，输出各类场景数据服务，支持应用迭代和创新需求。

4. 资源管理能力
   
利用云原生技术，屏蔽基础资源的差异，实现平台和底层资源解藕，对接虚拟计算资源、容器平台、云存储资源、区块链、物联网等各类基础云服务平台，实现跨云资源适配和管理能力，统一向应用层输出基础能力。


## 如何参与项目

### 概述
Orginone采用开放、开源共建模式，避免重复造轮子，以持续迭代，不断演进的模式，完善公共平台的建设。引入开放社区治理模式，保障平台的开放和中心，建设成果以开放或开源模式输出，鼓励在公共平台基础上开展商业服务，以市场化竞争方式提高资源效率，降低社会运行成本。

Orginone对kubernetes能力封装基于[kubevela](https://github.com/oam-dev/kubevela)，应用层开发基于[go-zero](https://github.com/tal-tech/go-zero)。

正式进入实际开发之前，需要做一些准备工作，比如：熟悉kubernetes和kubevela，Go环境的安装，grpc代码生成使用的工具安装，必备工具Goctl的安装，Golang环境配置等,熟悉go-zero中的编码规范

### 议事规则

设立三级组织

- PMC 

对于有矛盾的议题进行最终仲裁。

- 维护者
 
+ 决定项目方向、功能、发布新版本; 
+ 拥有pr合并及拒绝的权利。
+ 维护者会议就分歧议题采用半数通过投票。

- 贡献者

+ 提交pr，批准pr的权利；
+ 新功能提案的权利。
 
此外:维护者会根据贡献者和开发者的机构来源，周期性进行公示


### 奥集能官网

https://www.orginone.cn

### 开发流程

> goctl环境准备[1]    
> 数据库设计 参考assetcloud平台项目数据库设计 (/geneOS/service/geneOS.sql)   
> 业务开发   
> 创建服务类型（api/rpc/rmq/job/script）   
> 编写api、proto文件 参考现有assetx后台api文档   
> 代码生成   
> 生成数据库访问层代码model 以现有数据表生成对应文件   
> 配置config，yaml变更   
> 资源依赖填（ServiceContext）   
> 添加中间件   
> 业务代码填充   
> 错误处理   


### 工程介绍

#### 一、服务介绍
<!--
    ms为服务主要逻辑需求
    apigw 路由入口此处配有负载均衡,降级,熔断,鉴权等
    user-ms 用户相关服务
    user-apigw 用户相关服务网关
    company-ms 单位相关
    company-apigw 用户相关服务网关
    common 封装工具类

#### 二、分支说明

    master 主分支具有权限设置,想往此分支合并修改请联系管理,tag对应prod环境所运行的代码
    staging 测试环境分支,对应云服务器测试环境,建议不要直接在次分支编写代码而是merge develop分支,合并之前确定develop能run 
    develop 代码编写分支, 次分支解决冲突日常代码编写提交,本地自测相对稳定想发布测试请merge 到staging 分支



#### 示例代码 

利用goctl工具，根据assetcloud平台user表生成crudAPI服务。

#### 示例代码启动服务   

etc/user-api.yaml文件中配置好数据库，在user目录下面执行
```
go run cmd/api/user.go -f cmd/api/etc/user-api.yaml
```
-->
