import { BaseSubject } from '../api/v1'
import { SchoolType } from './common'

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
}

export interface SubjectMap {
  [key: number]: Subject
}

export type SubjectsMap = {
  [key in Exclude<SchoolType, 'その他'>]: Subject[]
}
