import { Role, SchoolType } from './common'
import { Subject } from './lesson'

export interface Auth {
  id: string
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
  mail: string
  role: Role
  subjects: {
    [key in SchoolType]: Subject[]
  }
}

export interface AuthState {
  uid: string
  token: string
  emailVerified: boolean
  auth: Auth
}
