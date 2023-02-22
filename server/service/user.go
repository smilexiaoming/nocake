package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"nocake/common"
	"nocake/global"
	"nocake/models/app"
)

type AppUserService struct {
}

func (u *AppUserService) Login(code string) *app.UserInfo {
	var acsJson app.Code2SessionResult
	acs := app.Code2Session{
		Code:      code,
		AppId:     global.Config.Code2Session.AppId,
		AppSecret: global.Config.Code2Session.AppSecret,
	}
	api := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	res, err := http.DefaultClient.Get(fmt.Sprintf(api, acs.AppId, acs.AppSecret, acs.Code))
	if err != nil {
		fmt.Println("微信登录凭证校验接口请求错误")
		return nil
	}
	if err := json.NewDecoder(res.Body).Decode(&acsJson); err != nil {
		fmt.Println("decoder error")
		return nil
	}

	rows := global.Db.Where("open_id = ?", acsJson.OpenId).First(&app.User{}).RowsAffected
	if rows == 0 {
		fmt.Println(acsJson.OpenId)
		user := app.User{
			OpenId:  acsJson.OpenId,
			Status:  1,
			Created: common.NowTime(),
		}
		row := global.Db.Create(&user).RowsAffected
		if row == 0{
			fmt.Println("add app user error")
		}
	}
	return &app.UserInfo{OpenId: acsJson.OpenId}
}
