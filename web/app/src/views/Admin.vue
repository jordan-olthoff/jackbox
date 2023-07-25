<template>
  <div class="user-list">
    <ul>
      <li v-for="(user, i) in users" v-bind:key="i">
        {{ user.username }} | {{ user.firstname }} | {{ user.lastname }} | {{ user.email }}
      </li>
    </ul>
  </div>
</template>

<script>

export default {
  name: "Admin",
  data() {
    return {
      users: []
    }
  },
  methods: {
    async fetchUsers() {
      // TODO: separate out actions to an API service
      let token = localStorage.getItem('token')
      const response = await fetch("http://localhost:8080/api/v1/users", {
        method: "GET",
        headers: new Headers(
            {
              'content-type': 'application/json',
              'Authorization': `Bearer ${token}`
            })
      });
      let jResp = await response.json();
      console.log(jResp)
      this.users = jResp
    }
  },
  beforeMount() {
    this.fetchUsers()
  }
}
</script>

<style scoped>

</style>