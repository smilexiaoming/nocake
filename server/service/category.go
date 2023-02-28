package service

import (
	"encoding/json"
	"nocake/constant"
	"nocake/global"
	"nocake/models/app"
	"strconv"
	"strings"
	"time"
)

type AppCategoryService struct {
}

func (c *AppCategoryService) GetOption(param app.CategoryQueryParam) (option []app.CategoryOption, errMessage string) {
	pid := strconv.Itoa(param.Pid)
	level := strconv.Itoa(param.Level)
	key := strings.Join([]string{"category", "option", pid, level}, ":")
	redisContent := global.Rdb.Get(ctx, key).Val()
	if redisContent != "" {
		err := json.Unmarshal([]byte(redisContent), &option)
		if err != nil {
			return nil, err.Error()
		}
		return option, ""
	}
	categorys := make([]app.Category, 0)
	categoryOptions := make([]app.CategoryOption, 0)
	Db := global.Db.Table("t_category")
	if param.Level > 1 && param.Pid == 0 {
		errMessage = constant.ParamInvalid
		return nil, errMessage
	}
	Db.Debug().Where("level = ? and pid = ?", param.Level, param.Pid).Find(&categorys)
	for _, item := range categorys {
		option := app.CategoryOption{
			Id:   item.Id,
			Text: item.Name,
		}
		categoryOptions = append(categoryOptions, option)
	}
	redisByte, _ := json.Marshal(categoryOptions)
	redisContent = string(redisByte)
	global.Rdb.Set(ctx, key, redisContent, time.Hour)
	return categoryOptions, ""
}
