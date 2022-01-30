import { Lesson, ShiftDetailLesson, Subject, StudentShift, TeacherShift } from '~/types/store'

export interface LessonDetail {
  lesson: Lesson
  teacher?: TeacherShift
  student?: StudentShift
  subject?: Subject
}

// ShiftDialogKey シフトダイアログ種別
export const ShiftDialogArray = ['未選択', '講師シフト', '講師授業', '生徒授業希望', '生徒授業', '授業登録'] as const
export type ShiftDialogKey = typeof ShiftDialogArray[number]

export interface TableHeader {
  text: string
  value: string
  align?: 'start' | 'center' | 'end'
  sortable?: boolean
}

export interface SubmissionTableItem {
  date: string
  lessons: ShiftDetailLesson[]
}

export interface TeacherLessonTableItem {
  date: string
  duration: string
  subjectId: number
  studentId: string
}

export interface StudentLessonTableItem {
  date: string
  duration: string
  subjectId: number
  teacherId: string
}

export interface LessonFormItemTeacher {
  id: string
  name: string
  enabled: boolean
}

export interface LessonFormItemStudent {
  id: string
  name: string
  enabled: boolean
}

export interface LessonFormItemSubject {
  id: number
  name: string
  enabled: boolean
}
