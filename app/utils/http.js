// 通用请求路径basePath
// var basePath = 'http://127.0.0.1:8000/app'
var basePath = 'https://www.nocake.cn/app'
// var basePath = 'http://1.14.106.241/app'

// 发送GET请求
function GET(url, data = {}){
  return new Promise((resolve, reject) => {
    wx.request({
      url: basePath + url,
      data: data,
      header: { 'content-type': 'application/json' , 'skey':wx.getStorageSync('skey')},
      method: 'GET',
      success: (res) => { console.log(res); resolve(res) },
      fail: (err) => { console.log(err); reject(err) }
    })
  })
}

// 发送POST请求
function POST(url, data = {}){
  return new Promise((resolve, reject) => {
    wx.request({
      url: basePath + url,
      data: data,
      header: { 'content-type': 'application/x-www-form-urlencoded' , 'skey':wx.getStorageSync('skey')},
      method: 'POST',
      success: (res) => { console.log(res); resolve(res) },
      fail: (err) => { console.log(err); reject(err) }
    })
  })
}

// 发送PUT请求
function PUT(url, data = {}){
  return new Promise((resolve, reject) => {
    wx.request({
      url: basePath + url,
      data: data,
      header: { 'content-type': 'application/x-www-form-urlencoded' , 'skey':wx.getStorageSync('skey')},
      method: 'PUT',
      success: (res) => { console.log(res); resolve(res) },
      fail: (err) => { console.log(err); reject(err) }
    })
  })
}

// 发送DELETE请求
function DELETE(url, data = {}){
  return new Promise((resolve, reject) => {
    wx.request({
      url: basePath + url,
      data: data,
      header: { 'content-type': 'application/x-www-form-urlencoded' , 'skey':wx.getStorageSync('skey')},
      method: 'DELETE',
      success: (res) => { console.log(res); resolve(res) },
      fail: (err) => { console.log(err); reject(err) }
    })
  })
}

module.exports = {
  GET, POST, PUT, DELETE
}