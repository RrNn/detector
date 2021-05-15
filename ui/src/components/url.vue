<template>
  <div class="url-details">
    <section class="header">
      <h1>
        <a :href="url.Link" target="_blank">{{ url.Link }}</a>
        <span v-if="totalPings">Checked ({{ totalPings }}) times</span>
      </h1>
    </section>
    <section class="details">
      <pings-table v-if="!isEmpty(url)" :pings="url.Pings" />
      <span v-else>Loading...</span>
      <pagination
        :total="totalPings"
        :limit="limit"
        :offset="offset"
        :onNextClick="onNextClick"
        :onPrevClick="onPrevClick"
        :onPageNumberClick="onPageNumberClick"
      />
    </section>
  </div>
</template>

<script>
  import http from '../services';
  import PingsTable from './pingsTable';
  import Pagination from './pagination';
  import { merge, isEmpty } from 'lodash';
  export default {
    components: { PingsTable, Pagination },
    data() {
      return {
        url: {},
        totalPings: 0, // undefined maybe??
        limit: 15,
        offset: 0,
      }
    },
    methods: {
      isEmpty,
      onNextClick: function() {
        this.offset = this.offset + this.limit;
      },
      onPrevClick: function() {
        this.offset = this.offset - this.limit;
      },
      onPageNumberClick: function(page) {
        this.offset = this.limit * page;
      }
    },
    mounted: async function() {
      const { id } = this.$route.params;
      const { data: url } = await http.get(`/auth/urls/${id}?limit=${this.limit}&offset=${this.offset}&withpings=true`);
      this.url = merge(this.url, url.URL);
      this.totalPings = url.TotalPings;
    },
    updated: async function() {
      const { id } = this.$route.params;
      const { data: url } = await http.get(`/auth/urls/${id}?limit=${this.limit}&offset=${this.offset}&withpings=true`);
      this.url = merge(this.url, url.URL);
      this.totalPings = url.TotalPings;
    },
    computed: {
      id() { return this.$route.params.id; }
    }
  }
</script>
<style scoped lang="scss">
  .url-details {
    display: flex;
    justify-content: center;
    flex-direction: column;
    align-items: center;
    .header {
      h1 {
        padding: 0.5rem;
        span { margin-left: 0.5rem; }
      }
    }
    .details {
      width: 90%;
    }
  }
</style>
