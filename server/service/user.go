package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"nocake/global"
	"nocake/models/app"
	"nocake/models/web"
	"time"
)

type UserService struct {
}

type AppUserService struct {
}

// 商家用户登录
func (u *UserService) Login(param web.LoginParam) uint64 {
	var user web.User
	global.Db.Table("t_web_user").Where("username = ? and password = ?", param.Username, param.Password).First(&user)
	return uint64(user.Id)
}

// 小程序登录
func (u *AppUserService) Login(code string) (*app.UserInfo, string) {
	var acsJson app.Code2SessionResult
	acs := app.Code2Session{
		Code:      code,
		AppId:     global.Config.Code2Session.AppId,
		AppSecret: global.Config.Code2Session.AppSecret,
	}
	api := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	res, err := http.DefaultClient.Get(fmt.Sprintf(api, acs.AppId, acs.AppSecret, acs.Code))
	if err != nil {
		return nil, err.Error()
	}
	if err := json.NewDecoder(res.Body).Decode(&acsJson); err != nil {
		return nil, err.Error()
	}
	if acsJson.ErrCode != 0 || acsJson.OpenId == "" {
		return nil, acsJson.ErrMsg
	}
	rows := global.Db.Table("t_user").Where("open_id = ?", acsJson.OpenId).First(&app.User{}).RowsAffected
	if rows == 0 {
		user := app.User{
			OpenId:      acsJson.OpenId,
			Status:      1,
			CreatedTime: time.Now(),
		}
		row := global.Db.Table("t_user").Create(&user).RowsAffected
		if row == 0 {
			fmt.Println("add app user error")
		}
	}
	return &app.UserInfo{OpenId: acsJson.OpenId}, ""
}
