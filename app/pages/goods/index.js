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
    if (this.data.cart_number == 0){
      this.deleteOneGoods()
    }else{
      await http.PUT('/cart/set',{ 
        goods_id: this.data.goodsId, 
        options: JSON.stringify({"option":this.data.option_list,"count":this.data.cart_number}),
        open_id: wx.getStorageSync('open_id')
      })
    }
    this.onShow()
  },
  
  /**
     * 选择可选配件
     */
    async additionSelect(e) {
      const propertyindex = e.currentTarget.dataset.propertyindex
      const propertychildindex = e.currentTarget.dataset.propertychildindex

      const goodsAddition = this.data.option_list
      const property = goodsAddition[propertyindex]
      const child = property.item[propertychildindex]
      if (child.active) {
        // 该操作为取消选择
        child.active = false
        this.setData({
          option_list:goodsAddition
        })
        this.addToCart()
        return
      }
      // 单选配件取消所有子栏目选中状态
      if (property.type == 0) {
        property.item.forEach(child => {
          child.active = false
        })
      }
      // 设置当前选中状态
      child.active = true
      this.setData({
        option_list:goodsAddition
      })
      console.log("goodsAddition  ", goodsAddition)
      this.addToCart()
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
    let goodsItem = res.data.data.cart_item
    if (goodsItem){
      for (let i = 0; i < goodsItem.length; i++) {
        let options = JSON.parse(goodsItem[i].options)
        goodsItem[i].cart_number = options.count
        if (String(goodsItem[i].id) == this.data.goodsId) {
          this.data.option_list = options.option
          this.setData({cart_number: options.count})
        }
      }
    }
    console.log("this.data.option_list : ", this.data.option_list)
    this.setData({
      goodsItem: goodsItem,
      option_list: this.data.option_list,
      totalPrice: res.data.data.total_price,
      totalGoodsCount:res.data.data.total_cart,
    })
  },

  // 清空购物车
  async clearCart(){
    await http.DELETE('/cart/clear?open_id='+wx.getStorageSync('open_id'))
    this.setData({ show: false, totalGoodsCount: 0, totalPrice: 0 });
  },

  // 删除购物车某个商品
  async deleteOneGoods(){
    await http.DELETE('/cart/delete?open_id='+wx.getStorageSync('open_id') + "&goods_id="+this.data.goodsId)
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