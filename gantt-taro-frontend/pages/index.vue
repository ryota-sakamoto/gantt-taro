<template>
<v-card
    class="pa-2"
    outlined
    tile
>
  <chart-options @change-mode="d => mode = d" />
  <gantt-chart :mode="mode" @open-task-edit-modal="open_modal" :update_task="update_task" />
  <task-edit-modal :dialog="is_show_modal" :task="task" @update-task="update" @close-modal="is_show_modal = false" />
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
      task: {},
      update_task: undefined,
    }
  },
  methods: {
    open_modal(task) {
      this.task = task
      this.is_show_modal = true
    },
    update(task) {
      this.update_task = task
    }
  }
}
</script>