<template>
  <div class="login-page">
    <form v-on:submit.prevent="login">
      <legend><h1>Login</h1></legend>
      <auth-input
        type="email"
        name="email"
        placeholder="Email"
        :value="email"
        :onChange="onChange"
      />
      <auth-input
        type="password"
        name="password"
        placeholder="Password"
        :value="password"
        :onChange="onChange"
      />
      <button type="submit">Login</button>
    </form>
    <span class="goto-register-link">
      Don't have an account?
      <button @click="toggle">Register</button>
    </span>
  </div>
</template>

<script>
import http from '../../services';
import emitter from '../../eventHub';
import { storage } from '../../utils';
import AuthInput from './authInput';
export default {
  emits: ['login-success'],
  props: ['toggle'],
  components: { AuthInput },
  data() {
    return {
      email: '',
      password: '',
    };
  },
  methods: {
    onChange: function (event) {
      const {
        target: { value, name },
      } = event;
      this[name] = value;
    },
    async login() {
      try {
        const {
          data: { token },
        } = await http.post('/auth/login', {
          email: this.email,
          password: this.password,
        });
        if (token) {
          Promise.resolve(storage.setItem('token', token)).then(() => {
            emitter.$emit('login-success');
            // this.$router.push("/");
            /* this makes me feel silly, need a way to make sure the axios
            instance after login has the Authorization header with the token,
            that way i wont have to manually refresh to pick it from storage
            by using this stupid window.location. As it is now the event emitter
            above is even useless 🚧
            */
            this.$router.push('/');
            window.location = '/';
          });
        }
      } catch (error) {
        alert('password or email is incorrect');
        console.info('An error happened while logging you in', error.response);
      }
    },
  },
};
</script>

<style scoped lang="scss">
@import '../../scss/_variables.scss';
.login-page {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  height: 80vh;
  form {
    width: 50%;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    button {
      width: 100%;
      font-size: 2rem;
      background: $app-neutral;
      color: $white;
      border-radius: 0.2rem;
      margin: 0;
      &:hover {
        cursor: pointer;
        background-color: $nav-bg;
      }
    }
  }
  .goto-register-link {
    margin-top: 0.5rem;
    font-size: 1.2rem;
    font-family: sans-serif;
    button {
      font-size: 1.2rem;
      background: transparent;
      color: $black;
      width: fit-content;
      border: none;
      color: $nav-active-link;
      &:hover {
        cursor: pointer;
      }
    }
  }
}
</style>
