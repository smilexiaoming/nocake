import { areaList } from '@vant/area-data';
import http from '../../../utils/http'
Page({

  // 页面的初始数据
  data: {
    areaList,
    id: '',
    name: '',
    tel: '',
    detail:"",
    isDefault: 0,
    show: false,
    checked: false,
    submitType:"submit",
  },

  // 获取地理位置
  async chooseLocation(){
    wx.chooseLocation({
      altitude: 'altitude',
      highAccuracyExpireTime: 0,
      isHighAccuracy: true,
      type: 'type',
      success: (result) => {
        console.log("res.name ", result.name)
        console.log("res.address ", result.address)
        this.setData({
          detail: result.name + result.address,
        })
      },
      fail: (res) => {},
      complete: (res) => {},
  })
  },
  // 改变默认地址
  changeSwitch({detail}){
    this.setData({ checked: detail })
    if (detail) {
      this.setData({ isDefault: 1 })
    }
  },

  // 生命周期函数--监听页面加载
  async onLoad(options){
    if (options.id > 0){
      let res = await http.GET('/address/detail',{ address_id: options.id, open_id:wx.getStorageSync('open_id') });
      if(res.data.code === 200){
        var response = res.data.data;
        this.setData({
          id: response.id,
          name: response.name,
          tel: response.tel,
          province: response.province,
          city: response.city,
          county: response.county,
          detail: response.detail,
          area: response.province + ' ' + response.city + ' ' + response.county,
          isDefault: response.is_default,
          submitType:"update"
        })
        if (this.data.isDefault === 1){
          this.setData({ checked: true })
        }
      }
    }
  },

  // 保存地址
  async submitForm(options) {
    console.log("this.data.submitType ", this.data.submitType)
    let data = {
      address_id: this.data.id,
      name: this.data.name,
      tel: this.data.tel,
      detail: this.data.detail,
      is_default: this.data.isDefault,
      open_id: wx.getStorageSync('open_id')
    }
    let res
    if (this.data.submitType == "update"){
      res = await http.PUT('/address/'+this.data.submitType,data)
    }else{
      res = await http.POST('/address/'+this.data.submitType,data)
    }
    if(res.data.code === 200){ 
      wx.navigateBack({ url: '/pages/address/list/index' })
    }
  }
})