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
  date: string
  isClosed: boolean
  lessons: ShiftDetailLesson[]
}

export interface TeacherShift {
  id: string
  name: string
  nameKana: string
  lessonTotal: number
}

export interface ShiftState {
  summary: ShiftSummary
  summaries: ShiftSummary[]
  details: ShiftDetail[]
  rooms: number
  teachers: TeacherShift[]
}
