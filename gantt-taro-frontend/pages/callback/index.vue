<template>
<div>
</div>
</template>

<script>
/* eslint no-console: 0 */

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

        const query = location.search
        if (query.includes("code=") && query.includes("state=")) {
            this.auth.callback()

            this.$router.push({ name: "index" })
        }
    },
}
</script>

<style>

</style>
