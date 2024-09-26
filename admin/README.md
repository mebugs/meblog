# smart

![](https://gitee.com/siteol-com/smart/raw/master/docs/logo.png)

## 介绍

轻量易用的授权基座smart，提供一个RBAC授权模型的基础开箱即用中台管理服务。

支持响应码自定义国际化、高粒度权限分配、接口路由配置和实时生效，同时支持部门数据权限横向访问限制。

本系统的全部接口采用【POST】【application/json】方式传输数据。

除开放接口以外的其他接口均需要通过【ApiKeyAuth:请求头[Token]】完成鉴权。

系统技术栈：Golang、VueNext、MySQL、Redis、Gin、ArcoDesign

Demo地址：[http://101.132.76.203/](http://101.132.76.203/)

Demo账密：admin  123456

接口文档提供Swagger[支持调试]和ReDoc[阅读增强]两个版本。

[Swagger[支持调试]：http://101.132.76.203/docs/swagger/index.html](http://101.132.76.203/docs/swagger/index.html)

[ReDoc[阅读增强]：http://101.132.76.203/docs/redoc/index.html](http://101.132.76.203/docs/redoc/index.html)

## 当前说明

数据库初始化脚本位于：config/smart_init.sql

## 启动准备

### 依赖准备

1. golang 1.20
2. 工程目录执行`go mod tidy`更新依赖
3. 配置文件位置`config`目录下，配置注释`config_comment.js`

### 环境变量

应用启动需要添加以下环境变量，IDE（如：goLand）可临时添加。

生产部署，可放在启动脚本中，参考`restart.sh`

如：`ENV=test;NODE=APP01`

- ENV=test/prod
- NODE=APP01/APP0X

# SwaggerAPI文档

轻度依赖集成，通过静态HTML加载yaml文件进行打开接口文档。

## 生成说明

- 首次集成

go install github.com/swaggo/swag/cmd/swag@latest

- 格式化注释代码 (选用)

swag fmt

- 初始化swagger.yaml文件

swag init

- 如果像本工程一样依赖了api.md和独立的api.go文件描述项目（推荐，后续持续使用）

swag init -g api.go --md .

- 删除多余的生成

rm .\docs\docs.go

删除docs.go是因为本项目中并未采用下述依赖来集成。

- github.com/swaggo/swag
- github.com/swaggo/gin-swagger
- github.com/swaggo/files

而是通过HTML模板+JS+YAML引入集成，移除docs.go用于避免依赖编译报错。

源帮助页：[https://github.com/swaggo/swag/blob/master/README_zh-CN.md](https://github.com/swaggo/swag/blob/master/README_zh-CN.md)

# 本地工具

本地工具以测试类形式归纳在`src/zoom`包下。

所有测试类的启动目录都是代码所在位置，测试类提供了详细的注释说明。

## 代码生成器

src/zoom/codeGen/gen_test.go

生成数据库实体类、请求响应实体类、路由入口、响应码常量、控制层、业务层。

## 响应码生成器

src/zoom/respCode/response_code.xlsx

在Excel填写响应内容，其中编号由公式自动计算生成。

src/zoom/respCode/make_test.go

运行测试类自动生成初始化数据库的SQL以及响应码常量文件。

- src/zoom/respCode/response_code.sql
- src/common/constant/response_code.go

## 国际化生成器

src/zoom/i18n/i18n.xlsx

填写国际化，不同模块分不同Sheet，第一个Sheet填写语言，注意后续的国际化Sheet的语言列要完整。

src/zoom/i18n/make_test.go

依照表格生成模块已经对应的前端国际化文件，TypeScript格式。

src/zoom/i18n/i18n.make.exe

编译的Windows执行文件，可以放在前端工程中，位置您可参考本项目的前端项目（大致在：src\locale\i18n.make.exe）。

# 设计思想

## RBAC模型

账号 - 角色（多） - 权限（多） - 接口路由（多）。

权限颗粒度为五层，系统、模块、页面、按钮、路由（路由层已权限关联的形式构成路由集）。

## 数据权限

部门数据权限体系，支持多维度的数据权限配置，本级与下级、仅本级、仅个人、全部、指定部门（TODO）、指定人（TODO）。

账号的数据权限可以自由选择，继承部门、本部门、本人、全局。

## 动态系统配置

系统配置入库，支持热配置生效。

默认的系统配置主要包含安全方面，登陆风控、多端登录限制、会话时长等。

## 响应码国际化

未配置的响应码翻译以默认成功和默认失败响应，支持响应码国际化翻译，支持变量注入。

## 动态路由配置

未配置的路由禁止访问，支持热处理，接口加入/移除、授权配置、日志记录、报文入库、脱敏加密等。

