/* eslint-disable import/no-mutable-exports */
import { Store } from 'vuex'
import { getModule } from 'vuex-module-decorators'
import CommonModule from '~/store/common'
import LessonModule from '~/store/lesson'
import UserModule from '~/store/user'

let CommonStore: CommonModule
let LessonStore: LessonModule
let UserStore: UserModule

function initialiseStores(store: Store<any>): void {
  CommonStore = getModule(CommonModule, store)
  LessonStore = getModule(LessonModule, store)
  UserStore = getModule(UserModule, store)
}

export { initialiseStores, CommonStore, LessonStore, UserStore }
