package user_service

import (
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"xhgfast/models"
	"xhgfast/utils/jwt"
	"xhgfast/utils/response"
)

type RegiterService struct {
	Name      string `form:"name" json:"name"`
	Telephone string `form:"telephone" json:"telephone" binding:"required,len=11"`
	Password  string `form:"password" json:"password" binding:"required,min=6,max=16"`
}

func (this *RegiterService) Register() (httpCode int, res response.Res) {
	res = response.Res{
		Code: 0,
		Data: "",
		Msg:  "",
	}
	httpCode = http.StatusOK

	user := models.User{}
	if ok := user.IsExist(this.Telephone); ok {
		res.Code = response.ERROR_CREATE_EXIST_TELEPHONE
		httpCode = http.StatusForbidden
		return
	}
	user.Name = this.Name
	user.Telephone = this.Telephone
	// hash password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(this.Password), bcrypt.DefaultCost)
	if err != nil {
		res.Code = response.ERROR
		httpCode = http.StatusInternalServerError
		return
	}
	user.Password = string(hashPassword)
	if err := models.DB.Create(&user).Error; err != nil {
		res.Code = response.ERROR
		httpCode = http.StatusInternalServerError
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
