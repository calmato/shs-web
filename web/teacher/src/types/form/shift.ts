import { IDatePicker, IHidden } from './util'

/**
 * ---------------------------
 * interface - params
 * ---------------------------
 */
export interface IShiftSummaryEditScheduleParams {
  summaryId: number
  openDate: string
  endDate: string
}

export interface IShiftsNewParams {
  yearMonth: string
  openDate: string
  endDate: string
  closedDates: string[]
}

export interface IShiftLessonParams {
  lessonId: number
  summaryId: number
  shiftId: number
  room: number
  teacherId: string
  studentId: string
  subjectId: number
}

/**
 * ---------------------------
 * interface - options
 * ---------------------------
 */
export interface IShiftSummaryEditScheduleOptions {
  summaryId: IHidden
  openDate: IDatePicker
  endDate: IDatePicker
}

export interface IShiftsNewOptions {
  yearMonth: IDatePicker
  openDate: IDatePicker
  endDate: IDatePicker
  closedDates: IDatePicker
}

/**
 * ---------------------------
 * interface - form
 * ---------------------------
 */
export interface ShiftSummaryEditScheduleForm {
  params: IShiftSummaryEditScheduleParams
  options: IShiftSummaryEditScheduleOptions
}

export interface ShiftsNewForm {
  params: IShiftsNewParams
  options: IShiftsNewOptions
}

export interface ShiftLessonForm {
  params: IShiftLessonParams
}

/**
 * ---------------------------
 * const - params
 * ---------------------------
 */
export const ShiftSummaryEditScheduleParams: IShiftSummaryEditScheduleParams = {
  summaryId: 0,
  openDate: '',
  endDate: '',
}

export const ShiftsNewParams: IShiftsNewParams = {
  yearMonth: '',
  openDate: '',
  endDate: '',
  closedDates: [],
}

export const ShiftLessonParams: IShiftLessonParams = {
  lessonId: 0,
  summaryId: 0,
  shiftId: 0,
  room: 0,
  teacherId: '',
  studentId: '',
  subjectId: 0,
}

/**
 * ---------------------------
 * const - options
 * ---------------------------
 */
export const ShiftSummaryEditScheduleOptions: IShiftSummaryEditScheduleOptions = {
  summaryId: {} as IHidden,
  openDate: {
    label: 'シフト提出開始日',
    rules: {
      required: true,
    },
  } as IDatePicker,
  endDate: {
    label: 'シフト提出締切日',
    rules: {
      required: true,
      after: '@シフト提出開始日',
    },
  } as IDatePicker,
}

export const ShiftsNewOptions: IShiftsNewOptions = {
  yearMonth: {
    label: 'シフト募集年月',
    rules: {
      required: true,
    },
  } as IDatePicker,
  openDate: {
    label: 'シフト提出開始日',
    rules: {
      required: true,
    },
  } as IDatePicker,
  endDate: {
    label: 'シフト提出締切日',
    rules: {
      required: true,
      after: '@シフト提出開始日',
    },
  } as IDatePicker,
  closedDates: {
    label: '休校日',
    rules: {
      required: true,
    },
  } as IDatePicker,
}
