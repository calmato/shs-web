import { Module, VuexModule, Mutation, Action } from 'vuex-module-decorators'
import { CommonState, PromiseState } from '~/types/store'

const initialState: CommonState = {
  promiseState: PromiseState.NONE,
}

@Module({
  name: 'common',
  stateFactory: true,
  namespaced: true,
})
export default class CommonModule extends VuexModule {
  private promiseState: CommonState['promiseState'] = initialState.promiseState

  public get getPromiseState(): PromiseState {
    return this.promiseState
  }

  @Mutation
  private setPromiseState(promiseState: PromiseState): void {
    this.promiseState = promiseState
  }

  @Action({})
  public factory(): void {
    this.setPromiseState(initialState.promiseState)
  }

  @Action({ rawError: true })
  public startConnection(): void {
    this.setPromiseState(PromiseState.LOADING)
  }

  @Action({ rawError: true })
  public endConnection(): void {
    this.setPromiseState(PromiseState.NONE)
  }
}
