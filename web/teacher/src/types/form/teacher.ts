import { ISelect, ITextField } from './util'

/**
 * ---------------------------
 * interface - params
 * ---------------------------
 */
export interface ITeacherNewParams {
  lastName: string
  firstName: string
  lastNameKana: string
  firstNameKana: string
  mail: string
  password: string
  passwordConfirmation: string
  role: number
}

/**
 * ---------------------------
 * interface - options
 * ---------------------------
 */
export interface ITeacherNewOptions {
  lastName: ITextField
  firstName: ITextField
  lastNameKana: ITextField
  firstNameKana: ITextField
  mail: ITextField
  password: ITextField
  passwordConfirmation: ITextField
  role: ISelect
}

/**
 * ---------------------------
 * interface - form
 * ---------------------------
 */
export interface TeacherNewForm {
  params: ITeacherNewParams
  options: ITeacherNewOptions
}

/**
 * ---------------------------
 * const - params
 * ---------------------------
 */
export const TeacherNewParams: ITeacherNewParams = {
  lastName: '',
  firstName: '',
  lastNameKana: '',
  firstNameKana: '',
  mail: '',
  password: '',
  passwordConfirmation: '',
  role: 0,
}

/**
 * ---------------------------
 * const - options
 * ---------------------------
 */
export const TeacherNewOptions: ITeacherNewOptions = {
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
  role: {
    label: '役職',
    rules: {
      required: true,
    },
  } as ISelect,
}
