import Vue from 'vue'
import Vuex, { Store } from 'vuex'
import { createLocalVue } from '@vue/test-utils'
import { AuthStore, CommonStore, initialiseStores } from '~/store'
import response from '~~/test/helpers/api-mock'
import CommonModule from '~/store/common'
import AuthModule from '~/store/auth'

const localVue: Vue.VueConstructor<Vue> = createLocalVue()
localVue.use(Vuex)

const store = new Store({
  modules: {
    auth: AuthModule,
    common: CommonModule,
  },
})

function setup(): void {
  initialiseStores(store)
}

function refresh(): void {
  AuthStore.factory()
  CommonStore.factory()
}

// Error を返したいときだけ false にする
let isSafetyMode: boolean = true

function setSafetyMode(mode: boolean): void {
  isSafetyMode = mode
}

jest.mock('~/plugins/axios', () => ({
  $axios: {
    $get: (key: string) => (isSafetyMode ? Promise.resolve(response['get'][key]) : Promise.reject(response['error'])),
    // $post: (key: string) => (isSafetyMode ? Promise.resolve(response['post'][key]) : Promise.reject(response['error'])),
    // $patch: (key: string) => isSafetyMode ? Promise.resolve(response['patch'][key]) : Promise.reject(response['error']),
    // $put: (key: string) => (isSafetyMode ? Promise.resolve(response['put'][key]) : Promise.reject(response['error'])),
    // $delete: (key: string) => isSafetyMode ? Promise.resolve(response['delete'][key]) : Promise.reject(response['error']),
  },
}))

export { localVue, setup, refresh, setSafetyMode }
