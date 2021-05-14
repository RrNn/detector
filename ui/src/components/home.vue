<template>
  <div class="">
    <ul class="urls">
      <li v-for="url in urls" :key="url.Link">
        <router-link :to="{ name: 'url', params: { id: url.ID } }">{{ url.Link }}</router-link>
        <!-- <a :href="url.Link" target="_blank">{{ url.Link }}</a> -->
      </li>
    </ul>
  </div>
</template>

<script>
  import http from '../services';
  import moment from 'moment';
  export default {
    data() {
      return {
        urls: [],
        error: ''
      };
    },
    methods: {
      moment
    },
    created: async function() {
      console.log(this.$router)
      try {
        const { data } = await http.get('/auth/urls?limit=10&offset=0');
        console.log('data', data);
        this.urls = [...this.urls, ...data]
      } catch (error) {
        this.error = http.getErrorMessageFromResponse(error).message;
      }
    },
  }
</script>

<style scoped lang="scss">
@import "../scss/_variables.scss";
.urls {
  padding: 0;
  display: flex;
  flex-direction: column;
  li {
    list-style-type: none;
    a {
      display: inline-block;
      width: 50%;
      text-decoration: none;
      background: $li-bg-color;
      margin: 0.2rem 0.85rem;
      padding: 0.5rem;
      border-radius: 0.2rem;
    }
  }
}
</style>
