import { Subject } from './lesson'
import { Role, SchoolType } from './common'

export interface Auth {
  id: string
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
  mail: string
  role: Role
  subjects: {
    [key in SchoolType.ELEMENTARY_SCHOOL | SchoolType.JUNIOR_HIGH_SCHOOL | SchoolType.HIGH_SCHOOL]: Subject[]
  }
}

export interface AuthState {
  uid: string
  token: string
  emailVerified: boolean
  auth: Auth
}
