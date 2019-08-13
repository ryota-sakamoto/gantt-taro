<template>
<v-card
    class="pa-2"
    outlined
    tile
>
  <chart-options @change-mode="d => mode = d" />
  <gantt-chart :mode="mode" @open-task-edit-modal="open_modal" />
  <task-edit-modal :dialog="is_show_modal" :task="task" @close-modal="is_show_modal = false" />
</v-card>
</template>

<script>
import GanttChart from "~/components/GanttChart.vue";
import ChartOptions from "~/components/chart/ChartOptions.vue";
import TaskEditModal from "~/components/chart/TaskEditModal.vue";

export default {
  components: {
    GanttChart,
    ChartOptions,
    TaskEditModal
  },
  data() {
    return {
      mode: "Day",
      is_show_modal: false,
      task: {}
    }
  },
  methods: {
    open_modal(task) {
      this.is_show_modal = true
      this.task = task
    }
  }
}
</script>