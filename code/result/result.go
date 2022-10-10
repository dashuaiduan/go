package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	LAN_ZH_CN = "zh_CN" //和zhCN互相兼容
	LAN_ZHCN  = "zhCN"  //和zh_CN互相兼容
	LAN_ZH_TW = "zh_TW" //和zhTW互相兼容
	LAN_ZHTW  = "zhTW"  //和zh_TW互相兼容
	LAN_EN    = "en"
	LAN_DE    = "de"
	LAN_IT    = "it"
	LAN_ES    = "es"
	LAN_FR    = "fr"
)

// 定义响应结果
type Result struct {
	RequestId interface{} `json:"request_id"` //服务请求request_id
	Code      int         `json:"code"`       // 错误码
	Msg       string      `json:"msg"`        // 错误信息  支持中英文切换
	ErrMsg    string      `json:"err_msg"`    // 保存内部错误信息
	Data      interface{} `json:"data"`       //返回数据
}

func New(c *gin.Context, err Errno, data interface{}) {
	requestId := c.MustGet("request_id")
	result := Result{
		RequestId: requestId,
		Code:      err.Code,
		Msg:       err.Message,
		Data:      data}
	lan := c.GetHeader("lan")
	if lan != LAN_EN { // 如果是英文  替换英文错误信息
		result.Msg = err.EnMessage
	}
	c.AbortWithStatusJSON(http.StatusOK, result)
}
