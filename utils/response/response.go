package response

import "C"
import "github.com/gin-gonic/gin"

// Response 基础序列化器
type Res struct {
	Code int         `json:"code"`
	Data interface{} `json:"data,omitempty"`
	Msg  string      `json:"msg"`
}

func Response(ctx *gin.Context, httpCode, errCode int, data interface{}) {

	ctx.JSON(httpCode, Res{
		Code: httpCode,
		Data: data,
		Msg:  GetMsg(errCode),
	})

	return
}
