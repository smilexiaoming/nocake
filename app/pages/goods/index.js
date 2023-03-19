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
    totalGoodsCount: 0,
    option_list : [],
  },

  // 生命周期函数--监听页面加载
  async onLoad(options) {
    this.setData({ goodsId: options.id })
    this.setData({ windowWidth: wx.getSystemInfoSync().windowWidth })
    let res = await http.GET('/goods/detail',{ goods_id: options.id });
    this.setData({ goods: res.data.data })
    let option_list = JSON.parse(res.data.data.options)
    console.log("obj ", option_list)
    this.setData({
      option_list:option_list
    })
    console.log(option_list)
  },

  // 返回首页
  backToHome() {
    wx.switchTab({ url: '/pages/home/index' })
  },

  // 增加商品数量
  addGoods() {
    this.setData({cart_number: 1})
    this.addToCart()
  },

  // 当进步器改变时，修改商品数量
  stepperChange(event){
    console.log(event)
    this.setData({cart_number: event.detail})
    this.addToCart()
  },

  // 添加商品到购物车
  async addToCart() {
    await http.PUT('/cart/set',{ 
      goods_id: this.data.goodsId, 
      options: this.data.options,
      open_id: wx.getStorageSync('open_id')
    })
    this.onShow()
  },
    /**
     * 可选项
     */
    async additionSelect(e) {
      const propertyindex = e.currentTarget.dataset.propertyindex
      const propertychildindex = e.currentTarget.dataset.propertychildindex

      const goodsAddition = this.data.goodsAddition
      const property = goodsAddition[propertyindex]
      const child = property.items[propertychildindex]
      if (child.active) {
        // 该操作为取消选择
        child.active = false
        this.setData({
          goodsAddition
        })
        this.calculateGoodsPrice()
        return
      }
      // 单选配件取消所有子栏目选中状态
      if (property.type == 0) {
        property.items.forEach(child => {
          child.active = false
        })
      }
      // 设置当前选中状态
      child.active = true
      this.setData({
        goodsAddition
      })
      this.calculateGoodsPrice()
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
      totalGoodsCount:res.data.data.total_cart,
    })
    if (this.data.goodsItem){
      for (let i = 0; i < this.data.goodsItem.length; i++) {
        if (String(this.data.goodsItem[i].id) == this.data.goodsId) {
          console.log(this.data.goodsItem[i].cart_number);
          this.setData({cart_number: this.data.goodsItem[i].cart_number})
        }
      }
    }
  },

  // 清空购物车
  async clearCart(){
    await http.DELETE('/cart/clear?open_id='+wx.getStorageSync('open_id'))
    this.setData({ show: false, totalGoodsCount: 0, totalPrice: 0 });
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