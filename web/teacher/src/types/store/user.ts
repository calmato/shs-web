import { Role } from './common'
import { SubjectsMap } from './lesson'

export interface Student {
  id: string
  name?: string
  nameKana?: string
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
  mail: string
  type: 0 | 1 | 2 | 3
  grade: number
  createdAt: string
  updatedAt: string
}

export interface Teacher {
  id: string
  name?: string
  nameKana?: string
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
  mail: string
  role: Role
  subjects: SubjectsMap
  createdAt: string
  updatedAt: string
}

export interface UserState {
  students: Student[]
  teacher: Teacher
  teachers: Teacher[]
  teachersTotal: number
}

export interface StudentMap {
  [key: string]: Student
}

export interface TeacherMap {
  [key: string]: Teacher
}
