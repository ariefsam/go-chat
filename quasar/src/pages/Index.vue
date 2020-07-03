<template>
  <q-page class="flex flex-center">
    <img
      alt="Quasar logo"
      src="~assets/quasar-logo-full.svg"
    >
  </q-page>
</template>

<script>
import { getToken } from "pages/deviceID.js";
export default {
  name: 'PageIndex',
  mounted(){
    console.log(getToken())
  }
}
</script>
