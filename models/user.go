package models

import (
	"github.com/jinzhu/gorm"
)
var UserService = func () *User {
	return &User{}
}

type User struct {
	Model
	WxId       string `db:"wx_id"`
	UserName   string `db:"user_name"`
	NickName   string `db:"nick_name"`
	MiniOpenid string `db:"mini_openid"`
	VipLevel   string `db:"vip_level"`
	HeadUrl    string `db:"head_url"`
}

// ExistUserByID checks if an User exists based on ID
func (u *User) ExistUserByID(id int) (bool, error) {
	err := db.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(u).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if u.ID > 0 {
		return true, nil
	}

	return false, nil
}

// GetUserTotal gets the total number of Users based on the constraints
func GetUserTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&User{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// GetUsers gets a list of Users based on paging constraints
func GetUsers(pageNum int, pageSize int, maps interface{}) ([]*User, error) {
	var Users []*User
	err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&Users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return Users, nil
}

// GetUser Get a single User based on ID
func GetUser(id int) (*User, error) {
	var User User
	err := db.Where("id = ? AND deleted_on = ? ", id, 0).First(&User).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &User, nil
}

// EditUser modify a single User
func EditUser(id int, data interface{}) error {
	if err := db.Model(&User{}).Where("id = ? AND deleted_on = ? ", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

// AddUser add a single User
func AddUser(data map[string]interface{}) error {
	User := User{
		WxId:       data["wx_id"].(string),
		UserName:   data["user_name"].(string),
		NickName:   data["nick_name"].(string),
		MiniOpenid: data["mini_openid"].(string),
		VipLevel:   data["vip_level"].(string),
		HeadUrl:    data["head_url"].(string),
	}
	if err := db.Create(&User).Error; err != nil {
		return err
	}

	return nil
}

// DeleteUser delete a single User
func DeleteUser(id int) error {
	if err := db.Where("id = ?", id).Delete(User{}).Error; err != nil {
		return err
	}

	return nil
}

// CleanAllUser clear all User
func CleanAllUser() error {
	if err := db.Unscoped().Where("deleted_on != ? ", 0).Delete(&User{}).Error; err != nil {
		return err
	}

	return nil
}
