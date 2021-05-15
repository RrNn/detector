<template>
  <div class="pagination">
    <button :disabled="currentPage === 1" :onclick="onPrevClick">Prev</button>
    <span
    :class="
    page === '...'
    ? 'ellipsis'
    : page === currentPage && page !== '...' ? 'active page-number'
    :'page-number'"
    v-for="page in pagesToDisplay"
    :onClick="onPageNumberClick ? () => onPageNumberClick(page) : () => {}"
    :key="page">
    {{ page }}
  </span>
  <button :disabled="currentPage === totalPages" :onclick="onNextClick">Next</button>
</div>
</template>
<script>
  export default {
    props: {
      total: {
        type: Number,
        required: true,
      },
      limit: {
        type: Number,
        required: true,
        default: function() { return 10; }
      },
      offset: {
        type: Number,
        required: true,
        default: function() { return 0; }
      },
      onNextClick: {
        type: Function,
        required: true,
      },
      onPrevClick: {
        type: Function,
        required: true,
      },
      onPageNumberClick: {
        type: Function,
      }
    },
    data() {
      return {};
    },
    computed: {
      currentPage: function() {
        return this.offset > this.limit ? Math.ceil(this.offset / this.limit) : 1;
      },
      totalPages: function() {
        return Math.floor(this.total / this.limit);
      },
      pagesToDisplay: function() {
        if (this.totalPages <= 6) {
          return Array.from({ length: this.totalPages }).map((_, i) => i + 1);
        }
        const leftPages = this.currentPage > this.totalPages - 3
        ?  [1, 2, 3] // [this.currentPage - 3, this.currentPage - 2, this.currentPage - 1]
        : [this.currentPage, this.currentPage + 1, this.currentPage + 2];
        return [...leftPages, '...', this.totalPages - 2, this.totalPages - 1, this.totalPages]
      }
    },
    methods: {},
  }
</script>
<style scoped lang="scss">
.pagination {
  display: flex;
  justify-content: center;
  align-items: center;
  margin: 2rem 0;
  .page-number {
    display: inline-block;
    margin: 0 1rem;
    &:hover {
      cursor: pointer;
    }
    &.active {
      width: 2.5rem;
      display: flex;
      justify-content: center;
      border-radius: 0.2rem;
      border: 1px solid grey;
      &:hover {
        cursor: not-allowed;
      }
    }
  }
}
</style>
