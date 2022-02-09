import { Role, SchoolType } from './common'
import { Subject, SubjectsMap } from './lesson'

export interface Student {
  id: string
  name?: string
  nameKana?: string
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
  mail: string
  schoolType: SchoolType
  grade: number
  subjects: Subject[]
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
  student: Student
  students: Student[]
  teacher: Teacher
  teachers: Teacher[]
  teachersTotal: number
  studentsTotal: number
}

export interface StudentMap {
  [key: string]: Student
}

export interface TeacherMap {
  [key: string]: Teacher
}
