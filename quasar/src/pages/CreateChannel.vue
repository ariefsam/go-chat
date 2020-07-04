<template>
  <q-page padding>
    <div class="row">
      <q-form @submit="onSubmit" class="q-gutter-md">
        <q-input
          class="q-ma-md"
          filled
          v-model="name"
          type="tel"
          label="Channel Name"
          placeholder="62852123456"
          stack-label
        />
        <q-btn type="submit" class="q-ma-md" color="teal">
          <q-icon left name="map" />
          <div>Create</div>
        </q-btn>
      </q-form>
    </div>
  </q-page>
</template>

<script>
import { getDeviceID, setPhoneNumber, getToken } from "pages/deviceID.js";
export default {
  name: "CreateChannel",
  data() {
    return {
      name: "",
      dataSubmit: {}
    };
  },
  mounted() {
    
  },
  methods: {
    onSubmit() {
      
      this.dataSubmit = {
        name: this.name,
        token: getToken()
      };
      setPhoneNumber(this.phoneNumber)

      var vm = this;
      this.$axios
        .post("/api/channel/create", this.dataSubmit)
        .then(function(response) {
          channelID=response.data.channel_id;
          vm.$router.push("/channel/" + channelID)
        })
        .catch(function(error) {
          console.log(error);
        });
    }
  }
};
</script>
