/**
 * ---------------------------
 * Request
 * ---------------------------
 */
export interface SubmissionRequest {
  shiftIds: number[]
}

/**
 * ---------------------------
 * Response
 * ---------------------------
 */
export interface SubmissionSummary {
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

export interface SubmissionResponse {
  summary: SubmissionSummary
  shifts: SubmissionDetail[]
}

export interface SubmissionsResponse {
  summaries: SubmissionSummary[]
}
