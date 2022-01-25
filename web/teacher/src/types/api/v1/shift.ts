import { TeacherShiftSummary } from './submission'
import { Lesson } from './lesson'
import { Student, Teacher } from './user'

/**
 * ---------------------------
 * Request
 * ---------------------------
 */
export interface UpdateShiftSummaryScheduleRequest {
  openDate: string
  endDate: string
}

export interface CreateShiftsRequest {
  yearMonth: string
  openDate: string
  endDate: string
  closedDates: string[]
}

/**
 * ---------------------------
 * Response
 * ---------------------------
 */
export interface ShiftSummary {
  id: number
  year: number
  month: number
  status: number
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
  teacher: Teacher
  lessonTotal: number
}

export interface StudentShift {
  student: Student
  lessonTotal: number
  suggestedLessonsTotal: number
}

export interface StudentShiftSummary {
  id: number
  year: number
  month: number
  shiftStatus: number
  submissionStatus: number
  openAt: string
  endAt: string
  createdAt: string
  updatedAt: string
}

export interface SuggestedLesson {
  subjectId: number
  total: number
}

export interface ShiftSummariesResponse {
  summaries: ShiftSummary[]
}

export interface ShiftDetailsResponse {
  summary: ShiftSummary
  shifts: ShiftDetail[]
  rooms: number
  teachers: TeacherShift[]
  students: StudentShift[]
  lessons: Lesson[]
}

export interface TeacherShiftsResponse {
  summary: TeacherShiftSummary
  shifts: ShiftDetail[]
}

export interface StudentShiftsResponse {
  summary: StudentShiftSummary
  shifts: ShiftDetail[]
  suggestedLessons: SuggestedLesson[]
}
