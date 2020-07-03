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
    </q-drawer>

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script>
import EssentialLink from "components/EssentialLink.vue";

export default {
  name: "MainLayout",

  components: {
    EssentialLink
  },

  data() {
    return {
      leftDrawerOpen: false,
      essentialLinks: [
        {
          title: "Login",
          caption: "Login",
          icon: "school",
          link: "/login"
        }
      ]
    };
  }
};
</script>
