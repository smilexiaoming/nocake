import http from './utils/http'

App({
  onLaunch() {
    // 自动登录
    wx.login({
      async success (res) {
        if (res.code) {
          let response = await http.POST('/login',{ 
            code: res.code 
          })
          wx.setStorageSync('open_id', response.data.data.openId);
          wx.setStorageSync('skey', response.data.data.skey);
        }
      }
    })
    // 店铺id
    wx.setStorageSync('level', '1');
    wx.setStorageSync('pid', '0');
  },
  globalData: {
    userInfo: null
  }
})
