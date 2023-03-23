// pages/home/index.js
import http from '../../utils/http'
const app = getApp()
Page({

    /**
     * 页面的初始数据
     */
    data: {
        banner_list : {},
        options :[],
        goodsList : [],
        hot_goods : [],
    },

    // 获取banner列表
    async getBanner(){
        this.setData({
            banner_list: ["../../images/微信图片_20220820161406.jpg", "../../images/微信图片_20220820161406.jpg", "../../images/微信图片_20220820161406.jpg"],
        })
    },
    // 获取分类选项
    async getCategoryOption() {
        console.log("options !")
        let res = await http.GET('/category/option',{
            level: parseInt(wx.getStorageSync('level')),
            pid: parseInt(wx.getStorageSync('pid'))
        })
        this.setData({options: res.data.data})
    },
    // 获取热门商品
    async getHotGoods() {
        console.log("hot goods !")
        let res = await http.GET('/goods/hot')
        this.setData({hot_goods: res.data.data})
    },

    // 跳转到分类
    tabClick(event){
        console.log("event ", event)
        app.globalData.option_id = event.currentTarget.id 
        wx.switchTab({ url: '/pages/category/index'})
    },
    // 跳转到搜索页面
    toSearch(){
        wx.tabbar({url: '/pages/search/index'})
    },
    // 查看商品详情
    checkGoodsDetail(event){
        wx.navigateTo({ url: '/pages/goods/index?id=' + event.currentTarget.id })
    },
    /**
     * 生命周期函数--监听页面加载
     */
    onLoad: function (options) {
        this.getBanner()
        this.getCategoryOption()
        this.getHotGoods()
    },

    /**
     * 生命周期函数--监听页面初次渲染完成
     */
    onReady: function () {

    },

    /**
     * 生命周期函数--监听页面显示
     */
    onShow: function () {
      this.getTabBar().init()
        this.setData({
            navHeight: app.globalData.navHeight,
            navTop: app.globalData.navTop,
            windowHeight: app.globalData.windowHeight,
            menuButtonObject: app.globalData.menuButtonObject //小程序胶囊信息
        })
    },

    /**
     * 生命周期函数--监听页面隐藏
     */
    onHide: function () {

    },

    /**
     * 生命周期函数--监听页面卸载
     */
    onUnload: function () {

    },

    /**
     * 页面相关事件处理函数--监听用户下拉动作
     */
    onPullDownRefresh: function () {

    },

    /**
     * 页面上拉触底事件的处理函数
     */
    onReachBottom: function () {

    },

    /**
     * 用户点击右上角分享
     */
    onShareAppMessage: function () {

    }
})