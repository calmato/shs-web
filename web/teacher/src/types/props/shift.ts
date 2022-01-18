import { Lesson, Subject, Teacher, Student } from '~/types/store'

export interface LessonDetail {
  lesson: Lesson
  teacher: Teacher
  student: Student
  subject: Subject
}
