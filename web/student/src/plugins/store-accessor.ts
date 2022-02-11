/* eslint-disable import/no-mutable-exports */
import { Store } from 'vuex'
import { getModule } from 'vuex-module-decorators'
import AuthModule from '~/store/auth'
import CommonModule from '~/store/common'
import SubmissionModule from '~/store/submission'

let AuthStore: AuthModule
let CommonStore: CommonModule
let SubmissionStore: SubmissionModule

function initialiseStores(store: Store<any>): void {
  AuthStore = getModule(AuthModule, store)
  CommonStore = getModule(CommonModule, store)
  SubmissionStore = getModule(SubmissionModule, store)
}

export { initialiseStores, AuthStore, CommonStore, SubmissionStore }
