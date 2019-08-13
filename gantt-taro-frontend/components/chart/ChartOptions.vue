<template>
<v-row align="center">
    <v-col class="d-flex" cols="12" sm="3" offset="9">
        <v-select
            :items="items"
            label="Mode"
            @change="change_mode"
        ></v-select>
    </v-col>
</v-row>
</template>

<script>
export default {
    data() {
        return {
            items: [
                "Day",
                "Week",
                "Month",
                "Year"
            ]
        }
    },
    methods: {
        change_mode(item) {
            this.$emit('change-mode', item)
        }
    }
}
</script>