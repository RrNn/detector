<template>
  <div id="app">
    <nav>
      <router-link :to="{ name: 'home' }">Home</router-link>
      <router-link :to="{ name: 'about' }">About</router-link>
    </nav>
    <router-view></router-view>
  </div>
</template>

<script>

export default {
  name: 'App',
  components: {},
}
</script>

<style scoped lang="scss">
@import "./scss/_variables.scss";
#app {
  nav {
    background: $nav-bg;
    height: 3.2rem;
    display: flex;
    align-items: center;
    a {
      color: $white;
      font-size: 1.2rem;
      text-decoration: none;
      height: 100%;
      display: flex;
      justify-content: center;
      align-items: center;
      min-width: 5rem;
      &:hover {
        background: $nav-active-link;
      }
      &.router-link-active {
        background: $nav-active-link;
      }
    }
  }
}
</style>
