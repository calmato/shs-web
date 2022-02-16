import { Lesson } from './lesson'
import { ShiftStatus } from './common'
import { Student, SubmissionStatus, Teacher } from '.'

export interface ShiftSummary {
  id: number
  year: number
  month: number
  decided: boolean
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
  remainingTotal: number
}

export interface ShiftLesson {
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

export interface ShiftLessonDetail {
  lessonId: number
  summaryId: number
  shiftId: number
  room: number
  date: string
  current?: ShiftLesson
  teachers: Teacher[]
  students: Student[]
  lessons: ShiftLesson[]
}

export interface TeacherShift {
  id: string
  name: string
  nameKana: string
  isSubmit: boolean
  lessonTotal: number
}

export interface TeacherShiftSummary {
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
  isSubmit: boolean
  suggestedLessons: SuggestedLesson[]
  suggestedLessonsTotal: number
  lessonTotal: number
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

export interface ShiftUserLesson {
  current: string
  lessons: ShiftLesson[]
  teachers: { [key: string]: TeacherShift }
  students: { [key: string]: StudentShift }
  total: number
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
  teacherLessons: ShiftUserLesson
  studentSubmission: StudentSubmissionDetail
  studentLessons: ShiftUserLesson
  lessonDetail: ShiftLessonDetail
}
