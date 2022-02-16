import { SchoolType } from '~/types/store'

export interface Subject {
  id: number
  name: string
  fullname: string
  color: string
  schoolType: SchoolType
  createdAt: string
  updatedAt: string
}

export interface SubjectMap {
  [key: number]: Subject
}

export type SubjectsMap = {
  [key in SchoolType]: Subject[]
}

export interface Teacher {
  id: string
  name?: string
  nameKana?: string
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
}

export interface TeacherMap {
  [key: string]: Teacher
}

export interface Lesson {
  id: number
  shiftId: number
  subjectId: number
  room: number
  teacherId: string
  studentId: string
  startAt: string
  endAt: string
  notes: string
  createdAt: string
  updatedAt: string
}

export interface LessonState {
  subjects: Subject[]
  lessons: Lesson[]
  teachers: Teacher[]
}
