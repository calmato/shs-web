import { getAuth, onAuthStateChanged, signInWithEmailAndPassword } from 'firebase/auth'
import { Module, VuexModule, Mutation, Action } from 'vuex-module-decorators'
import { AxiosError } from 'axios'
import { $axios } from '~/plugins/axios'
import { app } from '~/plugins/firebase'
import { ErrorResponse } from '~/types/api/exception'
import { AuthResponse } from '~/types/api/v1'
import { ApiError } from '~/types/exception'
import { SignInForm, SubjectUpdateForm } from '~/types/form'
import { Auth, AuthState, Role } from '~/types/store'
import { authResponse2Auth } from '~/lib'

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
    role: Role.TEACHER,
    subjects: {
      小学校: [],
      中学校: [],
      高校: [],
    },
  },
}

@Module({
  name: 'auth',
  stateFactory: true,
  namespaced: true,
})
export default class AuthModule extends VuexModule {
  private uid: AuthState['uid'] = initialState.uid
  private token: AuthState['token'] = initialState.token
  private emailVerified: AuthState['emailVerified'] = initialState.emailVerified
  private auth: AuthState['auth'] = initialState.auth

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

  public get getRole(): Role {
    return this.auth?.role || Role.TEACHER
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
        this.setApiAuth(authResponse2Auth(res))
      })
      .catch((err: AxiosError) => {
        const res: ErrorResponse = { ...err.response?.data }
        throw new ApiError(res.status, res.message, res)
      })
  }

  @Action({ rawError: true })
  public async updateOwnSubjects(formData: SubjectUpdateForm): Promise<void> {
    try {
      await $axios.$patch('/v1/me/subjects', formData)
    } catch (err) {
      if ($axios.isAxiosError(err)) {
        const res: ErrorResponse = { ...err.response?.data }
        throw new ApiError(res.status, res.message, res)
      }
      throw new Error('internal server error')
    }
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
