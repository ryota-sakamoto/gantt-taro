<template>
<div class="bar"
    :style="{ 'margin-left': margin_left + 'px' }"
    @mousemove="mouse_move"
    >
        <div class="bar-front mouse-left"
            @mousedown="mouse_down($event, 'left')"
            @mouseup="mouse_up($event, 'left')"
        ></div>
        <div class="bar-middle mouse-middle"
            :style="{ 'width': width + 'px' }"
            @mousedown="mouse_down($event, 'middle')"
            @mouseup="mouse_up($event, 'middle')"
        ></div>
        <div class="bar-end mouse-right"
            @mousedown="mouse_down($event, 'right')"
            @mouseup="mouse_up($event, 'right')"
        ></div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            can_move: false,
            start_left: 0,
            start_right: 0,
            current_target: null,

            width: 30,
            margin_left: 0,
        }
    },
    methods: {
        mouse_down(e, v) {
            this.can_move = true
            this.start_right = e.clientX
            this.start_left = e.clientX
            this.current_target = v
        },
        mouse_up() {
            this.can_move = false
        },
        mouse_move(e) {
            if (this.can_move) {
                if (this.current_target === "right") {
                    const len = e.clientX - this.start_right
                    this.width += len
                    this.start_right = e.clientX
                } else if (this.current_target === "left") {
                    const len = e.clientX - this.start_left
                    this.width -= len
                    this.margin_left += len
                    this.start_left = e.clientX
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
.bar-front, .bar-end {
    height: 100%;
    width: 10px;
    background-color: grey;
    display: inline-block;
}
.bar-middle {
    height: 100%;
    background-color: steelblue;
    display: inline-block;
}

.mouse-right {
    cursor: e-resize	;
}
.mouse-middle {
    cursor: move;
}
.mouse-left {
    cursor: w-resize	;
}
</style>
