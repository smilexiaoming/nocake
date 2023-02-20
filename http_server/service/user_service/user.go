package User_service

import (
	"encoding/json"

	"nocake/http_server/service/cache_service"
	"nocake/models"
	"nocake/pkg/gredis"
	"nocake/pkg/logging"
)

type User struct {
	ID         int
	WxId       string
	UserName   string
	NickName   string
	MiniOpenid string
	VipLevel   string
	HeadUrl    string

	PageNum  int
	PageSize int
}

func (u *User) Add() error {
	User := map[string]interface{}{
		"wx_id":       u.WxId,
		"user_name":   u.UserName,
		"nick_name":   u.NickName,
		"mini_openid": u.MiniOpenid,
		"vip_level":   u.VipLevel,
		"head_url":    u.HeadUrl,
	}

	if err := models.AddUser(User); err != nil {
		return err
	}

	return nil
}

func (u *User) Edit() error {
	return models.EditUser(u.ID, map[string]interface{}{
		"wx_id":       u.WxId,
		"user_name":   u.UserName,
		"nick_name":   u.NickName,
		"mini_openid": u.MiniOpenid,
		"vip_level":   u.VipLevel,
		"head_url":    u.HeadUrl,
	})
}

func (u *User) Get() (*models.User, error) {
	var cacheUser *models.User

	cache := cache_service.User{ID: u.ID}
	key := cache.GetUserKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheUser)
			return cacheUser, nil
		}
	}

	User, err := models.GetUser(u.ID)
	if err != nil {
		return nil, err
	}

	gredis.Set(key, User, 3600)
	return User, nil
}

func (a *User) GetAll() ([]*models.User, error) {
	var (
		Users, cacheUsers []*models.User
	)

	cache := cache_service.User{
		PageNum:  a.PageNum,
		PageSize: a.PageSize,
	}
	key := cache.GetUsersKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheUsers)
			return cacheUsers, nil
		}
	}

	Users, err := models.GetUsers(a.PageNum, a.PageSize, a.getMaps())
	if err != nil {
		return nil, err
	}

	gredis.Set(key, Users, 3600)
	return Users, nil
}

func (a *User) Delete() error {
	return models.DeleteUser(a.ID)
}

func (a *User) ExistByID() (bool, error) {
	return models.UserService().ExistUserByID(a.ID)
}

func (a *User) Count() (int, error) {
	return models.GetUserTotal(a.getMaps())
}

func (a *User) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0
	return maps
}
