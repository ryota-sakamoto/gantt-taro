<template>
<div class="container">
  <bar-item v-for="i in 3" :key="i" style="margin-bottom: 10px" />
</div>
</template>

<script>
import BarItem from "~/components/BarItem.vue";

export default {
  components: {
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

