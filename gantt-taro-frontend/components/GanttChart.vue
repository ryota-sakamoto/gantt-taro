<template>
<div id="gantt">
</div>
</template>

<script>
import Gantt from "frappe-gantt";

export default {
  data() {
    return {
      gantt: null,
      tasks: [
        {
          id: 'task_a',
          start: '2019-07-31',
          end: '2019-08-09',
          name: '',
          progress: 0
        },
        {
          id: 'task_b',
          start: '2019-08-05',
          end: '2019-08-23',
          name: '',
          progress: 0,
          dependencies: 'task_a',
        },
        {
          id: 'task_c',
          start: '2019-08-11',
          end: '2019-08-15',
          name: '',
          progress: 0,
          dependencies: 'task_b',
        }
      ]
    }
  },
  mounted() {
    this.gantt = new Gantt("#gantt", this.tasks)

    console.log(this.gantt)

    this.scroll_today()
  },
  methods: {
      scroll_today() {
        const oldest = this.gantt.get_oldest_starting_date().getTime()
        const t = new Date() - oldest

        this.gantt.gantt_start = new Date(this.gantt.gantt_start.getTime() - t)
        this.gantt.set_scroll_position()
      }
  }
}
</script>

<style>
.popup-wrapper {
  display: none;
}
</style>