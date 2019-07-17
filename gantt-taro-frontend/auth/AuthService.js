/* global createAuth0Client */

export default class AuthService {
    constructor() {
        this.auth = null
    }

    async init() {
        this.auth = await createAuth0Client({
            domain: '',
            client_id: '',
            redirect_uri: `${window.location.origin}/callback`,
        });
    }

    login() {
        this.auth.loginWithRedirect()
    }

    callback() {
        this.auth.handleRedirectCallback()
    }

    logout() {
        this.auth.logout()
    }

    async isAuthenticated() {
        return await this.auth.isAuthenticated()
    }

    async getUser() {
        return await this.auth.getUser({
            returnTo: `${window.location.origin}`
        })
    }
}