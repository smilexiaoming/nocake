package api

import (
	"fmt"
	"nocake/common"
	"nocake/constant"
	"nocake/models/web"
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

type WebUser struct {
	service.UserService
}

func GetWebUser() *WebUser {
	return &WebUser{}
}

// 获取验证码
func (u *WebUser) GetCaptcha(c *gin.Context) {
	id, base64s, _ := common.GenerateCaptcha()
	data := map[string]interface{}{"captchaId": id, "captchaImg": base64s}
	response.Success("操作成功", data, c)
}

// 用户登录
func (u *WebUser) UserLogin(c *gin.Context) {
	var param web.LoginParam
	if err := c.ShouldBind(&param); err != nil {
		response.Error(constant.ParamInvalid, c)
		return
	}

	// 检查验证码
	fmt.Printf("param: %v\n", param)
	if !common.VerifyCaptcha(param.CaptchaId, param.CaptchaValue) {
		response.Error("验证码错误", c)
		return
	}
	// 生成token
	uid := u.Login(param)
	if uid > 0 {
		token, _ := common.GenerateToke(param.Username)
		userInfo := web.UserInfo{
			Sid:   uid,
			Token: token,
		}
		response.Success("登录成功", userInfo, c)
		return
	}
	response.Error("用户名或密码错误", c)
}

// @Summary 用户登陆
// @Description 传入code进行鉴权
// @Accept  multipart/form-data
// @Produce  json
// @Param code formData string true "code"
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求错误"
// @Failure 500 {object} response.Response "内部错误"
// @Router /app/login [post]
// @RequestMapping(value = "/user")
func (u *AppUser) UserLogin(context *gin.Context) {
	code := context.PostForm("code")
	if code == "" {
		response.Error(constant.ParamInvalid, context)
		return
	}
	userInfo, errMessage := u.Login(code)
	if errMessage != "" {
		response.Error(errMessage, context)
		return
	}
	if userInfo != nil {
		response.Success(constant.LoginSuccess, userInfo, context)
		return
	}
	response.Error(constant.LoginFailed, context)
}
