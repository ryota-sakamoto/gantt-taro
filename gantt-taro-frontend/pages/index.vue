<template>
<div>
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
