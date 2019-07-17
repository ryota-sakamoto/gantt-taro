<template>
<div>
</div>
</template>

<script>
import AuthService from "~/auth/AuthService.js";

export default {
    data() {
        return {
            auth: undefined,
        }
    },
    async mounted() {
        this.auth = new AuthService()
        await this.auth.init()

        this.auth.callback()
    },
}
</script>

<style>

</style>
