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
  subjects: Map<SchoolType, Subject[]>
}

export interface AuthState {
  uid: string
  token: string
  emailVerified: boolean
  auth: Auth
}
