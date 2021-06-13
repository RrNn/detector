<template>
  <div class="register-page">
    <div class="error" v-if="error">{{ error }}</div>
    <form v-on:submit.prevent="register">
      <legend><h1>Register</h1></legend>
      <auth-input
        name="name"
        type="text"
        placeholder="Name"
        :value="name"
        :onChange="onChange"
      />
      <auth-input
        name="email"
        type="email"
        placeholder="Email"
        :value="email"
        :onChange="onChange"
      />
      <auth-input
        name="address"
        type="text"
        placeholder="Address"
        :value="address"
        :onChange="onChange"
      />
      <auth-input
        name="contact"
        type="text"
        placeholder="Contact"
        :value="contact"
        :onChange="onChange"
      />
      <auth-input
        name="password"
        type="password"
        placeholder="Password"
        :value="password"
        :onChange="onChange"
      />
      <auth-input
        name="confirmPassword"
        type="password"
        placeholder="Confirm Password"
        :value="confirmPassword"
        :onChange="onChange"
      />
      <button type="submit">Register</button>
    </form>
    <span class="backto-login-link">
      Already have an account?
      <button @click="toggle">Login</button>
    </span>
  </div>
</template>

<script>
import http from '../../services';
import emitter from '../../eventHub';
import { storage } from '../../utils';
import { compact } from 'lodash';
import AuthInput from './authInput';
export default {
  emits: ['login-success'],
  props: ['toggle'],
  components: { AuthInput },
  data() {
    return {
      name: '',
      email: '',
      address: '',
      contact: '',
      password: '',
      confirmPassword: '',
      error: '',
    };
  },
  methods: {
    onChange: function (event) {
      const {
        target: { value, name },
      } = event;
      this[name] = value;
    },
    async register() {
      const validate = compact([
        this.name,
        this.email,
        this.address,
        this.contact,
        this.password,
        this.confirmPassword,
      ]);
      if (validate.length !== 6) {
        this.error = 'Some fields are invalid or not filled in';
        return;
      }
      if (this.password !== this.confirmPassword) {
        this.error = 'Password are not the same';
        return;
      }
      try {
        const {
          data: { token },
        } = await http.post('/register', {
          name: this.name,
          email: this.email,
          password: this.password,
          address: this.address,
          contact: this.contact,
        });
        if (token) {
          Promise.resolve(storage.setItem('token', token)).then(() => {
            emitter.$emit('login-success');
            this.$router.push('/');
            window.location = '/'; // check comment in the login file
          });
        }
      } catch (error) {
        const { message } = http.getErrorMessageFromResponse(error);
        this.error = message;
        alert('An error happened while registering you');
        console.info('An error happened while registering you', message);
      }
    },
  },
};
</script>

<style scoped lang="scss">
@import '../../scss/_variables.scss';
.register-page {
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
    input {
      width: 97%;
      padding: 0.5rem;
      font-size: 1.2rem;
      margin-bottom: 1rem;
    }
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
  .backto-login-link {
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
  .error {
    padding: 0.75rem;
    font-size: 1.2rem;
    text-transform: capitalize;
    font-weight: bold;
    background: $app-warning;
    border: 1px solid $app-error;
    border-radius: 0.5rem;
    width: 95%;
    text-align: center;
  }
}
</style>
