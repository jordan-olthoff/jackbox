<template>
  <div>
    <h>Login</h>
    <form name="login-form">
      <div className="mb-3">
        <label htmlFor="username">Username: </label>
        <input type="text" id="username" v-model="input.username"/>
      </div>
      <div className="mb-3">
        <label htmlFor="password">Password: </label>
        <input type="password" id="password" v-model="input.password"/>
      </div>
      <button className="btn btn-outline-dark" type="submit" v-on:click.prevent="login()">
        Submit
      </button>
    </form>
    <h2 v-show="error!=null">Error: {{ this.error }}</h2>
  </div>
</template>

<script>
export default {
  name: 'Login',
  data() {
    return {
      input: {
        username: "",
        password: ""
      },
      error: null
    }
  },
  methods: {
    async login() {
      const data = {
        username: this.input.username,
        password: this.input.password
      }
      // TODO: separate actions into API service
      const response = await fetch("http://localhost:8080/api/v1/auth/login", {
        method: "POST",
        headers: new Headers({'content-type': 'application/json'}),
        body: JSON.stringify(data)
      });
      let jResp = await response.json();
      console.log(jResp)
      if (response.status === 200) {
        this.error = null;
        let tokenWrapper = jResp.token;
        console.log(tokenWrapper)
        this.$emit('loginSuccess', tokenWrapper['token']);
      } else {
        this.error = jResp.message
      }
      return jResp;
    }
  },
}
</script>