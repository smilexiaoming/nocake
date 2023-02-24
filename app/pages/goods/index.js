// pages/goods/index.js
import http from '../../utils/http'
Page({

  // 页面的初始数据
  data: {
    goods: null,
    windowWidth: 0,
    goodsId: 0,
    cart_number: 0,
    show: false,
    goodsItem: [],
    totalPrice: 0.00,
    checked: [],
    totalGoodsCount: 0
  },

  // 生命周期函数--监听页面加载
  async onLoad(options) {
    this.setData({ goodsId: options.id })
    this.setData({ windowWidth: wx.getSystemInfoSync().windowWidth })
    let res = await http.GET('/goods/detail',{ id: options.id });
    this.setData({ goods: res.data.data })
  },

  // 返回首页
  backToHome() {
    wx.switchTab({ url: '/pages/home/index' })
  },

  // 增加商品数量
  addGoods() {
    this.setData({cart_number: 1})
  },

  // 当进步器改变时，修改商品数量
  stepperChange(event){
    this.setData({cart_number: event.detail})
    this.addToCart()
  },

  // 添加商品到购物车
  async addToCart() {
    await http.POST('/cart/add',{ 
      goods_id: this.data.goodsId, 
      cart_number: this.data.cart_number,
      open_id: wx.getStorageSync('open_id')
    })
  },

  // 展示购物车
  showCartPopup() {
    if(this.data.show){
      this.setData({ show: false });
    } else {
      this.setData({ show: true });
    }
    console.log(this.data.show);
  },

  // 生命周期函数--监听页面显示
  async onShow() {
    // 获取购物车信息
    let res = await http.GET('/cart/query', {open_id: wx.getStorageSync('open_id')})
    this.setData({
      goodsItem: res.data.data.cart_item,
      totalPrice: res.data.data.total_price,
      cart_number: this.data.goodsItem[this.data.goodsId],
      totalGoodsCount:res.data.data.total_cart,
    })
    // let totalGoodsCount = 0
    // for (let i = 0; i < this.data.goodsItem.length; i++) {
    //   totalGoodsCount = totalGoodsCount + this.data.goodsItem[i].count
    //   if (this.data.goodsItem[i].id == this.data.goodsId) {
    //     this.setData({cart_number: this.data.goodsItem[i].count})
    //     console.log(this.data.goodsItem[i].count);
    //   }
    // }
    // this.setData({totalGoodsCount: totalGoodsCount})
  },

  // 清空购物车
  async clearCart(){
    await http.DELETE('/cart/clear?open_id='+wx.getStorageSync('open_id'))
  },
  
  // 购物车选中
  onChange(event) {
    this.setData({ checked: event.detail});
  },

  // 点击结算
  settleAccounts(){
    wx.navigateTo({url: '/pages/order/confirm/index'})
  }
})