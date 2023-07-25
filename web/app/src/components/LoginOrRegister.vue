<template>
  <div>
    <Login v-if="activeButton==='login'" @login-success="successfulLoginAction"/>
    <Register v-else @register-success="successfulRegisterAction"/>
    <button className="btn btn-outline-dark"
            type="button"
            v-on:click="setActive('register')"
            v-if="activeButton==='login'">
      Register
    </button>
    <button className="btn btn-outline-dark"
            type="button"
            v-on:click="setActive('login')"
            v-else>
      Login
    </button>
    <h3 v-show="lastRegisteredUser!==''">{{ lastRegisteredUser }} registered</h3>
  </div>
</template>

<script>
import Login from "@/components/Login";
import Register from "@/components/Register";
export default {
  name: "LoginOrRegister",
  components: {Register, Login},
  data() {
    return {
      activeButton: 'login',
      lastRegisteredUser: ''
    }
  },
  methods: {
    setActive(form) {
      this.activeButton = form;
    },
    successfulRegisterAction(username) {
      console.log("successfulRegisterAction:", username)
      this.activeButton = 'login'
      this.lastRegisteredUser = username
    },
    successfulLoginAction(token) {
      console.log("successfulLoginAction:", token)
      this.$emit("login", token)
    }
  }
}
</script>

<style scoped>

</style>