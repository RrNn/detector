<template>
  <div>
    <login :toggle="togglePages" v-if="login"></login>
    <register :toggle="togglePages" v-else></register>
  </div>
</template>
<script>
import Login from './login';
import Register from './register';
export default {
  components: { Login, Register },
  data() {
    return {
      login: true,
    };
  },
  methods: {
    togglePages: function () {
      this.login = !this.login;
    },
  },
};
</script>
<style></style>
