<template>
<div>
    <button @click="click_login" v-if="!is_login">login</button>
    <button @click="click_logout" v-if="is_login">logout</button>
</div>
</template>

<script>
/* eslint no-console: 0 */

import AuthService from "~/auth/AuthService.js";

export default {
    data() {
        return {
            auth: undefined,
            is_login: false,
        }
    },
    async mounted() {
        this.auth = new AuthService()
        await this.auth.init()

        this.is_login = await this.auth.isAuthenticated()

        if (this.is_login) {
            console.log(await this.auth.getUser())
        }
    },
    methods: {
        click_login() {
            this.auth.login()
        },
        click_logout() {
            this.auth.logout()
        }
    }
}
</script>

<style>

</style>
