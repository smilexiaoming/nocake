package service

import (
	"nocake/global"
	"nocake/models/app"
	"time"
)

type AppAddressService struct {
}

func (a *AppAddressService) Add(param app.AddressAddParam) int64 {
	address := app.Address{
		Name:        param.Name,
		OpenId:      param.OpenId,
		Province:    param.Province,
		City:        param.City,
		County:      param.County,
		Detail:      param.Detail,
		AreaCode:    param.AreaCode,
		PostalCode:  param.PostalCode,
		Tel:         param.Tel,
		IsDefault:   param.IsDefault,
		CreatedTime: time.Now(),
	}
	if param.IsDefault == 1 {
		global.Db.Debug().Table("t_address").Where("open_id = ? and is_default = 1", param.OpenId).Update("is_default", 0)
	}
	return global.Db.Debug().Table("t_address").Create(&address).RowsAffected
}

func (a *AppAddressService) Updata(param app.AddressUpdateParam) int64 {
	address := app.Address{
		Name:        param.Name,
		Province:    param.Province,
		City:        param.City,
		County:      param.County,
		Detail:      param.Detail,
		AreaCode:    param.AreaCode,
		PostalCode:  param.PostalCode,
		Tel:         param.Tel,
		IsDefault:   param.IsDefault,
		UpdatedTime: time.Now(),
	}
	if param.IsDefault == 1 {
		global.Db.Debug().Table("t_address").Where("open_id = ? and is_default = 1", param.OpenId).Update("is_default", 0)
	}
	return global.Db.Debug().Table("t_address").Where("id = ? and open_id = ?", param.AddressId, param.OpenId).Updates(address).RowsAffected
}

func (a *AppAddressService) Delete(param app.AddressDeleteParam) int64 {
	address := app.Address{
		Deleted:     1,
		DeletedTime: time.Now(),
	}
	return global.Db.Debug().Table("t_address").Where("id = ? and open_id = ?", param.AddressId, param.OpenId).Updates(address).RowsAffected
}

func (a *AppAddressService) GetList(param app.AddressListQueryParam) []app.Address {
	addressList := make([]app.Address, 0)
	global.Db.Debug().Table("t_address").Where("open_id = ? and deleted != 1", param.OpenId).Order("is_default desc").Find(&addressList)
	return addressList
}
