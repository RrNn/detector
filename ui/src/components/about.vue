<template>
  <div class="about-the-system">
    <div>
      <p>
        Add links in your application to this system to monitor and be able to
        track when they are down and when they are up. If you want to track many
        pages in your site, add as many pages as you can. This system will show
        you a table of when each url was checked/pinged, ie, a get request will
        be sent to every page after a specific time interval and the links page
        will show you the time at which the link was checked/pinged, the status
        code and error if any happened
      </p>
      <p>Note that you cannot add more than one similar page.</p>
    </div>
  </div>
</template>

<script>
export default {
  data: function () {
    return {};
  },
};
</script>
<style scoped lang="scss">
@import "../scss/_variables.scss";
.about-the-system {
  display: flex;
  height: 80vh;
  justify-content: center;
  align-items: center;
  width: 50%;
  margin: auto;
  font-size: 1.2rem;
  font-family: sans-serif;
  color: $app-neutral;
  div {
    background: $lighter-neutral;
    padding: 0.5rem;
    text-align: justify;
    border-radius: 0.5rem;
  }
}
</style>
