<template>
  <div class="url-details">
    <section class="header">
      <h1><a :href="url.Link" target="_blank">{{ url.Link }}</a></h1>
    </section>
    <section class="details">
      <pings-table v-if="!isEmpty(url)" :pings="url.Pings" />
      <span v-else>Loading...</span>
    </section>
  </div>
</template>

<script>
  import http from '../services';
  import PingsTable from './pingsTable';
  import { merge, isEmpty } from 'lodash';
  export default {
    components: { PingsTable },
    data() {
      return {
        url: {},
      }
    },
    methods: { isEmpty },
    mounted: async function(){
      const { id } = this.$route.params;
      const { data: url } = await http.get(`/auth/urls/${id}?limit=10&offset=10&withpings=true`);
      this.url = merge(this.url, url);
    },
    computed: {
      id() {
        return this.$route.params.id;
      }
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
      }
    }
    .details {
      width: 90%;
    }
  }
</style>
