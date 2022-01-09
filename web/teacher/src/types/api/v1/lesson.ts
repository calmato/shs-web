/**
 * ---------------------------
 * Response
 * ---------------------------
 */
export interface Subject {
  id: number
  name: string
  color: string
  schoolType: 1 | 2 | 3
  createdAt: string
  updatedAt: string
}

export interface SubjectsResponse {
  subjects: Subject[]
}
