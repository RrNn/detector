<template>
  <div class="url-details">
    <section class="header">
      <h1>
        <a :href="url.Link" target="_blank">{{ url.Link }}</a>
        <span v-if="totalPings">Checked ({{ totalPings }}) times</span>
      </h1>
      <button class="delete-btn" :onClick="deleteUrl">Delete URL</button>
    </section>
    <section class="details">
      <div v-if="!isEmpty(url)">
        <pings-table classname="no-margin-bottom" :pings="url.Pings" />
      </div>
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
import { isEmpty } from 'lodash';
export default {
  components: { PingsTable, Pagination },
  data() {
    return {
      url: {},
      totalPings: 0, // undefined maybe??
      limit: 10,
      offset: 0,
    };
  },
  methods: {
    isEmpty,
    onNextClick: function () {
      this.offset = this.offset + this.limit;
    },
    onPrevClick: function () {
      this.offset = this.offset - this.limit;
    },
    onPageNumberClick: function (page) {
      this.offset = this.limit * (page - 1);
    },
    async fetch() {
      const { id } = this.$route.params;
      const { data: url } = await http.get(
        `/auth/urls/${id}?limit=${this.limit}&offset=${this.offset}&withpings=true`
      );
      this.url = url.URL;
      this.totalPings = url.TotalPings;
    },
    async deleteUrl() {
      const { id } = this.$route.params;
      try {
        const {
          data: { message },
        } = await http.delete(`/auth/urls/${id}`);
        alert(message);
        this.$router.push({ path: '/' });
      } catch (err) {
        const { message } = http.getErrorMessageFromResponse(err);
        alert(message);
      }
    },
  },
  mounted: async function () {
    this.fetch();
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
  computed: {
    id() {
      return this.$route.params.id;
    },
  },
};
</script>
<style lang="scss">
@import '../scss/_variables.scss';
.url-details {
  display: flex;
  justify-content: center;
  flex-direction: column;
  align-items: center;
  .header {
    width: 90%;
    display: flex;
    justify-content: space-between;
    align-items: center;
    h1 {
      width: 85%;
      padding: 0.5rem;
      display: inline-flex;
      flex-wrap: wrap;
      span {
        margin-left: 0.5rem;
        border: 1px solid $app-neutral;
        border-radius: $border-radius;
        padding: 0 1rem;
      }
    }
    .delete-btn {
      background: $app-warning;
      border: none;
      padding: 0.5rem 1rem;
      font-size: 1rem;
      border-radius: $border-radius;
      &:hover {
        cursor: pointer;
      }
    }
  }
  .details {
    width: 90%;
    div {
      max-height: 70vh;
      overflow: scroll;
      .no-margin-bottom {
        margin-bottom: 0;
      }
    }
  }
}
</style>
