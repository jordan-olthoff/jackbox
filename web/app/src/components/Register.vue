<template>
  <div>
    <h>Register</h>
    <form name="login-form">
      <div className="mb-3">
        <label htmlFor="username">Username: </label>
        <input type="text" id="username" v-model="input.username"/>
      </div>
      <div className="mb-3">
        <label htmlFor="firstname">Firstname: </label>
        <input type="text" id="firstname" v-model="input.firstname"/>
      </div>
      <div className="mb-3">
        <label htmlFor="lastname">Lastname: </label>
        <input type="text" id="lastname" v-model="input.lastname"/>
      </div>
      <div className="mb-3">
        <label htmlFor="email">Email: </label>
        <input type="text" id="email" v-model="input.email"/>
      </div>
      <div className="mb-3">
        <label htmlFor="password">Password: </label>
        <input type="password" id="password" v-model="input.password"/>
      </div>
      <button className="btn btn-outline-dark" type="submit" v-on:click.prevent="register()">
        Submit
      </button>
    </form>
    <h2 v-show="error!=null">Error: {{ this.error }}</h2>
  </div>
</template>

<script>
export default {
  name: 'Register',
  data() {
    return {
      input: {
        username: "",
        firstname: "",
        lastname: "",
        email: "",
        password: ""
      },
      error: null
    }
  },
  methods: {
    async register() {
      const data = {
        username: this.input.username,
        firstname: this.input.firstname,
        lastname: this.input.lastname,
        email: this.input.email,
        password: this.input.password
      }
      // TODO: separate out actions to an API service
      const response = await fetch("http://localhost:8080/api/v1/user", {
        method: "POST",
        headers: new Headers({'content-type': 'application/json'}),
        body: JSON.stringify(data)
      });

      console.log("register status", response.status)
      let jResp = await response.json();
      if (response.status === 201) {
        this.$emit('registerSuccess', jResp.username);
        this.error = null;
      } else {
        this.error = jResp.message
      }
    }
  },
}
</script>