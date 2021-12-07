// Role 権限
export enum Role {
  TEACHER = 0, // 講師
  ADMINISTRATOR = 1, // 管理者
}

// SchoolType 校種
export enum SchoolType {
  ElementarySchool = 0, // 小学校
  JuniorHighSchool = 1, // 中学校
  HighSchool = 2, // 高等学校
}

export interface Teacher {
  id: string
  name?: string
  lastname: string
  firstname: string
  role: Role
  createdAt: string
  updatedAt: string
}

export interface Student {
  id: string
  name?: string
  lastname: string
  firstname: string
  type: SchoolType
  grade: number
  createdAt: string
  updatedAt: string
}

export interface UserState {
  message: string // TODO: remove
  teachers: Teacher[]
  students: Student[]
}

export interface TeacherMap {
  [key: string]: Teacher
}

export interface StudentMap {
  [key: string]: Student
}
