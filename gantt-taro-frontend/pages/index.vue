<template>
<div>
  <bar-item v-for="i in 3" :key="i" style="margin-bottom: 10px" />
</div>
</template>

<script>
import BarItem from "~/components/BarItem.vue";
import axios from "axios";

export default {
  components: {
    BarItem,
  },
  async asyncData() {
    try {
      await axios.get("/api/user/1")
    } catch (e) {
      console.log(e.response.data.message)
    }
  }
}
</script>
