<template>
<v-card
  class="pa-2"
  outlined
  tile
>
  <gantt-chart />
</v-card>
</template>

<script>
import GanttChart from "~/components/GanttChart.vue";

export default {
  components: {
    GanttChart,
  },
  data() {
    return {
      tasks: [
        {
          id: 'task_a',
          start: '2019-08-03',
          end: '2019-08-09',
          name: '',
          progress: 0
        },
        {
          start: '2019-08-05',
          end: '2019-08-23',
          name: '',
          progress: 0,
          dependencies: 'task_a',
        }
      ]
    }
  },
}
</script>