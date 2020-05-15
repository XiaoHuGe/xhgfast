package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xhgfast/dto"
	"xhgfast/models"
	"xhgfast/service/user_service"
	"xhgfast/utils/response"
)

func Register(ctx *gin.Context) {
	serviceReg := user_service.RegiterService{}
	if err := ctx.ShouldBind(&serviceReg); err != nil {
		panic(err)
		return
	}
	httpCode, res := serviceReg.Register()
	response.Response(ctx, httpCode, res.Code, res.Data)
	return
}

func Login(ctx *gin.Context) {
	serviceLogin := user_service.LoginService{}
	if err := ctx.ShouldBind(&serviceLogin); err != nil {
		panic(err)
		return
	}
	httpCode, res := serviceLogin.Login()
	response.Response(ctx, httpCode, res.Code, res.Data)
	return
}

func Info(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		response.Response(ctx, http.StatusInternalServerError, response.ERROR, nil)
		return
	}
	response.Response(ctx, http.StatusOK, response.SUCCESS, dto.ToUserDto(user.(models.User)))
	return
}
