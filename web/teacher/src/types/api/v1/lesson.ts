/**
 * ---------------------------
 * Response
 * ---------------------------
 */
export interface Subject {
  id: number
  name: string
  color: string
  schoolType: number
  createdAt: string
  updatedAt: string
}

export interface SubjectsResponse {
  subjects: Subject[]
}
