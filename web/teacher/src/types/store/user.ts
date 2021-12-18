import { SchoolType, Role } from './common'

export interface Student {
  id: string
  name?: string
  nameKana?: string
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
  mail: string
  type: SchoolType
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
  createdAt: string
  updatedAt: string
}

export interface UserState {
  students: Student[]
  teachers: Teacher[]
  teachersTotal: number
}

export interface StudentMap {
  [key: string]: Student
}

export interface TeacherMap {
  [key: string]: Teacher
}
