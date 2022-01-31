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

export interface UpdateShiftSummaryDecidedRequest {
  decided: boolean
}

export interface CreateShiftsRequest {
  yearMonth: string
  openDate: string
  endDate: string
  closedDates: string[]
}

export interface CreateShiftLessonRequest {
  shiftId: number
  room: number
  subjectId: number
  teacherId: string
  studentId: string
}

export interface UpdateShiftLessonRequest {
  shiftId: number
  room: number
  subjectId: number
  teacherId: string
  studentId: string
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
  decided: boolean
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

export interface SuggestedLesson {
  subjectId: number
  total: number
}

export interface TeacherShift {
  teacher: Teacher
  lessonTotal: number
}

export interface TeacherShiftSummary {
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

export interface StudentShift {
  student: Student
  suggestedLessons: SuggestedLesson[]
  suggestedLessonsTotal: number
  lessonTotal: number
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

export interface ShiftLessonResponse {
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

export interface ShiftLessonsResponse {
  teachers: Teacher[]
  students: Student[]
  lessons: ShiftLesson[]
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
