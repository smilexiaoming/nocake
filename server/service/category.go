package service

import (
	"encoding/json"
	"nocake/constant"
	"nocake/global"
	"nocake/models/app"
	"nocake/models/web"
	"strconv"
	"strings"
	"time"
)

type WebCategoryService struct {
}

type AppCategoryService struct {
}

// 获取树形结构的选项
func getTreeOptions(id int, cateList []web.CategoryList) (option []web.CategoryOption) {
	optionList := make([]web.CategoryOption, 0)
	for _, opt := range cateList {
		if opt.Pid == id && (opt.Level == 1 || opt.Level == 2) {
			option := web.CategoryOption{
				Value:    opt.Id,
				Label:    opt.Name,
				Children: getTreeOptions(opt.Id, cateList),
			}
			if opt.Level == 2 {
				option.Children = nil
			}
			optionList = append(optionList, option)
		}
	}
	return optionList
}

// 创建商品类目
func (c *WebCategoryService) Create(param web.CategoryCreateParam) int64 {
	category := web.Category{
		Name:        param.Name,
		Brief:       param.Brief,
		Keywords:    param.Keywords,
		Pid:         param.Pid,
		Level:       param.Level,
		IconUrl:     param.IconUrl,
		Weight:      param.Weight,
		CreatedTime: time.Now(),
	}
	return global.Db.Debug().Table("t_category").Create(&category).RowsAffected
}

// 删除商品类目
func (c *WebCategoryService) Delete(param web.CategoryDeleteParam) int64 {
	category := app.Category{
		Deleted:     2,
		DeletedTime: time.Now(),
	}
	return global.Db.Debug().Table("t_category").Where("id = ?", param.Id).Updates(category).RowsAffected

}

// 更新商品类目
func (c *WebCategoryService) Update(param web.CategoryUpdateParam) int64 {
	category := web.Category{
		Id:          param.Id,
		Name:        param.Name,
		Brief:       param.Brief,
		Keywords:    param.Keywords,
		Pid:         param.Pid,
		Level:       param.Level,
		IconUrl:     param.IconUrl,
		Weight:      param.Weight,
		UpdatedTime: time.Now(),
	}
	return global.Db.Debug().Table("t_category").Model(&category).Updates(&category).RowsAffected
}

// 获取商品类目列表
func (g *WebCategoryService) GetList(param web.CategoryQueryParam) []web.CategoryList {
	query := &web.Category{
		Name:    param.Name,
		Pid:     param.Pid,
		Level:   param.Level,
		Deleted: 1,
	}
	categoryList := make([]web.CategoryList, 0)
	global.Db.Debug().Table("t_category").Where(query).Find(&categoryList)
	return categoryList
}

// 获取商品类目选项
func (c *WebCategoryService) GetOption() (option []web.CategoryOption) {
	selectList := make([]web.CategoryList, 0)
	global.Db.Table("t_category").Find(&selectList)
	return getTreeOptions(0, selectList)
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
	Db.Debug().Where("level = ? and pid = ? and deleted = 1", param.Level, param.Pid).Find(&categorys)
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
