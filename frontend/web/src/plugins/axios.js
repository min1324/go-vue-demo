import axios from 'axios';
import storageService from '../service/storageService';

const url = window.index.url ? window.index.url : 'http://127.0.0.1:8000';

const service = axios.create({
  // baseURL: glob.url,
  baseURL: url,
  timeout: 1000 * 5,
});

// Add a request interceptor
service.interceptors.request.use((config) => {
  // Do something before request is sent

  Object.assign(config.headers, { Authorization: `Bearer ${storageService.get(storageService.USER_TOKEN)}` });
  return config;
}, (error) => {
  // Do something with request error
  Promise.reject(error);
});

// 添加响应拦截器
service.interceptors.response.use((response) => {
  // 2xx 范围内的状态码都会触发该函数。
  // 对响应数据做点什么
  return response;
}, ((error) => {
  if (error.response) {
    // 请求成功发出且服务器也响应了状态码，但状态代码超出了 2xx 的范围
  } else if (error.request) {
    // 请求已经成功发起，但没有收到响应
    // `error.request` 在浏览器中是 XMLHttpRequest 的实例，
    // 而在node.js中是 http.ClientRequest 的实例
  } else {
    // 发送请求时出了点问题
  }
  return Promise.reject(error);
}));

export default service;
