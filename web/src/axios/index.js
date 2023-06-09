import axios from "axios";
import router from '../router/index';
import { ElMessage } from 'element-plus';

// axios.defaults.baseURL = "http://1.14.106.241/web";
axios.defaults.baseURL = "https://www.nocake.cn/web";

const request = axios.create({
    timeout: 5000,
    headers: {
        'Content-Type': 'application/json;charset=UTF-8'
    }
});

request.interceptors.request.use(config => {
    config.headers['token'] = localStorage.getItem('token')
    return config
})

request.interceptors.response.use(response => {
    if (response.data.code === 200){
        console.log(response)
        return response;
    } else {
        ElMessage.error(response.data.message)
        console.log(response)
        return Promise.reject(response)
    }
},error => {
    console.log(error)
    router.push('/404');
    return Promise.reject(error)
})

export default request;