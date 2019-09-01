<template>
  <v-row justify="center">
    <v-dialog v-model="dialog" persistent max-width="600px">
      <v-card>
        <v-card-title>
          <span class="headline">Task Edit</span>
        </v-card-title>
        <v-card-text>
          <v-container>
            <v-row>
              <v-col cols="12" sm="12" md="12">
                <v-text-field
                  label="Task Name*"
                  required
                  v-model="name"
                ></v-text-field>
              </v-col>
              <v-col cols="12" sm="6">
                <task-edit-date-picker label="Start Date" :date="start" @update-date="v => start = v" />
              </v-col>
              <v-col cols="12" sm="6">
                <task-edit-date-picker label="End Date" :date="end" @update-date="v => end = v" />
              </v-col>
            </v-row>
          </v-container>
          <small>*indicates required field</small>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="blue darken-1" text @click="close_modal">Close</v-btn>
          <v-btn color="blue darken-1" text @click="update_task">Save</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-row>
</template>

<script>
import TaskEditDatePicker from "~/components/chart/TaskEditDatePicker.vue";

export default {
  components: {
    TaskEditDatePicker
  },
  props: {
    dialog: Boolean,
    task: Object,
  },
  data() {
    return {
      id: undefined,
      name: undefined,
      start: undefined,
      end: undefined,
    }
  },
  methods: {
    update_task() {
      this.$emit('update-task', {
        id: this.id,
        name: this.name,
        start: this.start,
        end: this.end,
      })
      this.close_modal()
    },
    close_modal() {
      this.$emit('close-modal')
    }
  },
  watch: {
    dialog() {
      this.id = this.task.id
      this.name = this.task.name
      this.start = this.task.start
      this.end = this.task.end
    }
  }
}
</script>