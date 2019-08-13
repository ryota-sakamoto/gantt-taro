import AuthService from "~/auth/AuthService.js";

export const strict = false

export const state = () => ({
  auth: new AuthService(),
  is_login: false,
  user: null,
  token: ''
})

export const mutations = {
  login(state) {
    state.auth.login()
  },
  logout(state) {
    state.auth.logout()
  }
}

export const actions = {
  async nuxtClientInit({ state }) {
    await state.auth.init()
    state.is_login = await state.auth.isAuthenticated()

    if (state.is_login) {
      state.user = await state.auth.getUser()
      state.token = await state.auth.auth.getTokenSilently()
    }
  }
}

export const getters = {
  is_login: (state) => state.is_login,
  user: (state) => state.auth,
  token: (state) => state.token,
}