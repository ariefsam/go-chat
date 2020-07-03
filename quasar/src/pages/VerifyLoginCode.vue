<template>
  <q-page padding>
    <div class="row">
      <q-form @submit="onSubmit" class="q-gutter-md">
        <q-input
          class="q-ma-md"
          filled
          v-model="code"
          type="tel"
          label="Code"
          placeholder
          stack-label
        />
        <q-btn type="submit" class="q-ma-md" color="primary">
          <q-icon left name="map" />
          <div>Verify</div>
        </q-btn>
      </q-form>
    </div>
  </q-page>
</template>

<script>
import { getDeviceID, getPhoneNumber, setToken } from "pages/deviceID.js";
export default {
  name: "PageVerifyLogin",
  data() {
    return {
      code: "",
      dataSubmit: {}
    };
  },
  mounted() {
    console.log(getDeviceID);
  },
  methods: {
    onSubmit() {
      this.dataSubmit = {
        phoneNumber: getPhoneNumber(),
        deviceID: getDeviceID(),
        verificationCode: this.code
      };

      var vm = this;
      this.$axios
        .post("/api/login/verify", this.dataSubmit)
        .then(function(response) {
          if (typeof response.data.token != "undefined") {
            setToken(response.data.token);
          }
        })
        .catch(function(error) {
          console.log(error);
        });
    }
  }
};
</script>
