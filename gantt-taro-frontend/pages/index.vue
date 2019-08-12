<template>
<div>
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
