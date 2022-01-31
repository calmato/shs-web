import { BaseSubject } from '../api/v1'
import { SchoolType } from './common'
import { Student, Teacher } from '.'

export interface Subject extends BaseSubject {
  schoolType: SchoolType
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
  students: Student[]
}

export interface SubjectMap {
  [key: number]: Subject
}

export type SubjectsMap = {
  [key in SchoolType]: Subject[]
}
