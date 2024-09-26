![Smart](http://localhost:8000/docs/file/logo.png)

# 概述

轻量易用的授权基座smart，提供一个非官方RBAC授权模型的基础开箱即用中台管理服务。

支持响应码自定义国际化、高粒度权限分配、接口路由配置和实时生效，同时支持部门数据权限横向访问限制。

本系统的全部接口采用【POST】【application/json】方式传输数据。

除开放接口以外的其他接口均需要通过【ApiKeyAuth:请求头[Token]】完成鉴权。

系统技术栈：Golang、VueNext、MySQL、Redis、Gin、ArcoDesign

# 接口文档

轻度依赖集成，通过静态HTML加载yaml文件进行打开接口文档。

该接口文档提供Swagger[支持调试]和ReDoc[阅读增强]两个版本。

[Swagger[支持调试]：/docs/swagger/index.html](http://localhost:8000/docs/swagger/index.html) 丨 [ReDoc[阅读增强]：/docs/redoc/index.html](http://localhost:8000/docs/redoc/index.html)

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

# V1.0Fixed