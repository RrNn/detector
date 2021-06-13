<template>
  <div id="app" v-if="name">
    <nav>
      <div>
        <router-link :to="{ name: 'home' }">Home</router-link>
        <router-link :to="{ name: 'about' }">About</router-link>
      </div>
      <div>
        <span class="user">{{ name }}</span>
        <button class="logout-btn" @click="logout">Logout</button>
      </div>
    </nav>
    <router-view></router-view>
  </div>
  <auth v-else></auth>
</template>

<script>
import jwt_decode from 'jwt-decode';
import emitter from './eventHub';
import { storage } from './utils';
import Auth from './components/auth/index';
export default {
  name: 'App',
  data: function () {
    return { name: '' };
  },
  components: { Auth },
  methods: {
    checkTokenValidity: function () {
      const token = storage.getItem('token');
      var decoded = token ? jwt_decode(token) : 'unauthorized';
      this.name = decoded?.name || '';
    },
    logout: function () {
      storage.removeItem('token');
      this.name = '';
    },
  },
  mounted() {
    this.checkTokenValidity();
  },
  created() {
    emitter.$on('login-success', () => this.checkTokenValidity(), this);
  },
};
</script>

<style scoped lang="scss">
@import './scss/_variables.scss';
#app {
  nav {
    background: $nav-bg;
    display: flex;
    justify-content: space-between;
    > div {
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
    .user {
      color: $white;
      font-family: sans-serif;
      border: 1px solid black;
      border-radius: 0.5rem;
      padding: 0.5rem 1rem;
      margin-right: 1rem;
    }
    .logout-btn {
      background: transparent;
      border: none;
      color: $white;
      font-size: 1.2rem;
      &:hover {
        cursor: pointer;
      }
    }
  }
}
</style>
