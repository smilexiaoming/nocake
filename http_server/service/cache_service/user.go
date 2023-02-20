package cache_service

import (
	"strconv"
	"strings"

	"nocake/pkg/e"
)

type User struct {
	ID    int
	TagID int
	State int

	PageNum  int
	PageSize int
}

func (u *User) GetUserKey() string {
	return e.CACHE_USER + "_" + strconv.Itoa(u.ID)
}

func (u *User) GetUsersKey() string {
	keys := []string{
		e.CACHE_USER,
		"LIST",
	}

	if u.ID > 0 {
		keys = append(keys, strconv.Itoa(u.ID))
	}
	if u.TagID > 0 {
		keys = append(keys, strconv.Itoa(u.TagID))
	}
	if u.State >= 0 {
		keys = append(keys, strconv.Itoa(u.State))
	}
	if u.PageNum > 0 {
		keys = append(keys, strconv.Itoa(u.PageNum))
	}
	if u.PageSize > 0 {
		keys = append(keys, strconv.Itoa(u.PageSize))
	}

	return strings.Join(keys, "_")
}
