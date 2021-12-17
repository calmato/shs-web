export interface Teacher {
  id: string
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
  mail: string
  role: number
  createdAt: string
  updatedAt: string
}

export interface TeachersResponse {
  teachers: Teacher[]
}
