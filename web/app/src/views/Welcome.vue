<template>
  <div class="welcome">
    <h1>Welcome {{ userName }}</h1>
    <Admin v-if="userRole==='admin'"/>
    <button className="btn btn-outline-dark" type="button" id="logout" v-on:click="logout()">
      Logout
    </button>
  </div>
</template>

<script>
import Admin from "@/views/Admin";

export default {
  name: "Welcome",
  props: {
    userName: String,
    userRole: String
  },
  components: {
    Admin
  },
  methods: {
    logout() {
      this.$emit('logout')
    }
  }
}
</script>

<style scoped>

</style>