// pages/order/index.js
import http from '../../utils/http'
import Dialog from '@vant/weapp/dialog/dialog';

Page({

  /**
   * 页面的初始数据
   */
  data: {
    // tabList: ['未完成','已完成'],
    tabList: [{"text":'未付款', value:"1"},{"text":'已取消', value:"2"},{"text":'已付款', value:"3"}, {"text":'配送中', value:"4"},{"text":'已完成', value:"5"}],
    active: 0,
    orderList: []
  },

  tabsChange(event) {
    console.log("event.detail.name : ", event.detail.name)
    this.getOrderList(event.detail.name+1)
  },

  // 添加商品到购物车
  async addToCart() {
    await http.POST('/cart/add',{ 
      goods_id: this.data.goodsId, 
      cart_number: this.data.cart_number,
      open_id: wx.getStorageSync('open_id')
    })
  },

  // 生命周期函数--监听页面显示
  onShow() {
    this.getTabBar().init();
    this.getOrderList(1)
    this.setData({
      active:0
    })
  },

  // 获取订单列表
  async getOrderList(orderType){
    let res = await http.GET('/order/list',{
      status: orderType,
      open_id: wx.getStorageSync('open_id'),
    });
    
    this.setData({
      orderList: res.data.data
    })
  },

  // 取消订单
  async cancelOrder(event){
    let res = await http.PUT('/order/update',{
      order_id: event.currentTarget.id,
      status: 2
    });
  },
    // 付款
    async payOrder(event){
      let res = await http.PUT('/order/update',{
        order_id: event.currentTarget.id,
        status: 3
      });
      this.setData({
        active:2
      })
      this.getOrderList(3)
    },

  // 确认取消订单
  confirmCancelOrder(event){
    let _this = this
    Dialog.confirm({
      title: '确认取消订单吗'
    }).then(() => {
      _this.cancelOrder(event)
      this.setData({
        active:1,
      })
      _this.getOrderList(1)
    });
  },
})

  