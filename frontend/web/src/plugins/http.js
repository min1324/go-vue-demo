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

export default service;
