package api

import (
	"nocake/constant"
	"nocake/response"
	"nocake/service"

	"github.com/gin-gonic/gin"
)

type AppUser struct {
	service.AppUserService
}

func GetAppUser() *AppUser {
	return &AppUser{}
}

func (u *AppUser) UserLogin(context *gin.Context) {
	code := context.PostForm("code")
	if code == "" {
		response.Error(constant.ParamInvalid, context)
		return
	}
	if userInfo := u.Login(code); userInfo != nil {
		response.Success(constant.LoginSuccess, userInfo, context)
		return
	}
	response.Error(constant.LoginFailed, context)
}
