<template>
  <div class="container">
    <div class="vd">
      <div class="vditor-reset" id="preview">
        <h1>smart</h1>
        <p><img src="http://101.132.76.203/static/img/logo.png" alt="" /></p>
        <h2>介绍</h2>
        <p>轻量易用的授权基座smart，提供一个RBAC授权模型的基础开箱即用中台管理服务。</p>
        <p>支持响应码自定义国际化、高粒度权限分配、接口路由配置和实时生效，同时支持部门数据权限横向访问限制。</p>
        <p>本系统的全部接口采用【POST】【application/json】方式传输数据。</p>
        <p>除开放接口以外的其他接口均需要通过【ApiKeyAuth:请求头[Token]】完成鉴权。</p>
        <p>系统技术栈：Golang、VueNext、MySQL、Redis、Gin、ArcoDesign</p>
        <p>Demo地址：<a href="http://101.132.76.203/">http://101.132.76.203/</a></p>
        <p>Demo账密：admin 123456</p>
        <p>接口文档提供Swagger[支持调试]和ReDoc[阅读增强]两个版本。</p>
        <p>
          <a href="http://101.132.76.203/docs/swagger/index.html">Swagger[支持调试]：http://101.132.76.203/docs/swagger/index.html</a>
        </p>
        <p><a href="http://101.132.76.203/docs/redoc/index.html">ReDoc[阅读增强]：http://101.132.76.203/docs/redoc/index.html</a></p>
        <h2>当前说明</h2>
        <p>数据库初始化脚本位于：init目录下</p>
        <h2>启动准备</h2>
        <h3>依赖准备</h3>
        <ol>
          <li>golang 1.20</li>
          <li>工程目录执行<code>go mod tidy</code>更新依赖</li>
          <li>配置文件位置<code>config</code>目录下，配置注释<code>config_comment.js</code></li>
        </ol>
        <h3>环境变量</h3>
        <p>应用启动需要添加以下环境变量，IDE（如：goLand）可临时添加。</p>
        <p>生产部署，可放在启动脚本中，参考<code>restart.sh</code></p>
        <p>如：<code>ENV=test;NODE=APP01</code></p>
        <ul>
          <li>ENV=test/prod</li>
          <li>NODE=APP01/APP0X</li>
        </ul>
        <h1>SwaggerAPI文档</h1>
        <p>轻度依赖集成，通过静态HTML加载yaml文件进行打开接口文档。</p>
        <h2>生成说明</h2>
        <ul>
          <li>首次集成</li>
        </ul>
        <p>go install github.com/swaggo/swag/cmd/swag@latest</p>
        <ul>
          <li>格式化注释代码 (选用)</li>
        </ul>
        <p>swag fmt</p>
        <ul>
          <li>初始化swagger.yaml文件</li>
        </ul>
        <p>swag init</p>
        <ul>
          <li>如果像本工程一样依赖了api.md和独立的api.go文件描述项目（推荐，后续持续使用）</li>
        </ul>
        <p>swag init -g api.go --md .</p>
        <ul>
          <li>删除多余的生成</li>
        </ul>
        <p>rm .\docs\docs.go</p>
        <p>删除docs.go是因为本项目中并未采用下述依赖来集成。</p>
        <ul>
          <li>github.com/swaggo/swag</li>
          <li>github.com/swaggo/gin-swagger</li>
          <li>github.com/swaggo/files</li>
        </ul>
        <p>而是通过HTML模板+JS+YAML引入集成，移除docs.go用于避免依赖编译报错。</p>
        <p>
          源帮助页：<a href="https://github.com/swaggo/swag/blob/master/README_zh-CN.md"
            >https://github.com/swaggo/swag/blob/master/README_zh-CN.md</a
          >
        </p>
        <h1>本地工具</h1>
        <p>本地工具以测试类形式归纳在<code>src/zoom</code>包下。</p>
        <p>所有测试类的启动目录都是代码所在位置，测试类提供了详细的注释说明。</p>
        <h2>代码生成器</h2>
        <p>src/zoom/codeGen/gen_test.go</p>
        <p>生成数据库实体类、请求响应实体类、路由入口、响应码常量、控制层、业务层。</p>
        <h2>响应码生成器</h2>
        <p>src/zoom/respCode/response_code.xlsx</p>
        <p>在Excel填写响应内容，其中编号由公式自动计算生成。</p>
        <p>src/zoom/respCode/make_test.go</p>
        <p>运行测试类自动生成初始化数据库的SQL以及响应码常量文件。</p>
        <ul>
          <li>src/zoom/respCode/response_code.sql</li>
          <li>src/common/constant/response_code.go</li>
        </ul>
        <h2>国际化生成器</h2>
        <p>src/zoom/i18n/i18n.xlsx</p>
        <p>填写国际化，不同模块分不同Sheet，第一个Sheet填写语言，注意后续的国际化Sheet的语言列要完整。</p>
        <p>src/zoom/i18n/make_test.go</p>
        <p>依照表格生成模块已经对应的前端国际化文件，TypeScript格式。</p>
        <p>src/zoom/i18n/i18n.make.exe</p>
        <p>编译的Windows执行文件，可以放在前端工程中，位置您可参考本项目的前端项目（大致在：src\locale\i18n.make.exe）。</p>
        <h1>设计思想</h1>
        <h2>RBAC模型</h2>
        <p>账号 - 角色（多） - 权限（多） - 接口路由（多）。</p>
        <p>权限颗粒度为五层，系统、模块、页面、按钮、路由（路由层已权限关联的形式构成路由集）。</p>
        <h2>数据权限</h2>
        <p>部门数据权限体系，支持多维度的数据权限配置，本级与下级、仅本级、仅个人、全部、指定部门（TODO）、指定人（TODO）。</p>
        <p>账号的数据权限可以自由选择，继承部门、本部门、本人、全局。</p>
        <h2>动态系统配置</h2>
        <p>系统配置入库，支持热配置生效。</p>
        <p>默认的系统配置主要包含安全方面，登陆风控、多端登录限制、会话时长等。</p>
        <h2>响应码国际化</h2>
        <p>未配置的响应码翻译以默认成功和默认失败响应，支持响应码国际化翻译，支持变量注入。</p>
        <h2>动态路由配置</h2>
        <p>未配置的路由禁止访问，支持热处理，接口加入/移除、授权配置、日志记录、报文入库、脱敏加密等。</p>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup></script>
<script lang="ts">
export default {
  name: 'CenterIndex'
}
</script>

<style scoped lang="less">
.vd {
  background: var(--color-bg-1);
  border-radius: 4px;
  width: 100%;
  height: 100%;
  position: relative;
  transition: all 0.3s ease-in-out;
  padding: 15px;
}
.vditor-reset {
  height: 100%;
  overflow: scroll;
  padding: 10px 20px;
  font-size: 16px;
  color: #555;
  h1 {
    font-size: 36px;
  }
  h2 {
    font-size: 32px;
  }
  h3 {
    font-size: 28px;
  }
  h4 {
    font-size: 24px;
  }
  p {
    line-height: 20px;
  }
  ol,
  ul {
    line-height: 30px;
    padding-left: 20px;
  }
  code {
    background-color: rgba(27, 31, 35, 0.05);
    padding: 5px;
    margin: 0 5px;
  }
}
</style>
