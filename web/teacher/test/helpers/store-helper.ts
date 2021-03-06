import Vue from 'vue'
import Vuex, { Store } from 'vuex'
import { createLocalVue } from '@vue/test-utils'
import {
  CommonStore,
  UserStore,
  initialiseStores,
  LessonStore,
  AuthStore,
  ShiftStore,
  SubmissionStore,
  ClassroomStore,
} from '~/store'
import AuthModule from '~/store/auth'
import CommonModule from '~/store/common'
import LessonModule from '~/store/lesson'
import ShiftModule from '~/store/shift'
import UserModule from '~/store/user'
import SubmissionModule from '~/store/submission'
import ClassroomModule from '~/store/classroom'
import response from '~~/test/helpers/api-mock'

const localVue: Vue.VueConstructor<Vue> = createLocalVue()
localVue.use(Vuex)

const store = new Store({
  modules: {
    auth: AuthModule,
    common: CommonModule,
    lesson: LessonModule,
    shift: ShiftModule,
    submission: SubmissionModule,
    user: UserModule,
    classroom: ClassroomModule,
  },
})

function setup(): void {
  initialiseStores(store)
}

function refresh(): void {
  AuthStore.factory()
  CommonStore.factory()
  LessonStore.factory()
  ShiftStore.factory()
  SubmissionStore.factory()
  UserStore.factory()
  ClassroomStore.factory()
}

// Error を返したいときだけ false にする
let isSafetyMode: boolean = true

function setSafetyMode(mode: boolean): void {
  isSafetyMode = mode
}

let isAxiosErrorMockValue: boolean = true

function setIsAxiosMockValue(val: boolean): void {
  isAxiosErrorMockValue = val
}

jest.mock('axios', () => ({
  isAxiosError: () => isAxiosErrorMockValue,
}))

jest.mock('~/plugins/axios', () => ({
  $axios: {
    $get: (key: string) => (isSafetyMode ? Promise.resolve(response['get'][key]) : Promise.reject(response['error'])),
    $post: (key: string) => (isSafetyMode ? Promise.resolve(response['post'][key]) : Promise.reject(response['error'])),
    $patch: (key: string) =>
      isSafetyMode ? Promise.resolve(response['patch'][key]) : Promise.reject(response['error']),
    // $put: (key: string) => (isSafetyMode ? Promise.resolve(response['put'][key]) : Promise.reject(response['error'])),
    $delete: (key: string) =>
      isSafetyMode ? Promise.resolve(response['delete'][key]) : Promise.reject(response['error']),
  },
}))

export { localVue, setup, refresh, setSafetyMode, setIsAxiosMockValue }
