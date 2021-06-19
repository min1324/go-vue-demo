import http from '../plugins/axios';

// 用户注册
const register = ({ name, phone, password }) => http.post('auth/register', { name, phone, password });

// 用户登录
const login = ({ phone, password }) => http.post('auth/login', { phone, password });

// 获取用户信息
const info = () => http.get('auth/info');

export default {
  register,
  login,
  info,
};
