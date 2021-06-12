<template>
  <div class="not-found">
    <h1>Page Not Found</h1>
  </div>
</template>
<style scoped lang="scss">
@import "../scss/_variables.scss";
.not-found {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 80vh;
  color: $app-neutral;
  font-family: sans-serif;
  h1 {
    font-weight: 100;
    font-size: 5rem;
  }
}
</style>
