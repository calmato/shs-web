import { Lesson, Subject, StudentShift, TeacherShift } from '~/types/store'

export interface LessonDetail {
  lesson: Lesson
  teacher?: TeacherShift
  student?: StudentShift
  subject?: Subject
}
