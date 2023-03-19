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

    // ---------------检测navbar高度
    let menuButtonObject = wx.getMenuButtonBoundingClientRect();
    console.log("小程序胶囊信息",menuButtonObject)
    wx.getSystemInfo({
      success: res => {
        let statusBarHeight = res.statusBarHeight,
          navTop = menuButtonObject.top,//胶囊按钮与顶部的距离
          navHeight = statusBarHeight + menuButtonObject.height + (menuButtonObject.top - statusBarHeight)*2;//导航高度
        this.globalData.navHeight = navHeight;
        this.globalData.navTop = navTop;
        this.globalData.windowHeight = res.windowHeight;
        this.globalData.menuButtonObject = menuButtonObject;
        console.log("navHeight",navHeight);
      },
      fail(err) {
        console.log(err);
      }
    })

    // 品类初始化
    wx.setStorageSync('level', '1');
    wx.setStorageSync('pid', '0');
  },
  globalData: {
    userInfo: null
  }
})
