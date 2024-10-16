package gen

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

const (
	routerCodeTemp = `
package routers

import (
	"github.com/gin-gonic/gin"
	"siteol.com/smart/src/router/middleware"
	"siteol.com/smart/src/service/${dbPack}/${dbPack}Handler"
)

/*
 * 这是一段引用路由的示例，您可拷贝到实际路由代码区后删除本文件
 */

// ${dbPack}DemoRouter 业务路由
func ${dbPack}DemoRouter(router *gin.Engine) {
	${dbPack}Router := router.Group("/${dbPack}Demo", middleware.CommMiddleWare) // 授权中间件
	{
		// ${tableComment}相关
		${tableRouter}Router := ${dbPack}Router.Group("/${tableRouter}")
		{
			${tableRouter}Router.POST("/add", ${dbPack}Handler.Add${tableStruct})
			${tableRouter}Router.POST("/page", ${dbPack}Handler.Page${tableStruct})
			${tableRouter}Router.POST("/get", ${dbPack}Handler.Get${tableStruct})
			${tableRouter}Router.POST("/edit", ${dbPack}Handler.Edit${tableStruct})
			${tableRouter}Router.POST("/del", ${dbPack}Handler.Del${tableStruct})
		}
	}
}

`
)

// MakeRouterDemoCode 生成路由引用类
func MakeRouterDemoCode(tc *TableConfig, t *testing.T) error {
	// 生成代码文件
	code := strings.ReplaceAll(routerCodeTemp, "${tableRouter}", tc.Router)
	code = strings.ReplaceAll(code, "${dbPack}", tc.PackName)
	code = strings.ReplaceAll(code, "${tableStruct}", tc.ObjName)
	code = strings.ReplaceAll(code, "${tableComment}", tc.Remark)
	// 没有目录建目录 src/router/routers
	dir := fmt.Sprintf("../../router/routers")
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		errM := os.MkdirAll(dir, 0777)
		if errM != nil {
			t.Logf("%s MakeRouterDemoCode Mkdir Err is %v", tc.TbName, errM)
			return errM
		}
	}
	// 创建文件
	file := fmt.Sprintf("%s/demo_router_%s.go", dir, tc.TbName)
	err := os.WriteFile(file, []byte(code), 0777)
	if err != nil {
		t.Logf("%s MakeRouterDemoCode WriteFile Err is %v", tc.TbName, err)
		return err
	}
	// 执行go fmt
	cmd := exec.Command("go", "fmt", file)
	_, err = cmd.CombinedOutput()
	if err != nil {
		t.Logf("%s MakeRouterDemoCode FMT Err is %v", tc.TbName, err)
		return err
	}
	return nil
}
