package service

import (
	"nocake/global"
	"nocake/models/app"
)

type AppCategoryService struct {
}

func (c *AppCategoryService) GetOption(param app.CategoryQueryParam) (option []app.CategoryOption) {
	categorys := make([]app.Category, 0)
	categoryOptions := make([]app.CategoryOption, 0)
	global.Db.Table("t_category").Where("Level = ?", param.Level).Find(&categorys)
	for _, item := range categorys {
		option := app.CategoryOption{
			Id:   item.Id,
			Text: item.Name,
		}
		categoryOptions = append(categoryOptions, option)
	}
	return categoryOptions
}
