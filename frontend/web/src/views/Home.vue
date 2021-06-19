<template>
  <v-container>
    <v-row no-gutters>
      <v-col
        md="6"
        offset-md="3"
      >
        <div style="height:350px">
          <p
            weight
            class="text--primary  text-center"
          >{{userInfo.name}}</p>
        </div>
      </v-col>
    </v-row>

    <v-container fluid>
      <v-row>
        <v-col
          v-for="card in cards"
          :key="card.title"
          :cols="card.flex"
        >
          <v-card>
            <v-img
              :src="card.src"
              class="white--text"
              gradient="to bottom, rgba(0,0,0,.1), rgba(0,0,0,.5)"
            >
              <v-card-title v-text="card.title"></v-card-title>
            </v-img>

            <v-card-actions>
              <v-spacer></v-spacer>

              <v-btn icon>
                <v-icon>mdi-heart</v-icon>
              </v-btn>

              <v-btn icon>
                <v-icon>mdi-bookmark</v-icon>
              </v-btn>

              <v-btn icon>
                <v-icon>mdi-share-variant</v-icon>
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </v-container>
</template>

<script>
import { mapState } from 'vuex';

export default {
  data: () => ({
    cards: [
      { title: 'Pre-fab homes', src: 'https://cdn.vuetifyjs.com/images/cards/house.jpg', flex: 12 },
      { title: 'Favorite road trips', src: 'https://cdn.vuetifyjs.com/images/cards/road.jpg', flex: 6 },
      { title: 'Best airlines1', src: 'https://cdn.vuetifyjs.com/images/cards/plane.jpg', flex: 6 },
      { title: 'Best airlines2', src: 'https://cdn.vuetifyjs.com/images/cards/plane.jpg', flex: 6 },
      { title: 'Best airlines3', src: 'https://cdn.vuetifyjs.com/images/cards/plane.jpg', flex: 6 },
      { title: 'Best airlines4', src: 'https://cdn.vuetifyjs.com/images/cards/plane.jpg', flex: 6 },
      { title: 'Best airlines5', src: 'https://cdn.vuetifyjs.com/images/cards/plane.jpg', flex: 6 },
      { title: 'Best airlines6', src: 'https://cdn.vuetifyjs.com/images/cards/plane.jpg', flex: 6 },
      { title: 'Best airlines7', src: 'https://cdn.vuetifyjs.com/images/cards/plane.jpg', flex: 6 },
      { title: 'Best airlines8', src: 'https://cdn.vuetifyjs.com/images/cards/plane.jpg', flex: 6 },
      { title: 'Best airlines9', src: 'https://cdn.vuetifyjs.com/images/cards/plane.jpg', flex: 6 },
      { title: 'Best airlines0', src: 'https://cdn.vuetifyjs.com/images/cards/plane.jpg', flex: 6 },

    ],
  }),
  computed: mapState({
    userInfo: (state) => state.userModule.userInfo,
  }),
};
</script>
<style scoped>
</style>
