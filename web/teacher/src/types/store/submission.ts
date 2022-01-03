import { ShiftStatus, SubmissionStatus } from './common'

export type SubmissionStatusType = '未提出' | '提出済み'

export interface Submission {
  title: string
  endDate: string
  submissionStatus: SubmissionStatusType
  editStatus: '入力する' | '編集する'
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
  submissions: Submission[] // TODO: remove
  summary: TeacherShiftSummary
  summaries: TeacherShiftSummary[]
  shifts: TeacherShiftDetail[]
}
