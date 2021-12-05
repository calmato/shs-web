import Vue from 'vue'
import Vuex, { Store } from 'vuex'
import { createLocalVue } from '@vue/test-utils'
import { CommonStore, UserStore, initialiseStores, LessonStore } from '~/store'
import CommonModule from '~/store/common'
import UserModule from '~/store/user'
import response from '~~/test/helpers/api-mock'
import LessonModule from '~/store/lesson'

const localVue: Vue.VueConstructor<Vue> = createLocalVue()
localVue.use(Vuex)

const store = new Store({
  modules: {
    common: CommonModule,
    lesson: LessonModule,
    user: UserModule,
  },
})

function setup(): void {
  initialiseStores(store)
}

function refresh(): void {
  CommonStore.factory()
  LessonStore.factory()
  UserStore.factory()
}

// Error を返したいときだけ false にする
let isSafetyMode: boolean = true

function setSafetyMode(mode: boolean): void {
  isSafetyMode = mode
}

jest.mock('~/plugins/axios', () => ({
  $axios: {
    // $get: (key: string) => (isSafetyMode ? Promise.resolve(response['get'][key]) : Promise.reject(response['error'])),
    $post: (key: string) => (isSafetyMode ? Promise.resolve(response['post'][key]) : Promise.reject(response['error'])),
    // $patch: (key: string) => (isSafetyMode ? Promise.resolve(response['patch'][key]) : Promise.reject(response['error'])),
    // $put: (key: string) => (isSafetyMode ? Promise.resolve(response['put'][key]) : Promise.reject(response['error'])),
    // $delete: (key: string) => (isSafetyMode ? Promise.resolve(response['delete'][key]) : Promise.reject(response['error'])),
  },
}))

export { localVue, setup, refresh, setSafetyMode }