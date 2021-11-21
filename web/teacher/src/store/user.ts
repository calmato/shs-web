import { Module, VuexModule, Mutation, Action } from 'vuex-module-decorators'
import { $axios } from '~/plugins/axios'
import { HelloRequest, HelloResponse } from '~/types/api/v1'
import { UserState } from '~/types/store'

const initialState: UserState = {
  message: '',
}

@Module({
  name: 'user',
  stateFactory: true,
  namespaced: true,
})
export default class UserModule extends VuexModule {
  private message: UserState['message'] = initialState.message

  public get getMessage(): string {
    return this.message
  }

  @Mutation
  private setMessage(message: string): void {
    this.message = message
  }

  @Action({})
  public factory(): void {
    this.setMessage(initialState.message)
  }

  @Action({ rawError: true })
  public async hello(): Promise<void> {
    const req: HelloRequest = { name: 'test' }

    await $axios
      .$post('/v1/hello', req)
      .then((res: HelloResponse) => {
        this.setMessage(res.message)
      })
      .catch((err: Error) => {
        throw err
      })
  }
}
