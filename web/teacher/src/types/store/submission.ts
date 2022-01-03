import { ShiftStatus, SubmissionStatus } from './common'

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

export interface TeacherShiftDetailLesson {
  id: number
  enabled: boolean
  startTime: string
  endTime: string
}

export interface TeacherShiftDetail {
  date: string
  isClosed: boolean
  lessons: TeacherShiftDetailLesson[]
}

export interface SubmissionState {
  summary: TeacherShiftSummary
  summaries: TeacherShiftSummary[]
  shifts: TeacherShiftDetail[]
}
