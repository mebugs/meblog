package worker

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"siteol.com/smart/src/common/constant"
	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/common/model/baseModel"
	"siteol.com/smart/src/common/model/cacheModel"
	"siteol.com/smart/src/common/mysql/platDB"
	"siteol.com/smart/src/common/utils/security"
	"strings"
	"time"
)

// ReturnMsgTrans 执行响应码 => 响应文言 翻译
func ReturnMsgTrans(res *baseModel.ResBody, c *gin.Context, router *cacheModel.CacheRouter, traceID string) {
	// 语言类型
	lang := c.GetString(constant.ContextLang)
	// 非400校验错误执行翻译
	if res.HttpCode != http.StatusBadRequest {
		// 执行翻译
		tableMsgTrans(res, lang, traceID)
	}
	// 响应日志
	printStr := ""
	// 响应安全日志
	resBts, _ := json.Marshal(res)
	printSafeStr := security.SafeJson(string(resBts), router.ResSecure)
	if router.ResPrint {
		// 如需打印日志
		printStr = fmt.Sprintf("Res Code Is : %d . Res Body Is : %s", res.HttpCode, printSafeStr)
	} else {
		// 仅打印响应消息
		printStr = fmt.Sprintf("Res Code Is : %d . Res Body Is : Do Not Print . %s", res.HttpCode, res.Code)
	}
	log.InfoTF(traceID, printStr)
	// 日志入库
	if router.LogInDb {
		// 提取日志对象
		obj, ok := c.Get(constant.ContextRouterI)
		if ok {
			// 提取到则异步加入到库中
			now := time.Now()
			logInDb := &platDB.RouterLog{}
			logInDb = obj.(*platDB.RouterLog)
			if logInDb != nil {
				logInDb.ResStatus = res.HttpCode
				logInDb.ResBody = printSafeStr
				logInDb.ResTime = &now
				go func() {
					err := platDB.RouterLogTable.InsertOne(logInDb)
					if err != nil {
						log.ErrorTF(traceID, "Insert Router Log Fail . Err Is : %v", err)
					}
				}()
			}
		}
	}
}

// tableMsgTrans 执行Msg翻译
func tableMsgTrans(res *baseModel.ResBody, lang, traceID string) {
	// 出错不翻译
	transMap, err := cacheModel.GetResponseCache(traceID)
	if err != nil {
		return
	}
	// 读取配置
	codeMap, ok := transMap[res.Code]
	// 未能匹配的响应码，复用基础文言
	if !ok {
		useCode := constant.Success
		if strings.HasPrefix(res.Code, "F") {
			useCode = constant.Fail
		}
		if strings.HasPrefix(res.Code, "E") {
			useCode = constant.Error
		}
		codeMap, ok = transMap[useCode]
	}
	// 执行语言读取
	if ok {
		langTemple, lok := codeMap[lang]
		if lok {
			// 检查是否有变量，有则进行变量替换
			if strings.Index(langTemple, "}}") > -1 {
				res.Msg = TableValReplace(langTemple, res.Data)
			} else {
				res.Msg = langTemple
			}
		}
	}
	// 无相关翻译
	return
}

// TableValReplace 执行变量替换
// 实际生效，{{name}}修改成功 => 米虫修改成功
func TableValReplace(temple string, data any) string {
	if data == nil {
		return temple
	}
	dataStr, err := json.Marshal(data)
	if err != nil {
		return temple
	}
	dataMap := make(map[string]any)
	err = json.Unmarshal([]byte(dataStr), &dataMap)
	if err != nil {
		return temple
	}
	// 提取模板中的变量数据
	valArray := getTempleVal(temple)
	// 提取参数并替换
	for _, val := range valArray {
		valObj, ok := dataMap[val]
		if ok && valObj != nil {
			// 变量存在值则替换
			temple = strings.ReplaceAll(temple, "{{"+val+"}}", fmt.Sprintf("%v", valObj))
		}
	}
	return temple
}

// 提取模板中的变量数据
func getTempleVal(temple string) []string {
	valArray := make([]string, 0)
	strS := strings.Split(temple, "{{")
	for _, i := range strS {
		if strings.Index(i, "}}") > -1 {
			valStr := i[:strings.Index(i, "}}")]
			valArray = append(valArray, valStr)
		}
	}
	return valArray
}
