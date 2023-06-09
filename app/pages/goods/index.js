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

  // 返回
  backToHome() {
    wx.navigateBack({
      delta: 1
    })
  },

  // 增加商品数量
  addGoods() {
    this.setData({
      cart_number : this.data.cart_number + 1
    })
    this.addToCart(this.data.goodsId, this.data.cart_number)
  },

    // 当进步器改变时，修改商品数量
    stepperChange(event){
      console.log("event stepperchange ", event)
      this.setData({
        cart_number : event.detail
      })
      this.addToCart(event.currentTarget.id, event.detail)
    },

    // 添加商品到购物车
    async addToCart(id, count) {
      if (count > 0){
        await http.PUT('/cart/set',{ 
          goods_id: id, 
          options: JSON.stringify({"option":this.data.option_list,"count":count}),
          open_id: wx.getStorageSync('open_id')
        })
      }else{
        await http.DELETE('/cart/delete?open_id=' + wx.getStorageSync('open_id') + "&goods_id=" + id)
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
        this.addToCart(this.data.goodsId, this.data.cart_number)
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
      this.addToCart(this.data.goodsId, this.data.cart_number)
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
    if (res.data.data.cart_item){
      for (let index = 0; index < res.data.data.cart_item.length; index++) {
        const element = res.data.data.cart_item[index];
        let option = JSON.parse(element.options)
        res.data.data.cart_item[index].cart_number = option.count
        if (String(element.id) == this.data.goodsId) {
          let options = JSON.parse(element.options)
          this.setData({cart_number: options.count})
        }
      }
    }
    this.setData({
      goodsItem: res.data.data.cart_item,
      totalPrice: res.data.data.total_price,
      totalGoodsCount:res.data.data.total_cart,
    })
  },

  // 清空购物车
  async clearCart(){
    await http.DELETE('/cart/clear?open_id='+wx.getStorageSync('open_id'))
    this.setData({ show: false, totalGoodsCount: 0, totalPrice: 0 ,cart_number: 0});
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