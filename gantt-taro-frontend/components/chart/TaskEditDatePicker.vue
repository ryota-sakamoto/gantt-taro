<template>
<v-menu
  ref="menu"
  v-model="menu"
  :close-on-content-click="false"
  :return-value.sync="local_date"
  transition="scale-transition"
  offset-y
  full-width
  min-width="290px"
>
  <template v-slot:activator="{ on }">
    <v-text-field
      v-model="local_date"
      :label="label"
      readonly
      v-on="on"
    ></v-text-field>
  </template>
  <v-date-picker v-model="local_date" no-title scrollable>
    <v-spacer></v-spacer>
    <v-btn text color="primary" @click="cancel">Cancel</v-btn>
    <v-btn text color="primary" @click="save_date">OK</v-btn>
  </v-date-picker>
</v-menu>
</template>

<script>
export default {
  props: {
    label: String,
    date: String,
  },
  data() {
    return {
      menu: false,
      local_date: this.date,
    }
  },
  methods: {
    cancel() {
      this.local_date = this.date
      this.menu = false
    },
    save_date() {
      this.$refs.menu.save(this.local_date)
      this.$emit('update-date', this.local_date)
      this.menu = false
    }
  }
}
</script>