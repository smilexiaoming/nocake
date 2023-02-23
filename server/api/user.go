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

// @Summary 用户登陆
// @Description 传入code进行鉴权
// @Accept  json
// @Produce  json
// @Param code path string true "code"
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求错误"
// @Failure 500 {object} response.Response "内部错误"
// @Router /login [post]
func (u *AppUser) UserLogin(context *gin.Context) {
	code := context.PostForm("code")
	if code == "" {
		response.Error(constant.ParamInvalid, context)
		return
	}
	userInfo, errMessage := u.Login(code)
	if errMessage != ""{
		response.Error(errMessage, context)
		return
	}
	if userInfo != nil {
		response.Success(constant.LoginSuccess, userInfo, context)
		return
	}
	response.Error(constant.LoginFailed, context)
}
