/**
 * ---------------------------
 * Request
 * ---------------------------
 */
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
  isClosed: boolean
  lessons: ShiftDetailLesson[]
}

export interface ShiftSummariesResponse {
  summaries: ShiftSummary[]
}

export interface ShiftDetailsResponse {
  summary: ShiftSummary
  shifts: { [key: string]: ShiftDetail }
}