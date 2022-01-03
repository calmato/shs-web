/**
 * ---------------------------
 * Response
 * ---------------------------
 */
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

export interface TeacherSubmissionsResponse {
  summaries: TeacherShiftSummary[]
}

export interface TeacherShiftsResponse {
  summary: TeacherShiftSummary
  shifts: TeacherShiftDetail[]
}
