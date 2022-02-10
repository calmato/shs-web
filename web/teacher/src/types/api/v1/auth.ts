import { SubjectResponse } from '~/types/api/v1'

/**
 * ---------------------------
 * Request
 * ---------------------------
 */
export interface UpdateMyMailRequest {
  mail: string
}

export interface UpdateMyPasswordRequest {
  password: string
  passwordConfirmation: string
}

/**
 * ---------------------------
 * Response
 * ---------------------------
 */
export interface AuthResponse {
  id: string
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
  mail: string
  role: number
  subjects: {
    [key in 1 | 2 | 3]: SubjectResponse[]
  }
}
