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
                  v-model="task.name"
                ></v-text-field>
              </v-col>
              <v-col cols="12" sm="6">
                <task-edit-date-picker label="Start Date" :date="task.start" />
              </v-col>
              <v-col cols="12" sm="6">
                <task-edit-date-picker label="End Date" :date="task.end" />
              </v-col>
            </v-row>
          </v-container>
          <small>*indicates required field</small>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="blue darken-1" text @click="close_modal">Close</v-btn>
          <v-btn color="blue darken-1" text @click="close_modal">Save</v-btn>
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
  methods: {
    close_modal() {
      this.$emit('close-modal')
    }
  }
}
</script>