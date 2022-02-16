import { SchoolType } from './common'
import { Subject } from './lesson'

export interface Auth {
  id: string
  name?: string
  nameKana?: string
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
  mail: string
  schoolType: SchoolType
  grade: number
  subjects: Subject[]
}

export interface AuthState {
  uid: string
  token: string
  emailVerified: boolean
  auth: Auth
}
