<template>
  <div v-if="!pings?.length" class="no-pings">
    No Pings to show. Check back later
  </div>
  <table class="pings-table" :class="classname" v-if="pings?.length">
    <thead>
      <tr>
        <th>Ping Time</th>
        <th>Status</th>
        <th>Error</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="ping in pings" :key="ping.ID">
        <td>{{ moment(ping.Time).fromNow() }}</td>
        <td>{{ ping.Status || '-' }}</td>
        <td>{{ ping.Error || '-' }}</td>
      </tr>
    </tbody>
  </table>
</template>
<script>
import { isEmpty } from 'lodash';
import moment from 'moment';
export default {
  props: ['pings', 'classname'],
  data() {
    return {};
  },
  methods: {
    isEmpty,
    moment,
  },
  created: function () {},
  updated: function () {},
};
</script>
<style scoped lang="scss">
@import '../scss/_variables.scss';
.no-pings {
  height: 75vh;
  font-size: 2rem;
  display: flex;
  justify-content: center;
  align-items: center;
  font-weight: 100;
}
.pings-table {
  border-spacing: 0 0.0625rem;
  width: 100%;
  margin-bottom: 2rem;
  thead,
  tbody > tr > td,
  th {
    text-align: left;
    padding: 0.5rem;
    font-size: 1.2rem;
  }
  thead {
    tr {
      background: $nav-bg;
      color: $white;
    }
  }
  tbody {
    tr {
      background-color: #a3caec;
      &:nth-child(odd) {
        background-color: #86c1f5;
      }
      &:hover {
        background-color: #72baf9;
      }
    }
  }
}
</style>
