<template>
<v-card
    class="pa-2"
    outlined
    tile
>
  <chart-options @change-mode="d => mode = d" />
  <gantt-chart :mode="mode" />
</v-card>
</template>

<script>
import GanttChart from "~/components/GanttChart.vue";
import ChartOptions from "~/components/chart/ChartOptions.vue";

export default {
  components: {
    GanttChart,
    ChartOptions,
  },
  data() {
    return {
      mode: "Day",
    }
  },
}
</script>