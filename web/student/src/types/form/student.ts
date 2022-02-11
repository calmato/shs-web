import { ITextField } from './util'

/**
 * ---------------------------
 * interface - params
 * ---------------------------
 */
export interface IStudentUpdatePasswordParams {
  password: string
  passwordConfirmaion: string
}

/**
 * ---------------------------
 * interface - options
 * ---------------------------
 */
export interface IStudentUpdatePasswordOptions {
  password: ITextField
  passwordConfirmation: ITextField
}

/**
 * ---------------------------
 * interface - form
 * ---------------------------
 */
export interface StudentUpdatePasswordForm {
  params: IStudentUpdatePasswordParams
  options: IStudentUpdatePasswordOptions
}

/**
 * ---------------------------
 * const - params
 * ---------------------------
 */
export const StudentUpdatePasswordParams: IStudentUpdatePasswordParams = {
  password: '',
  passwordConfirmaion: '',
}

/**
 * ---------------------------
 * const - options
 * ---------------------------
 */
export const StudentUpdatePasswordOptions: IStudentUpdatePasswordOptions = {
  password: {
    label: '変更後パスワード',
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
      confirmed: '変更後パスワード',
    },
  } as ITextField,
}
