import { Module, VuexModule, Mutation, Action } from 'vuex-module-decorators'
import { MESSAGE } from '~/constants/exception'
import { ApiError } from '~/types/exception'
import { Snackbar } from '~/types/props/snackbar'
import { CommonState, InitialState, PromiseState } from '~/types/store'

const initialState: CommonState = {
  initialState: InitialState.PENDING,
  promiseState: PromiseState.NONE,
  snackbarColor: 'info',
  snackbarMessage: '',
}

@Module({
  name: 'common',
  stateFactory: true,
  namespaced: true,
})
export default class CommonModule extends VuexModule {
  private initialState: CommonState['initialState'] = initialState.initialState
  private promiseState: CommonState['promiseState'] = initialState.promiseState
  private snackbarColor: CommonState['snackbarColor'] = initialState.snackbarColor
  private snackbarMessage: CommonState['snackbarMessage'] = initialState.snackbarMessage

  public get getInitialState(): InitialState {
    return this.initialState
  }

  public get getPromiseState(): PromiseState {
    return this.promiseState
  }

  public get getSnackbarColor(): string {
    return this.snackbarColor
  }

  public get getSnackbarMessage(): string {
    return this.snackbarMessage
  }

  @Mutation
  private setInitialState(initialState: InitialState): void {
    this.initialState = initialState
  }

  @Mutation
  private setPromiseState(promiseState: PromiseState): void {
    this.promiseState = promiseState
  }

  @Mutation
  private setSnackbarColor(color: string): void {
    this.snackbarColor = color
  }

  @Mutation
  private setSnackbarMessage(message: string): void {
    this.snackbarMessage = message
  }

  @Action({})
  public factory(): void {
    this.setInitialState(initialState.initialState)
    this.setPromiseState(initialState.promiseState)
    this.setSnackbarColor(initialState.snackbarColor)
    this.setSnackbarMessage(initialState.snackbarMessage)
  }

  @Action({ rawError: true })
  public startInitialize(): void {
    this.setInitialState(InitialState.PROCESSING)
  }

  @Action({ rawError: true })
  public endInitialize(): void {
    this.setInitialState(InitialState.FINISHED)
  }

  @Action({ rawError: true })
  public startConnection(): void {
    this.setPromiseState(PromiseState.LOADING)
  }

  @Action({ rawError: true })
  public endConnection(): void {
    this.setPromiseState(PromiseState.NONE)
  }

  @Action({ rawError: true })
  public showSnackbar(payload: Snackbar): void {
    this.setSnackbarColor(payload.color)
    this.setSnackbarMessage(payload.message)
  }

  @Action({ rawError: true })
  public showSuccessInSnackbar(message: string): void {
    this.setSnackbarColor('success')
    this.setSnackbarMessage(message)
  }

  @Action({ rawError: true })
  public showErrorInSnackbar(err: Error): void {
    this.setSnackbarColor('error')

    if (err instanceof ApiError) {
      this.setSnackbarMessage(getApiErrorMessage(err))
    } else {
      this.setSnackbarMessage(MESSAGE.UNEXPEXTED_ERROR)
    }
  }

  @Action({ rawError: true })
  public hiddenSnackbar(): void {
    this.setSnackbarColor(initialState.snackbarColor)
    this.setSnackbarMessage(initialState.snackbarMessage)
  }
}

function getApiErrorMessage(err: ApiError): string {
  switch (err.status) {
    case 400:
      return MESSAGE.BAD_REQUEST
    case 401:
      return MESSAGE.UNAUTHORIZED
    case 403:
      return MESSAGE.FORBIDDEN
    case 404:
      return MESSAGE.PROCESS_FAILED
    case 409:
      return MESSAGE.CONFLICT
    case 500:
    case 501:
    case 503:
      return MESSAGE.SERVER_ERROR
    case 504:
      return MESSAGE.TIMEOUT
    default:
      return MESSAGE.UNEXPEXTED_ERROR
  }
}
