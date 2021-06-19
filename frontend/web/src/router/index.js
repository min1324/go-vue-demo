import Vue from 'vue';
import VueRouter from 'vue-router';
import store from '@/store';
import userRoutes from './module/userRouter';
import utilRoutes from './module/fileRouter';
import Home from '../views/Home.vue';

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
  },
  ...userRoutes,
  ...utilRoutes,
];

const router = new VueRouter({
  base: process.env.BASE_URL,
  routes,
});

router.beforeEach((to, from, next) => {
  if (to.meta.auth) { // 判断是否需要登陆
    // 判断已经登陆
    if (store.state.userModule.token) {
      // 判断token的有效性，如果过期，需要后台发放，带上 token 的有效期。
      // 如果 token 无效，需要请求 token.
      next();
    } else {
      // 跳转登陆
      router.push({ name: 'Login' });
    }
  } else {
    next();
  }
});
export default router;
