<template>
  <input
    :type="type"
    :placeholder="placeholder"
    :name="name"
    v-model="inputValue"
    @input="onInput"
  />
</template>
<script>
export default {
  props: {
    type: {
      type: String,
      required: true,
    },
    placeholder: {
      type: String,
      required: true,
    },
    name: {
      type: String,
    },
    value: {},
    onChange: {},
  },
  data() {
    return {
      inputValue: this.value,
    };
  },
  methods: {
    onInput: function (e) {
      this.onChange(e);
    },
  },
};
</script>
<style>
input {
  width: 97%;
  padding: 0.5rem;
  font-size: 1.2rem;
  margin-bottom: 1rem;
}
</style>
