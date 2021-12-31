export type SubmissionStatusType = '未提出' | '提出済み'

export interface Submission {
  title: string
  endDate: string
  submissionStatus: SubmissionStatusType
  editStatus: '入力する' | '編集する'
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
