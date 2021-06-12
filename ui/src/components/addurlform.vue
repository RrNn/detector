<template>
  <form class="url-input-form" v-on:submit.prevent="onSubmit">
    <input type="text" placeholder="Add URL/Link to monitor" v-model="link" />
    <button>{{ adding ? 'Adding URL...' : 'Add URL' }}</button>
    <span v-if="error" class="error">
      <span>{{ startCase(error) }}</span>
      <span class="remove-error" :onClick="removeError">X</span>
    </span>
    <div v-if="resp" class="success">
      <span
        ><strong>First Check Status Code:</strong>
        {{ resp.firstCheckStatusCode }}</span
      >
      <span><strong>Time Check:</strong> {{ resp.lastChecked }}</span>
      <span><strong>URL:</strong> {{ resp.url }}</span>
    </div>
  </form>
</template>
<script>
import http from '../services';
import emitter from '../eventHub';
import { startCase } from 'lodash';
export default {
  data() {
    return {
      link: '',
      resp: undefined,
      error: undefined,
      adding: false,
    };
  },
  methods: {
    startCase,
    async onSubmit() {
      this.adding = true;
      this.resp = undefined;
      try {
        const { data } = await http.post('/auth/url', { link: this.link });
        this.resp = data;
        emitter.$emit('new-url-added');
        this.link = '';
      } catch (err) {
        const { message } = http.getErrorMessageFromResponse(err);
        this.error = message;
      }
      this.adding = false;
    },
    removeError() {
      this.error = '';
      this.link = '';
    },
  },
};
</script>
<style scoped lang="scss">
@import '../scss/_variables.scss';
.url-input-form {
  width: 40%;
  display: flex;
  flex-direction: column;
  margin-top: 1.5rem;
  input {
    width: 70%;
    padding: 0.5rem;
    border: 1px solid $app-neutral;
    &:focus {
      outline: none;
      border: 1px solid $nav-active-link;
    }
  }
  .error {
    span {
      border-radius: $border-radius;
      padding: 0.2rem;
      display: inline-block;
      text-align: center;
      &:nth-child(1) {
        width: 69%;
        color: $app-error;
        margin-top: 1rem;
        background: $app-warning;
      }
      &:nth-child(2) {
        background: $app-neutral;
        width: 3%;
        color: white;
        margin-left: 0.1rem;
        &:hover {
          cursor: pointer;
        }
      }
    }
  }
  .success {
    display: flex;
    flex-direction: column;
    margin-top: 1rem;
    span {
      margin: 0.2rem 0;
    }
  }
  button {
    margin-top: 1rem;
    width: 73.5%;
    padding: 0.5rem;
    background: $app-neutral;
    color: white;
    border: none;
    border-radius: $border-radius;
    transition: opacity $transition-time;
    &:hover {
      cursor: pointer;
      opacity: 0.9;
    }
    &:focus {
      outline: none;
    }
  }
}
</style>
