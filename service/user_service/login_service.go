package user_service

import (
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"xhgfast/models"
	"xhgfast/utils/jwt"
	"xhgfast/utils/response"
)

type LoginService struct {
	Telephone string `form:"telephone" json:"telephone" binding:"required,len=11"`
	Password  string `form:"password" json:"password" binding:"required,min=6,max=16"`
}

func (this *LoginService) Login() (httpCode int, res response.Res) {
	res = response.Res{
		Code: 0,
		Data: "",
		Msg:  "",
	}
	httpCode = http.StatusOK

	user := models.User{}
	if ok := user.IsExist(this.Telephone); !ok {
		res.Code = response.ERROR_CREATE_NAME_OR_PASSWORD
		httpCode = http.StatusForbidden
		return
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(this.Password)) != nil {
		res.Code = response.ERROR_CREATE_NAME_OR_PASSWORD
		httpCode = http.StatusForbidden
		return
	}

	// Generate token
	token, err := jwt.GenerateToken(user)
	if token == "" || err != nil {
		res.Code = response.ERROR_AUTH_TOKEN
		httpCode = http.StatusInternalServerError
		return
	}
	res.Code = response.SUCCESS
	res.Data = map[string]string{"token": token}
	return
}
