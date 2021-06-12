<template>
  <div class="">
    <div class="error" v-if="error">{{ error }}</div>
    <div class="list-and-form">
      <ul class="urls">
        <li v-for="url in urls" :key="url.Link">
          <router-link :to="{ name: 'url', params: { id: url.ID } }">
            {{ url.Link }}
          </router-link>
          <!-- <a :href="url.Link" target="_blank">{{ url.Link }}</a> -->
        </li>
      </ul>
      <add-url-form></add-url-form>
    </div>
    <pagination
      :total="totalUrls"
      :limit="limit"
      :offset="offset"
      :onNextClick="onNextClick"
      :onPrevClick="onPrevClick"
      :onPageNumberClick="onPageNumberClick"
    />
  </div>
</template>

<script>
import moment from 'moment';
import http from '../services';
import emitter from '../eventHub';
import Pagination from './pagination';
import AddUrlForm from './addurlform';

export default {
  components: { Pagination, AddUrlForm },
  data() {
    return {
      urls: [],
      error: '',
      totalUrls: 0,
      limit: 15,
      offset: 0,
    };
  },
  methods: {
    moment,
    onNextClick() {
      this.offset = this.offset + this.limit;
    },
    onPrevClick() {
      this.offset = this.offset - this.limit;
    },
    onPageNumberClick(page) {
      this.offset = this.limit * (page - 1);
    },
    async fetch() {
      try {
        const { data } = await http.get(
          `/auth/urls?limit=${this.limit}&offset=${this.offset}`
        );
        this.urls = data.URLS;
        this.totalUrls = data.TotalUrls;
      } catch (error) {
        this.error = http.getErrorMessageFromResponse(error).message;
      }
    },
  },
  mounted() {
    this.fetch();
  },
  created() {
    emitter.$on('new-url-added', () => this.fetch(), this);
  },
  watch: {
    offset(value, previous) {
      if (value !== previous) {
        return this.fetch();
      }
    },
    limit(value, previous) {
      if (value !== previous) {
        return this.fetch();
      }
    },
  },
};
</script>

<style scoped lang="scss">
@import '../scss/_variables.scss';
.error {
  background: $app-warning;
  color: $app-neutral;
  width: 90%;
  padding: 0.5rem;
  margin: 1rem auto;
  display: flex;
  justify-content: center;
  font-size: 2rem;
  border-radius: 0.5rem;
  text-transform: capitalize;
}
.list-and-form {
  display: flex;
  justify-content: space-between;
  margin-top: 1rem;
  .urls {
    display: flex;
    flex-direction: column;
    width: 50%;
    li {
      list-style-type: none;
      a {
        display: inline-block;
        width: 90%;
        text-decoration: none;
        background: $li-bg-color;
        margin: 0.2rem 0.85rem;
        padding: 0.5rem;
        border-radius: $border-radius;
      }
    }
  }
}
</style>
