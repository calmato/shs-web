export interface BaseSubject {
  id: number
  name: string
  color: string
  createdAt: string
  updatedAt: string
}

export interface SubjectResponse extends BaseSubject {
  schoolType: 1 | 2 | 3
}
