import { Lesson } from './lesson'
import { ShiftStatus } from './common'
import { SubmissionStatus, TeacherShiftSummary } from '.'

export interface ShiftSummary {
  id: number
  year: number
  month: number
  status: ShiftStatus
  openAt: string
  endAt: string
  createdAt: string
  updatedAt: string
}

export interface ShiftDetailLesson {
  id: number
  startTime: string
  endTime: string
}

export interface ShiftDetail {
  date: string
  isClosed: boolean
  lessons: ShiftDetailLesson[]
}

export interface SuggestedLesson {
  subjectId: number
  total: number
}

export interface TeacherShift {
  id: string
  name: string
  nameKana: string
  lessonTotal: number
}

export interface TeacherSubmissionDetail {
  id: string
  name: string
  nameKana: string
  summary: TeacherShiftSummary
  shifts: ShiftDetail[]
  submissionTotal: number
}

export interface StudentShift {
  id: string
  name: string
  nameKana: string
  lessonTotal: number
  suggestedLessonsTotal: number
}

export interface StudentShiftSummary {
  id: number
  year: number
  month: number
  shiftStatus: ShiftStatus
  submissionStatus: SubmissionStatus
  openAt: string
  endAt: string
  createdAt: string
  updatedAt: string
}

export interface StudentSubmissionDetail {
  id: string
  name: string
  nameKana: string
  summary: StudentShiftSummary
  shifts: ShiftDetail[]
  suggestedLessons: SuggestedLesson[]
  submissionTotal: number
}

export interface ShiftState {
  summary: ShiftSummary
  summaries: ShiftSummary[]
  details: ShiftDetail[]
  rooms: number
  teachers: TeacherShift[]
  students: StudentShift[]
  lessons: Lesson[]
  teacherSubmission: TeacherSubmissionDetail
  studentSubmission: StudentSubmissionDetail
}
