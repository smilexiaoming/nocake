package service

import (
	"fmt"
	"nocake/constant"
	"nocake/global"
	"nocake/models/app"
)

type AppCategoryService struct {
}

func (c *AppCategoryService) GetOption(param app.CategoryQueryParam) (option []app.CategoryOption, errMessage string) {
	fmt.Printf("param: %v\n", param)
	categorys := make([]app.Category, 0)
	categoryOptions := make([]app.CategoryOption, 0)
	Db := global.Db.Table("t_category")
	if param.Level > 1 && param.Pid == 0 {
		errMessage = constant.ParamInvalid
		return nil, errMessage
	}
	Db.Where("level = ? and pid = ?", param.Level, param.Pid).Find(&categorys)
	for _, item := range categorys {
		option := app.CategoryOption{
			Id:   item.Id,
			Text: item.Name,
		}
		categoryOptions = append(categoryOptions, option)
	}
	return categoryOptions, ""
}
