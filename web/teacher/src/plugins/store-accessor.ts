/* eslint-disable import/no-mutable-exports */
import { Store } from 'vuex'
import { getModule } from 'vuex-module-decorators'
import CommonModule from '~/store/common'
import UserModule from '~/store/user'

let CommonStore: CommonModule
let UserStore: UserModule

function initialiseStores(store: Store<any>): void {
  CommonStore = getModule(CommonModule, store)
  UserStore = getModule(UserModule, store)
}

export { initialiseStores, CommonStore, UserStore }
