interface Subject {
  id: number
  name: string
  color: string
  schoolType: number
  createdAt: string
  updatedAt: string
}

export interface AuthResponse {
  id: string
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
  mail: string
  role: number
  subjects: Map<number, Subject>
}
