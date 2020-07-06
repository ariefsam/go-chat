<template>
  <q-layout view="lHh Lpr lFf">
    <q-header elevated>
      <q-toolbar>
        <q-btn
          flat
          dense
          round
          icon="menu"
          aria-label="Menu"
          @click="leftDrawerOpen = !leftDrawerOpen"
        />

        <q-toolbar-title>Go-Chat</q-toolbar-title>

        <div>v 0.1</div>
      </q-toolbar>
    </q-header>

    <q-drawer v-model="leftDrawerOpen" show-if-above bordered content-class="bg-grey-1">
      <q-btn color="primary" label="menu">
        <q-menu>
          <q-list style="min-width: 100px">
            <q-item clickable v-close-popup :to="'/channel/create'">
              <q-item-section>Create Channel</q-item-section>
            </q-item>
          </q-list>
        </q-menu>
      </q-btn>
      <q-list>
        <q-item-label header class="text-grey-8">Essential Links</q-item-label>
        <EssentialLink v-for="link in essentialLinks" :key="link.title" v-bind="link" />
      </q-list>

      <q-input
        bottom-slots
        v-model="searchChannel"
        label="Search Channel"
        counter
        maxlength="50"
        @input="onInputSearchChannel"
      >
        <template v-slot:before>
          <q-icon name="flight_takeoff" />
        </template>

        <template v-slot:append>
          <q-icon
            v-if="searchChannel !== ''"
            name="close"
            @click="searchChannel = ''"
            class="cursor-pointer"
          />
          <q-icon name="search" />
        </template>

        <template v-slot:hint>Field hint</template>
      </q-input>

      <q-list>
        <q-item-label header class="text-grey-8">Channels</q-item-label>
        <EssentialLink v-for="link in channelLinks" :key="link.id" v-bind="link" />
      </q-list>
    </q-drawer>

    <q-page-container>
      <router-view  :key="$route.fullPath" />
    </q-page-container>
  </q-layout>
</template>

<script>
import EssentialLink from "components/EssentialLink.vue";
import { getToken } from "pages/deviceID.js";
export default {
  name: "MainLayout",

  components: {
    EssentialLink
  },

  data() {
    return {
      searchChannel: "",
      awaitingSearch: false,
      leftDrawerOpen: false,
      essentialLinks: [
        {
          title: "Login",
          caption: "Login",
          icon: "school",
          link: "/login"
        }
      ],
      channelLinks: []
    };
  },

  methods: {
    onInputSearchChannel() {
      var vm = this;
      
      if (!this.awaitingSearch) {
        setTimeout(() => {
          
          vm.searchChannelAPI();
          // console.log(vm)
          this.awaitingSearch = false;
          
        }, 1000); // 1 sec delay
      }
      this.awaitingSearch = true;
    },
    searchChannelAPI() {
      this.dataSubmit = {
        name: this.searchChannel,
        token: getToken()
      };

      var vm = this;
      this.$axios
        .post("/api/channel/search", this.dataSubmit)
        .then(function(response) {
          var items = [];
          response.data.channels.forEach(element => {
            items.push({
              id: element.ID,
              title: element.Name,
              link: "/channel/detail/" + element.ID,
            })
            vm.channelLinks=items
          });
          
        })
        .catch(function(error) {
          console.log(error);
        });
    }
  }
};
</script>
