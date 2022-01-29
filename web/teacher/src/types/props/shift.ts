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
