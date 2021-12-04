export interface Teacher {
  id: string
  lastname: string
  firstname: string
  createdAt: string
  updatedAt: string
}

export interface Student {
  id: string
  lastname: string
  firstname: string
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
