import { SchoolType } from './common'

export interface Auth {
  id: string
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
  mail: string
  schoolType: SchoolType
  grade: number
}

export interface AuthState {
  uid: string
  token: string
  emailVerified: boolean
  auth: Auth
}
