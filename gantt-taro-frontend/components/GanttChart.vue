<template>
<div id="gantt">
</div>
</template>

<script>
import axios from "axios";
import Gantt from "frappe-gantt";

export default {
  props: {
    mode: String,
    update_task: Object,
  },
  components: {
  },
  data() {
    return {
      gantt: null,
      tasks: [],
    }
  },
  async mounted() {
    await this.get_all_tasks()
    this.create()

    this.scroll_today()
  },
  methods: {
    async _update_task() {  
      const id = Number(this.update_task.id.replace("task_", ""))

      await axios.post(`/api/v1/tasks/${id}`, {
        name: this.update_task.name,
        started_at: this.update_task.start,
        ended_at: this.update_task.end,
      },
      {
        headers: {
          "Authorization": `${this.$store.getters.token}`
        },
      })
    },
    async get_all_tasks() {
      const res = await axios.get("/api/v1/tasks", {
        headers: {
          "Authorization": `${this.$store.getters.token}`
        }
      })

      this.tasks = res.data.map(v => {
        return {
          id: `task_${v.id}`,
          start: v.started_at,
          end: v.ended_at,
          name: v.name,
          progress: 0,
        }
      })
    },
    create() {
      const f = this.tasks.every(o => o.id !== "task_new")
      if (f) {
        this.tasks.push({
          id: 'task_new',
          name: '',
          progress: 0,
        })
      }

      this.gantt = new Gantt("#gantt", this.tasks, {
        view_mode: this.mode,
        on_click: this.open_task_edit_modal,
      })

      this.gantt.bars[this.gantt.bars.length - 1].setup_click_event()
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
    },
    update_task() {
      this._update_task()
    }
  }
}
</script>

<style>
.popup-wrapper {
  display: none;
}
</style>