// pages/order/confirm/index.js
import http from '../../../utils/http'

Page({

  // 页面的初始数据
  data: {
    goodsItem: [],
    totalPrice: 0.00,
    totalGoodsCount: 0,
    show: false,
    showView: true,
    addressList:[]
  },

  // 生命周期函数--监听页面显示
  onShow() {
    this.getCartInfo()
    this.getDefaultAddress()
  },

  // 获取购物车信息
  async getCartInfo(){
    let res = await http.GET('/cart/query', {open_id: wx.getStorageSync('open_id')})
    this.setData({
      goodsItem: res.data.data.cart_item,
      totalPrice: res.data.data.total_price
    })
  },
  async getDefaultAddress(){
    let res = await http.GET('/address/list', {open_id: wx.getStorageSync('open_id')})
    this.setData({
      addressList: res.data.data,
      addressListDefault:res.data.data[0]
    })

  },

  // 提交订单
  submitOrder(){
    this.setData({show: true})
  },

  // 确认支付
  confirmPay(){
    this.setData({showView: false})
    http.POST('/order/submit', {
      open_id: wx.getStorageSync('open_id'),
      sid: parseInt(wx.getStorageSync('sid'))
    })
  },

  // 返回首页
  backToHome(){
    wx.switchTab({url: '/pages/home/index'})
  }
})