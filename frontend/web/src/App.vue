<template>

  <v-app
    class="overflow-hidden"
    id="top"
  >
    <img id="bgImg">

    <!-- app bar -->
    <v-app-bar
      app
      color="deep-purple accent-4"
      dark
      id="appbar"
    >

      <v-app-bar-nav-icon @click="onDrawer()"></v-app-bar-nav-icon>

      <v-toolbar-title @click="$router.push({name:'Home'})">
        <div color="primary">{{title}}</div>
      </v-toolbar-title>

      <v-spacer></v-spacer>

      <!-- 顶导航右边 -->
      <v-tooltip bottom>
        <template v-slot:activator="{ on, attrs }">
          <v-hover v-slot:default="{ hover }">
            <v-btn
              :elevation="hover ? 16 : 1"
              icon
              color="red"
              v-bind="attrs"
              v-on="on"
            >
              <v-icon>mdi-heart</v-icon>
            </v-btn>
          </v-hover>
        </template>
        <span>收藏</span>
      </v-tooltip>

      <v-tooltip bottom>
        <template v-slot:activator="{ on, attrs }">
          <v-hover v-slot:default="{ hover }">
            <v-btn
              icon
              :elevation="hover ? 16 : 1"
              color="orange"
              dark
              v-bind="attrs"
              v-on="on"
            >
              <v-icon>mdi-magnify</v-icon>
            </v-btn>
          </v-hover>
        </template>
        <span>搜索</span>
      </v-tooltip>

      <!-- 最右选项 -->
      <v-menu
        bottom
        open-on-hover
        offset-y
      >
        <!-- 使用 prop 传数据进入组件 -->
        <template v-slot:activator="{ on, attrs }">
          <v-hover v-slot:default="{ hover }">
            <v-btn
              icon
              :elevation="hover ? 16 : 1"
              color="light-green accent-4"
              v-bind="attrs"
              v-on="on"
            >
              <v-icon>mdi-dots-vertical</v-icon>
            </v-btn>
          </v-hover>
        </template>

        <v-list dense>
          <template v-for="(item, index) in OptionList">
            <v-list-item
              v-if="item.show"
              :key="index"
              @click="$router.push({ name:item.path }) "
            >
              <v-list-item-title v-if="item.show">{{item.title}}</v-list-item-title>
            </v-list-item>
          </template>
        </v-list>
      </v-menu>
    </v-app-bar>
    <!-- app bar end-->

    <!-- 侧边导航 -->
    <!-- navbar -->
    <v-navigation-drawer
      app
      dark
      class="deep-purple accent-4"
      v-model="drawer"
      temporary
      id="navbar"
    >
      <v-img
        :aspect-ratio="16/9"
        :src="navBG"
      >
        <v-row
          align="end"
          class="lightbox white--text pa-2 fill-height"
        >
          <v-list-item class="px-2">
            <v-list-item-avatar
              color="teal"
              size="48"
            >
              <v-img
                v-if="userInfo.img"
                :src="userInfo.img"
              ></v-img>
              <v-img v-else>
                <v-icon dark>mdi-account-circle</v-icon>
              </v-img>
            </v-list-item-avatar>
            <v-list-item-content @click="$router.push({name: 'Profile'})">
              {{userInfo.name}}</v-list-item-content>
          </v-list-item>
        </v-row>
      </v-img>

      <v-divider></v-divider>

      <!-- navbar list -->
      <v-list dense>
        <template v-for="item in items">
          <v-list-item
            v-if="item.show"
            :key="item.title"
            @click="$router.push({name:item.path})"
            link
          >
            <v-list-item-icon>
              <v-icon>{{ item.icon }}</v-icon>
            </v-list-item-icon>

            <v-list-item-content>
              <v-list-item-title>{{ item.title }}</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </template>
      </v-list>
      <template v-slot:append>
        <div class="pa-2">
          <v-btn
            block
            @click="logout"
          >Logout</v-btn>
        </div>
      </template>
    </v-navigation-drawer>
    <!-- navbar end-->

    <!-- main -->

    <v-main>
      <v-container style="my-10">
        <!-- using vue-router -->
        <router-view :key="$route.path"></router-view>
      </v-container>
      <div style="mt-10">
        <foot />
      </div>

    </v-main>
    <!-- main end -->
    <v-btn
      class="mx-2"
      bottom
      right
      fab
      dark
      fixed
      color="purple"
      v-if="scrollBtn"
      href="#top"
    >
      <v-icon dark>mdi-android</v-icon>
    </v-btn>
  </v-app>

</template>

<script>
import { mapActions, mapState } from 'vuex';
import foot from '@/components/Footer.vue';

export default {
  name: 'App',
  data: () => ({
    scrollInvoked: 0,
    scrollBtn: true,
    user: {
      name: '',
      img: 'https://randomuser.me/api/portraits/men/85.jpg',
      // img: 'https://randomuser.me/api/portraits/men/85.jpg',
    },
    title: 'Min',
    navBG: 'http://shp.qpic.cn/ishow/2735040920/1617970550_84828260_22886_sProdImgNo_8.jpg/0',
    drawer: false,
    items: [
      {
        path: 'Home', title: 'Home', icon: 'mdi-home', show: true,
      },
      { title: 'Folder', icon: 'mdi-folder', show: true },
      { title: 'Setting', icon: 'mdi-cog', show: true },
    ],
    OptionList: [
      { title: '登陆', path: 'Login', show: true },
      { title: '注册', path: 'Register', show: true },
    ],
  }),

  computed: mapState({
    userInfo: (state) => state.userModule.userInfo,
  }),
  methods: {
    ...mapActions('userModule', ['logout']),
    onScroll(e) {
      if (e.target.scrollTop > 100) {
        this.scrollBtn = true;
      }
    },
    onDrawer() {
      if (this.drawer) {
        this.drawer = false;
      } else {
        this.drawer = true;
      }
    },
  },
  components: {
    foot,
  },
};
</script>
<style scoped>
#bgImg {
  background-image: url("https://cdn.vuetifyjs.com/images/parallax/material.jpg");
  width: 100%;
  height: 100%;
  position: fixed;
  background-size: 100% 100%;
  overflow: hidden;
}
#appbar,
#navbar {
  filter: alpha(opacity=85);
  -moz-opacity: 0.85;
  -khtml-opacity: 0.85;
  opacity: 0.85;
}
</style>
