import { SchoolType } from '../store'
import { ISelect, ITextField } from './util'

/**
 * ---------------------------
 * interface - params
 * ---------------------------
 */
export interface IStudentNewParams {
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
  mail: string
  password: string
  passwordConfirmation: string
  schoolType: SchoolType
  grade: string
}

/**
 * ---------------------------
 * interface - options
 * ---------------------------
 */
export interface IStudentNewOptions {
  lastName: ITextField
  firstName: ITextField
  lastNameKana: ITextField
  firstNameKana: ITextField
  mail: ITextField
  password: ITextField
  passwordConfirmation: ITextField
  schoolType: ISelect
  grade: ISelect
}

/**
 * ---------------------------
 * interface - form
 * ---------------------------
 */
export interface StudentNewForm {
  params: IStudentNewParams
  options: IStudentNewOptions
}

/**
 * ---------------------------
 * const - params
 * ---------------------------
 */
export const StudentNewParams: IStudentNewParams = {
  lastName: '',
  firstName: '',
  lastNameKana: '',
  firstNameKana: '',
  mail: '',
  password: '',
  passwordConfirmation: '',
  schoolType: '',
  grade: '',
}

/**
 * ---------------------------
 * const - options
 * ---------------------------
 */
export const StudentNewOptions: IStudentNewOptions = {
  lastName: {
    label: '姓',
    rules: {
      required: true,
      max: 16,
    },
  } as ITextField,
  firstName: {
    label: '名',
    rules: {
      required: true,
      max: 16,
    },
  } as ITextField,
  lastNameKana: {
    label: '姓 (かな)',
    rules: {
      required: true,
      hiragana: true,
      max: 32,
    },
  } as ITextField,
  firstNameKana: {
    label: '名 (かな)',
    rules: {
      required: true,
      hiragana: true,
      max: 32,
    },
  } as ITextField,
  mail: {
    label: 'メールアドレス',
    rules: {
      required: true,
      email: true,
    },
  } as ITextField,
  password: {
    label: 'パスワード',
    rules: {
      required: true,
      password: true,
      min: 6,
      max: 32,
    },
  } as ITextField,
  passwordConfirmation: {
    label: 'パスワード(確認用)',
    rules: {
      required: true,
      confirmed: 'パスワード',
    },
  } as ITextField,
  schoolType: {
    label: '校種',
    rules: {
      required: true,
    },
  } as ISelect,
  grade: {
    label: '学年',
    rules: {
      required: true,
    },
  } as ISelect,
}
