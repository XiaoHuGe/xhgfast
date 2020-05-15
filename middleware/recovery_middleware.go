package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"xhgfast/utils/response"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx.JSON(http.StatusBadRequest, response.Res{Code: response.ERROR, Data: "", Msg: fmt.Sprint(err)})
			}
		}()
		ctx.Next()
	}
}
