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

export interface Teacher {
  id: string
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
}

export interface Lesson {
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

export interface SubjectsResponse {
  subjects: Subject[]
}

export interface LessonsResponse {
  lessons: Lesson[]
  teachers: Teacher[]
}
