/**
 * ---------------------------
 * Request
 * ---------------------------
 */
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
  schoolType: number
  grade: number
}
