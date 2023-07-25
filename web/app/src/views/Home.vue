<template>
  <div class="home">
    <LoginOrRegister v-if="!loggedIn" @login="loginAction"/>
    <Welcome :userName="userName" :userRole="userRole" v-else-if="loggedIn" @logout="logoutAction"/>
  </div>
</template>

<script>
import {decodeToken, fetchToken} from "@/services/TokenFetcher";
import LoginOrRegister from "@/components/LoginOrRegister";
import Welcome from "@/views/Welcome";


export default {
  name: 'Home',
  data() {
    return {
      loggedIn: false,
      userRole: '',
      userName: '',
      fetcher: null
    }
  },
  components: {
    LoginOrRegister,
    Welcome
  },
  methods: {
    logoutAction() {
      console.log("received logout action")
      this.stopFetcher()
      localStorage.removeItem('token')
      this.loggedIn = false
      this.userRole = ''
      this.userName = ''
    },
    loginAction(token) {
      console.log("received login action", token)
      localStorage.setItem('token', token)

      let parsedToken = decodeToken(token)
      console.log(parsedToken)
      this.userRole = parsedToken.role
      this.userName = parsedToken["sub"]
      localStorage.setItem('userName', this.userName)
      this.loggedIn = true
      this.startFetcher(parsedToken)
    },
    startFetcher(token) {
      const expiry = token.exp
      const now = Math.round(Date.now() / 1000)
      if (now < expiry) {
        this.fetcher = setInterval(fetchToken, 60 * 1000,
            localStorage.getItem('userName'), localStorage.getItem('token'))
      }
    },
    stopFetcher() {
      if (this.fetcher) clearInterval(this.fetcher);
    }
  },
  beforeMount() {
    let token = localStorage.getItem('token')
    if (token) {
      this.loginAction(token)
    }
  }
}
</script>
