import Vue from 'vue';
import axios from 'axios';
import VueAxios from 'vue-axios';
import Vuelidate from 'vuelidate';

import App from './App.vue';
import router from './router';
import store from './store';
import vuetify from './plugins/vuetify';

Vue.config.productionTip = false;
Vue.use(VueAxios, axios);
Vue.use(Vuelidate);

new Vue({
  router,
  store,
  vuetify,
  render: (h) => h(App),
}).$mount('#app');
