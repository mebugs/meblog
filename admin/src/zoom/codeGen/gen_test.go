package codeGen

import (
	"siteol.com/smart/src/zoom/codeGen/gen"
	"testing"
)

/**
 * 基于MySQL表结构构建的代码生成器
 * 1. 您需要具备访问information_schema库权限的账号来完成这些信息的处理
 * 2. 代码生成器自动为您生成基础的Controller、Server、Mapper，以及一些基本的结构体（注意：生成的结构体作为参考，实际使用需要您进行调整）
 * 3. 默认基于注释为您生成对应的Swagger，美观性的调整需要您自行调节（main.go的@x-tagGroups需要手写编写）
 * 4. 由于路由引用的高度自定义性，生成器会为您在routers创建一个临时的路由引用，您可复制到实际的引用位置
 */

// 数据库连接配置
var url = "root:123456@tcp(localhost:3306)/information_schema?charset=utf8mb4&parseTime=True&loc=Local"

// 用户生成的数据库
var dbName = "smart"

/**
 * 包名，该字段将影响为各个包下传入文件的处理
 * 例：src/common/mysql/{X}DB  src/common/model/{X}Model  src/service/{X}
 */
var packName = "plat"

// 计划生成的表名
var tableArray = []string{"login_record"}

// 计划生成的文件控制 DB文件 Model文件 Handle文件/Service文件/Constant示例/Router示例
var flagArray = []bool{true, false, false}

// TestGen
func TestGen(t *testing.T) {
	// 初始化数据库
	err := gen.InitDb(url)
	if err != nil {
		t.Logf("Init Db Fail . Err is %v", err)
		return
	}
	t.Logf("Init Db Success . ")
	// 循环计划处理的表
	for _, tbName := range tableArray {
		// 读取表注释
		tc, err := gen.GetTableBase(packName, dbName, tbName)
		if err != nil {
			t.Logf("%s GetTableBase Err is %v", tbName, err)
			continue
		}
		// 读取表结构
		err = gen.GetTableColumns(tc)
		if err != nil {
			t.Logf("%s GetTableColumns Err is %v", tbName, err)
			continue
		}
		// 生成数据库层代码
		if flagArray[0] {
			err = gen.MakeDbCode(tc, t)
			if err != nil {
				continue
			}
			t.Logf("%s MakeDbCode Success", tbName)
		}
		// 生成Model代码
		if flagArray[1] {
			err = gen.MakeModelCode(tc, t)
			if err != nil {
				continue
			}
			t.Logf("%s MakeModelCode Success", tbName)
		}
		// 生成控制层代码
		// 生成业务层代码
		// 生成路由Demo代码
		// 生成常量Demo代码
		if flagArray[2] {
			err = gen.MakeConstantDemoCode(tc, t)
			if err != nil {
				continue
			}
			t.Logf("%s MakeConstantDemoCode Success", tbName)
			err = gen.MakeServiceWorkCode(tc, t)
			if err != nil {
				continue
			}
			t.Logf("%s MakeServiceWorkCode Success", tbName)
			err = gen.MakeServiceCode(tc, t)
			if err != nil {
				continue
			}
			t.Logf("%s MakeServiceCode Success", tbName)
			err = gen.MakeHandlerCode(tc, t)
			if err != nil {
				continue
			}
			t.Logf("%s MakeHandlerCode Success", tbName)
			err = gen.MakeRouterDemoCode(tc, t)
			if err != nil {
				continue
			}
			t.Logf("%s MakeRouterDemoCode Success", tbName)

		}
	}

}
