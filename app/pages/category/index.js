import http from '../../utils/http'

Page({

  // 页面的初始数据
  data: {
    bannerList: [],
    options: [],
    goodsList: [],
    show: false,
    goodsCount: 0,
    goodsItem: [],
    totalPrice: 0.00,
    checkedGoods: ['25','28'],
    totalGoodsCount: 0
  },

  // 跳转到搜索页面
  toSearch(){
    wx.navigateTo({url: '/pages/search/index'})
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad(options) {
    this.getCategoryOption()
    this.getCartInfo()
    this.getBanners()
  },

  // 获取Banner，人气商品
  async getBanners(){
    let res = await http.GET('/goods/hot')
    this.setData({bannerList: res.data.data})
  },

  // 点击banner
  onClickBanner(event){
    console.log("========")
    wx.navigateTo({ url: '/pages/goods/index?id=' + event.currentTarget.id})
  },
  
  // 获取分类选项
  async getCategoryOption() {
    let res = await http.GET('/category/option',{
      level: parseInt(wx.getStorageSync('level')),
      pid: parseInt(wx.getStorageSync('pid'))
    })
    this.setData({options: res.data.data})
    this.getGoodsList(res.data.data[0].id)
  },

  // 获取商品列表
  async getGoodsList(categoryId) {
    let res = await http.GET('/goods/list', {
      category_id: categoryId,
    })
    this.setData({goodsList: res.data.data})
    console.log(this.data.goodsList);
  },

  // 商品分类切换
  changeOption(event){
    let categoryId;
    for (let i = 0; i < this.data.options.length; i++){
      if (i === event.detail.index){ categoryId = this.data.options[i].id }
    }
    this.getGoodsList(categoryId)
  },

  // 查看商品详情
  checkGoodsDetail(event){
    console.log(" detail id ", event.currentTarget)
    wx.navigateTo({ url: '/pages/goods/index?id=' + event.currentTarget.id })
  },

  // 生命周期函数--监听页面显示
  onShow() {
    this.getTabBar().init()
    this.getCategoryOption()
    this.getCartInfo()
  },

  // 展示购物车弹出框
  showCartPopup() {
    if(this.data.show){
      this.setData({ show: false });
    } else {
      this.getCartInfo()
      this.setData({ show: true });
    }
  },
  
  // 获取购物车信息
  async getCartInfo(){
    let res = await http.GET('/cart/query', {open_id: wx.getStorageSync('open_id')})
    this.setData({
      goodsItem: res.data.data.cart_item,
      totalPrice: res.data.data.total_price,
      totalGoodsCount:res.data.data.total_cart,
    })
  },

  // 清空购物车
  async clearCart(){
    await http.DELETE('/cart/clear?open_id='+wx.getStorageSync('open_id'))
    this.setData({ show: false, totalGoodsCount: 0, totalPrice: 0 });
  },

  // 点击结算
  settleAccounts(){
    wx.navigateTo({url: '/pages/order/confirm/index'})
  }
})