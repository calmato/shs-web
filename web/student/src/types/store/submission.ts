import { ShiftStatus, SubmissionStatus } from './common'

export interface SubmissionSummary {
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

export interface SubmissionDetailLesson {
  id: number
  enabled: boolean
  startTime: string
  endTime: string
}

export interface SubmissionDetail {
  date: string
  isClosed: boolean
  lessons: SubmissionDetailLesson[]
}

export interface SubmissionLesson {
  subjectId: number
  total: number
}

export interface SubmissionState {
  summary: SubmissionSummary
  summaries: SubmissionSummary[]
  shifts: SubmissionDetail[]
  suggestedLessons: SubmissionLesson[]
}

export interface SummaryParams {
  id: number
  year: number
  month: number
  status: number
  openAt: string
  endAt: string
  createdAt: string
  updatedAt: string
}

export interface LessonParams {
  id: number
  startTime: string
  endTime: string
}

export interface ShiftParams {
  date: string
  isClosed: boolean
  lessons: LessonParams[]
}

export interface SubmissionEditState {
  summary: SummaryParams
  shifts: ShiftParams[]
}
