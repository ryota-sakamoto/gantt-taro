<template>
<v-progress-linear
    height="30"
    :buffer-value="width"
    :value="width / 10"
    :style="{ 'margin-left': margin_left + 'px', cursor: cursor }"
    @mousemove="change_cursor"
    @mousedown="mouse_down"
    @mouseup="mouse_up"
>
</v-progress-linear>
</template>

<script>
export default {
    data() {
        return {
            can_move: false,
            start_left: 0,
            start_right: 0,
            current_target: null,

            width: 10,
            max_width_px: 0,
            margin_left: 0,
            cursor: "",
        }
    },
    methods: {
        is_target(e) {
            return e.target.className === "v-progress-linear__buffer"
        },
        mouse_down(e) {
            if (!this.is_target(e)) {
                return
            }
            this.can_move = true
            this.start_right = e.clientX
            this.start_left = e.clientX

            this.max_width_px = e.target.clientWidth / this.width * 100
        },
        mouse_up() {
            this.can_move = false
        },
        change_cursor(e) {
            if (this.can_move) {
                this.mouse_move(e)
                return
            }

            if (!this.is_target(e)) {
                this.cursor = ""
                return
            }

            if (e.target.clientWidth - e.offsetX <= 10) {
                this.cursor = "e-resize"
                this.current_target = "right"
            } else {
                this.cursor = "move"
                this.current_target = "middle"
            }
        },
        mouse_move(e) {
            if (this.can_move) {
                if (this.current_target === "right") {
                    this.width = (e.target.clientWidth + e.movementX) / this.max_width_px * 100
                } else if (this.current_target === "middle") {
                    const len = e.clientX - this.start_left
                    this.margin_left += len
                    this.start_left = e.clientX
                }
            }
        }
    }
}
</script>

<style>
.bar {
    height: 30px;
    font-size: 0;
}
</style>
