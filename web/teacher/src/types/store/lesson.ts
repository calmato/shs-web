export interface Subject {
  id: number
  name: string
  color: string
  createdAt: string
  updatedAt: string
}

export interface Lesson {
  id: number
  teacherId: string
  studentId: string
  subjectId: number
  startAt: string
  endAt: string
  createdAt: string
  updatedAt: string
}

export interface LessonState {
  subjects: Subject[]
  lessons: Lesson[]
}

export interface SubjectMap {
  [key: number]: Subject
}
