<template>
<div id="gantt">
</div>
</template>

<script>
import Gantt from "frappe-gantt";

export default {
  props: {
    mode: String,
  },
  components: {
  },
  data() {
    return {
      gantt: null,
      tasks: [
        {
          id: 'task_a',
          start: '2019-07-31',
          end: '2019-08-09',
          name: 'task_a',
          progress: 0
        },
        {
          id: 'task_b',
          start: '2019-08-05',
          end: '2019-08-23',
          name: 'task_b',
          progress: 0,
          dependencies: 'task_a',
        },
        {
          id: 'task_c',
          start: '2019-08-11',
          end: '2019-08-15',
          name: 'task_c',
          progress: 0,
          dependencies: 'task_b',
        },
        {
          id: 'task_d',
          start: '2019-08-11',
          end: '2019-08-15',
          name: 'task_d',
          progress: 0,
          dependencies: 'task_a',
        },
        {
          id: 'task_e',
          start: '2019-08-15',
          end: '2019-08-18',
          name: 'こんにちはする',
          progress: 0,
          dependencies: 'task_c',
        }
      ]
    }
  },
  mounted() {
    this.create()

    console.log(this.gantt)

    this.scroll_today()
  },
  methods: {
    create() {
      this.gantt = new Gantt("#gantt", this.tasks, {
        view_mode: this.mode,
        on_click: this.open_task_edit_modal,
      })
    },
    scroll_today() {
      const oldest = this.gantt.get_oldest_starting_date().getTime()
      const t = new Date() - oldest

      this.gantt.gantt_start = new Date(this.gantt.gantt_start.getTime() - t)
      this.gantt.set_scroll_position()
    },
    open_task_edit_modal(task) {
      this.$emit('open-task-edit-modal', task)
    }
  },
  watch: {
    mode() {
      this.create()
    }
  }
}
</script>

<style>
.popup-wrapper {
  display: none;
}
</style>