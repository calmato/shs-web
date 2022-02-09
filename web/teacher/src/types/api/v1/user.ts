/**
 * ---------------------------
 * Request
 * ---------------------------
 */
export interface CreateStudentRequest {
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
  mail: string
  password: string
  passwordConfirmation: string
  schoolType: number
  grade: number
}

export interface UpdateStudentSubjectsRequest {
  schoolType: number
  subjectIds: number[]
}

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

export interface UpdateTeacherMailRequest {
  mail: string
}

export interface UpdateTeacherPasswordRequest {
  password: string
  passwordConfirmation: string
}

/**
 * ---------------------------
 * Response
 * ---------------------------
 */
interface Subject {
  id: number
  name: string
  color: string
  schoolType: 1 | 2 | 3
  createdAt: string
  updatedAt: string
}

export interface Student {
  id: string
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
  mail: string
  schoolType: number
  grade: number
  subjects: Subject[]
  createdAt: string
  updatedAt: string
}

export interface Teacher {
  id: string
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
  mail: string
  role: number
  subjects: { [key: number]: Subject[] }
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

export interface StudentResponse {
  id: string
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
  mail: string
  schoolType: number
  grade: number
  createdAt: string
  updatedAt: string
  subjects: Subject[]
}

export interface TeachersResponse {
  teachers: Teacher[]
  total: number
}

export interface StudentsResponse {
  students: Student[]
  total: number
}
