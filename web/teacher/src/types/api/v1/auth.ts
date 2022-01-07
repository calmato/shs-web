import { SubjectResponse } from '~/types/api/v1'

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
