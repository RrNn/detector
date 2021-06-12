<template>
  <div class="pagination">
    <button :disabled="currentPage === 1" :onclick="onPrevClick">Prev</button>
    <span
      v-for="page in pagesToDisplay"
      :class="
        page === '...'
          ? 'ellipsis'
          : page === currentPage && page !== '...'
          ? 'active page-number'
          : 'page-number'
      "
      :onClick="onPageNumberClick ? () => onPageNumberClick(page) : () => {}"
      :key="page"
    >
      {{ page }}
    </span>
    <button :disabled="currentPage === totalPages" :onclick="onNextClick">
      Next
    </button>
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
      default: function () {
        return 10;
      },
    },
    offset: {
      type: Number,
      required: true,
      default: function () {
        return 0;
      },
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
    },
  },
  data() {
    return {};
  },
  computed: {
    currentPage: function () {
      return this.offset > 0 ? Math.floor(this.offset / this.limit) + 1 : 1;
    },
    totalPages: function () {
      return Math.ceil(this.total / this.limit);
    },
    pagesToDisplay: function () {
      let current =
          this.offset !== 0 ? (this.offset + this.limit) / this.limit : 1,
        last = Math.ceil(this.total / this.limit),
        delta = 2,
        left = current - delta,
        right = current + delta + 1,
        range = [],
        rangeWithDots = [],
        l;
      for (let i = 1; i <= last; i++) {
        if (i == 1 || i == last || (i >= left && i < right)) {
          range.push(i);
        }
      }
      for (let i of range) {
        if (l) {
          if (i - l === 2) {
            rangeWithDots.push(l + 1);
          } else if (i - l !== 1) {
            rangeWithDots.push('...');
          }
        }
        rangeWithDots.push(i);
        l = i;
      }
      return rangeWithDots;
    },
  },
  methods: {},
};
</script>
<style scoped lang="scss">
.pagination {
  display: flex;
  justify-content: center;
  align-items: center;
  margin: 2rem 0;
  position: fixed;
  bottom: 1rem;
  right: 0;
  left: 0;
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
