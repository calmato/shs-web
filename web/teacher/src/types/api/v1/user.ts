/**
 * ---------------------------
 * Request
 * ---------------------------
 */
export interface CreateTeacherRequest {
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
  mail: string
  password: string
  passwordConfirmation: string
  role: number
}

export interface UpdateTeacherSubjectsRequest {
  schoolType: number
  subjectIds: number[]
}

export interface UpdateTeacherRoleRequest {
  role: number
}

/**
 * ---------------------------
 * Response
 * ---------------------------
 */
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

interface Subject {
  id: number
  name: string
  color: string
  schoolType: number
  createdAt: string
  updatedAt: string
}

export interface TeacherResponse {
  id: string
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
  mail: string
  role: number
  createdAt: string
  updatedAt: string
  subjects: { [key: number]: Subject[] }
}

export interface TeachersResponse {
  teachers: Teacher[]
  total: number
}
