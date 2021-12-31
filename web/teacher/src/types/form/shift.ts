import { IDatePicker } from './util'

/**
 * ---------------------------
 * interface - params
 * ---------------------------
 */
export interface IShiftsNewParams {
  yearMonth: string
  openDate: string
  endDate: string
  closedDates: string[]
}

/**
 * ---------------------------
 * interface - options
 * ---------------------------
 */
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
export interface ShiftsNewForm {
  params: IShiftsNewParams
  options: IShiftsNewOptions
}

/**
 * ---------------------------
 * const - params
 * ---------------------------
 */
export const ShiftsNewParams: IShiftsNewParams = {
  yearMonth: '',
  openDate: '',
  endDate: '',
  closedDates: [],
}

/**
 * ---------------------------
 * const - options
 * ---------------------------
 */
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
