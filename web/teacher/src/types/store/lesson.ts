import { BaseSubject } from '../api/v1'
import { SchoolType } from '.'

export interface Subject extends BaseSubject {
  schoolType: SchoolType
}

export interface Lesson {
  id: number
  teacherId: string
  studentId: string
  subjectId: number
  startAt: string
  endAt: string
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
