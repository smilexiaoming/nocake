import { areaList } from '@vant/area-data';
import http from '../../../utils/http'

Page({

  // 页面的初始数据
  data: {
    areaList,
    id: '',
    name: '',
    mobile: '',
    province: '',
    city: '',
    district: '',
    detailedAddress: '',
    isDefault: 0,
    area: '',
    show: false,
    checked: false,
    submitType:"submit",
  },

  // 展示地区选择弹出框
  showPopup(){
    this.setData({ show: true })
  },

  // 确实选择
  confirmSelect(event){
    this.setData({ show: false });
    this.setData({ province: event.detail.values[0].name });
    this.setData({ city: event.detail.values[1].name });
    this.setData({ county: event.detail.values[2].name });
    this.setData({ area: event.detail.values[0].name + " " + event.detail.values[1].name + " " + event.detail.values[2].name});
  },

  // 取消选择
  cancelSelect({detail}){
    this.setData({ show: false })
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
      let res = await http.GET('/address/detail',{ id: options.id, open_id:wx.getStorageSync('open_id') });
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
    let res = await http.POST('/address/'+this.data.submitType,{
      address_id: this.data.id,
      name: this.data.name,
      tel: this.data.tel,
      province: this.data.province,
      city: this.data.city,
      county: this.data.county,
      detail: this.data.detail,
      is_default: this.data.isDefault,
      open_id: wx.getStorageSync('open_id')
    })
    if(res.data.code === 200){ 
      wx.navigateBack({ url: '/pages/address/list/index' })
    }
  }
})