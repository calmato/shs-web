import { SchoolType } from '~/types/store'

export interface SubjectNewForm {
  name: string
  color: string
  schoolType: SchoolType
}

export interface SubjectEditForm extends SubjectNewForm {
  subjectId: number
}
