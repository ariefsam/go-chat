<template>
  <q-page padding>
    <div class="row">
      <q-list>
        <q-item-label header class="text-grey-8">Channels</q-item-label>
        <q-chat-message
          v-for="chat in chats"
          v-bind="chat"
          :key="chat.ID"
          :name="chat.SenderName"
          avatar="~assets/quasar-logo-full.svg"
          :text="[
          chat.Message
        ]"
          stamp="4 minutes ago"
          text-color="white"
          bg-color="primary"
        />
      </q-list>
    </div>
    <div class="row">
      <q-input bottom-slots v-model="text" label="Label" counter maxlength="100">
        <template v-slot:before>
          <q-avatar>
            <img src="~assets/quasar-logo-full.svg" />
          </q-avatar>
        </template>

        <template v-slot:append>
          <q-icon v-if="text !== ''" name="close" @click="text = ''" class="cursor-pointer" />
          <q-icon name="schedule" />
        </template>

        <template v-slot:hint>New Chat</template>

        <template v-slot:after>
          <q-btn round dense flat icon="send" @click="sendMessage" />
        </template>
      </q-input>
    </div>
  </q-page>
</template>

<script>
import ReconnectingWebSocket from "reconnecting-websocket";

import { getToken } from "pages/deviceID.js";
export default {
  name: "CreateChannel",
  data() {
    return {
      chats: [],
      text: "",
      channelID: ""
    };
  },
  deactivated: function() {
    alert("pindah halaman" + this.channelID)
  },
  
  mounted() {
    var channelID = this.$route.params.channelID;
    this.channelID = channelID;

    this.dataSubmit = {
      channelID,
      token: getToken()
    };

    var vm = this;
    this.$axios
      .post("/api/channel/detail", this.dataSubmit)
      .then(function(response) {
        console.log(response.data);
        if (response.data.chats) {
          vm.chats = response.data.chats;
        }
      })
      .catch(function(error) {
        console.log(error);
      });

    const rws = new ReconnectingWebSocket(
      "ws://localhost:8889/api/channel/listen/" + this.channelID
    );

    rws.addEventListener("open", () => {
      rws.send('{"token": "' + getToken() + '"');
    });

    rws.addEventListener("message", msg => {
      console.log(vm.chats);
      console.log(msg.data);

      vm.chats.push(JSON.parse(msg.data));
    });
  },
  methods: {
    sendMessage() {
      this.dataSubmit = {
        channelID: this.channelID,
        text: this.text,
        token: getToken()
      };

      var vm = this;
      this.$axios
        .post("/api/channel/chat/create", this.dataSubmit)
        .then(function(response) {
          if (response.data.chat) {
            vm.text = "";
          }
        })
        .catch(function(error) {
          console.log(error);
        });
    }
  }
};
</script>
