<template>
<div class="container">
  <auth />
  <bar-item v-for="i in 3" :key="i" style="margin-bottom: 10px" />
</div>
</template>

<script>
import Auth from "~/components/Auth.vue";
import BarItem from "~/components/BarItem.vue";

export default {
  components: {
    Auth,
    BarItem,
  }
}
</script>

<style>
.container {
  background-color: whitesmoke;
  height: 100%;
}
</style>

