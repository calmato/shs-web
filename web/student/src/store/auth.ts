import { getAuth, onAuthStateChanged, signInWithEmailAndPassword } from 'firebase/auth'
import { Module, VuexModule, Mutation, Action } from 'vuex-module-decorators'
import { AxiosError } from 'axios'
import { $axios } from '~/plugins/axios'
import { app } from '~/plugins/firebase'
import { ErrorResponse } from '~/types/api/exception'
import { AuthResponse, UpdateMyMailRequest, UpdateMyPasswordRequest } from '~/types/api/v1'
import { ApiError } from '~/types/exception'
import { SignInForm, StudentUpdateMailForm, StudentUpdatePasswordForm } from '~/types/form'
import { Auth, AuthState, SchoolType } from '~/types/store'

const initialState: AuthState = {
  uid: '',
  token: '',
  emailVerified: false,
  auth: {
    id: '',
    lastName: '',
    firstName: '',
    lastNameKana: '',
    firstNameKana: '',
    mail: '',
    schoolType: SchoolType.UNKNOWN,
    grade: 0,
  },
}

@Module({
  name: 'auth',
  stateFactory: true,
  namespaced: true,
})
export default class AuthModule extends VuexModule {
  private uid: string = initialState.uid
  private token: string = initialState.token
  private emailVerified: boolean = initialState.emailVerified
  private auth: Auth = initialState.auth

  public get getUid(): string {
    return this.uid
  }

  public get getToken(): string {
    return this.token
  }

  public get getEmailVerified(): boolean {
    return this.emailVerified
  }

  public get getAuth(): Auth {
    return this.auth
  }

  @Mutation
  private setToken(token: string): void {
    this.token = token
  }

  @Mutation
  private setApiAuth(auth: Auth): void {
    this.auth = auth
  }

  @Mutation
  private setFirebaseAuth({ uid, emailVerified }: { uid: string; emailVerified: boolean }): void {
    this.uid = uid
    this.emailVerified = emailVerified
  }

  @Action({})
  public factory(): void {
    this.setToken(initialState.token)
    this.setApiAuth(initialState.auth)
    this.setFirebaseAuth({ ...initialState })
  }

  @Action({ rawError: true })
  public authentication(): Promise<void> {
    return new Promise((resolve: () => void, reject: (reason: Error) => void) => {
      const auth = getAuth(app)
      onAuthStateChanged(auth, async (user) => {
        if (user) {
          await this.getIdToken()
          this.getFirebaseAuth()
          resolve()
        } else {
          this.signOut()
          reject(new Error('unauthorized'))
        }
      })
    })
  }

  @Action({ rawError: true })
  public async signIn(form: SignInForm): Promise<void> {
    const auth = getAuth(app)
    await signInWithEmailAndPassword(auth, form.mail, form.password)
      .then(() => {
        this.authentication()
      })
      .catch((err) => {
        throw new ApiError(err.code, err.message, err)
      })
  }

  @Action({ rawError: true })
  public signOut(): void {
    const auth = getAuth(app)
    auth.signOut()
    this.factory()
  }

  @Action({ rawError: true })
  public async showAuth(): Promise<void> {
    await $axios
      .$get('/v1/me')
      .then((res: AuthResponse) => {
        this.setApiAuth({ ...res })
      })
      .catch((err: AxiosError) => {
        const res: ErrorResponse = { ...err.response?.data }
        throw new ApiError(res.status, res.message, res)
      })
  }

  @Action({ rawError: true })
  public async updatePassword({ form }: { form: StudentUpdatePasswordForm }): Promise<void> {
    const req: UpdateMyPasswordRequest = {
      password: form.params.password,
      passwordConfirmation: form.params.passwordConfirmaion,
    }

    await $axios.$patch(`/v1/me/password`, req).catch((err: AxiosError) => {
      const res: ErrorResponse = { ...err.response?.data }
      throw new ApiError(res.status, res.message, res)
    })
  }

  @Action({ rawError: true })
  public async updateMail({ form }: { form: StudentUpdateMailForm }): Promise<void> {
    const req: UpdateMyMailRequest = {
      mail: form.params.mail,
    }

    await $axios.$patch(`/v1/me/mail`, req).catch((err: AxiosError) => {
      const res: ErrorResponse = { ...err.response?.data }
      throw new ApiError(res.status, res.message, res)
    })
  }

  @Action({ rawError: true })
  private async getIdToken(): Promise<void> {
    const auth = getAuth(app)
    await auth.currentUser
      ?.getIdToken(true)
      .then((token: string): void => {
        this.setToken(token)
      })
      .catch((err: Error) => {
        throw err
      })
  }

  @Action({ rawError: true })
  private getFirebaseAuth(): void {
    const auth = getAuth(app)
    const user = auth.currentUser
    if (!user) {
      return
    }

    this.setFirebaseAuth({ ...user })
  }
}
