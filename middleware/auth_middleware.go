package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"xhgfast/models"
	"xhgfast/utils/jwt"
	"xhgfast/utils/response"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取及验证token
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			response.Response(ctx, http.StatusUnauthorized, response.ERROR_EMPTY_AUTH, nil)
			ctx.Abort()
			return
		}
		tokenString = tokenString[7:]
		claims, err := jwt.ParseToken(tokenString)
		if err != nil {
			response.Response(ctx, http.StatusUnauthorized, response.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
			ctx.Abort()
			return
		}
		// 验证用户是否存在
		UserId := claims.UserId
		user := models.User{}
		if err := user.GetUserById(int(UserId)); err != nil {
			response.Response(ctx, http.StatusUnauthorized, response.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
			ctx.Abort()
			return
		}
		ctx.Set("user", user)
		ctx.Next()
	}
}
