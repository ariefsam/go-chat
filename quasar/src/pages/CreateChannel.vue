<template>
  <q-page padding>
    <div class="row">
      <q-form @submit="onSubmit" class="q-gutter-md">
        <q-input
          class="q-ma-md"
          filled
          v-model="phoneNumber"
          type="tel"
          label="Phone Number"
          placeholder="62852123456"
          stack-label
        />
        <q-btn type="submit" class="q-ma-md" color="teal">
          <q-icon left name="map" />
          <div>Login</div>
        </q-btn>
      </q-form>
    </div>
  </q-page>
</template>

<script>
import { getDeviceID, setPhoneNumber } from "pages/deviceID.js";
export default {
  name: "CreateChannel",
  data() {
    return {
      phoneNumber: "",
      dataSubmit: {}
    };
  },
  mounted() {
    
  },
  methods: {
    onSubmit() {
      
      this.dataSubmit = {
        phoneNumber: this.phoneNumber,
        deviceID: getDeviceID()
      };
      setPhoneNumber(this.phoneNumber)

      var vm = this;
      this.$axios
        .post("/api/login", this.dataSubmit)
        .then(function(response) {
          vm.$router.push("/login/verify")
        })
        .catch(function(error) {
          console.log(error);
        });
    }
  }
};
</script>
