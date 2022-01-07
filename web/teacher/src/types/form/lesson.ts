import { SchoolType } from '~/types/store'

// ---------------------------
// interface - form
// ---------------------------
export interface SubjectUpdateForm {
  schoolType: SchoolType
  subjectIds: number[]
}
