import { ShiftStatus } from './common'

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
  isClosed: boolean
  lessons: ShiftDetailLesson[]
}

export interface ShiftState {
  summary: ShiftSummary
  summaries: ShiftSummary[]
  details: Map<string, ShiftDetail[]>
}

export interface ShiftSummariesMap {
  [key: number]: ShiftSummary[]
}
