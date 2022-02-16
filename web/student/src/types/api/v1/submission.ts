/**
 * ---------------------------
 * Request
 * ---------------------------
 */
export interface SubmissionLesson {
  subjectId: number
  total: number
}

export interface SubmissionTemplateLesson {
  enabled: boolean
  startTime: string
  endTime: string
}

export interface SubmissionTemplate {
  weekday: number
  lessons: SubmissionTemplateLesson[]
}

export interface SubmissionRequest {
  suggestedLessons: SubmissionLesson[]
  shiftIds: number[]
}

export interface SubmissionTemplateRequest {
  schedules: SubmissionTemplate[]
  suggestedLessons: SubmissionLesson[]
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
  suggestedLessons: SubmissionLesson[]
}

export interface SubmissionsResponse {
  summaries: SubmissionSummary[]
}

export interface SubmissionTemplateResponse {
  schedules: SubmissionTemplate[]
  suggestedLessons: SubmissionLesson[]
}
