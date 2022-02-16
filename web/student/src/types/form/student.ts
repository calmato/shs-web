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

export interface IStudentUpdateMailParams {
  mail: string
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

export interface IStudentUpdateMailOptions {
  mail: ITextField
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

export interface StudentUpdateMailForm {
  params: IStudentUpdateMailParams
  options: IStudentUpdateMailOptions
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

export const StudentUpdateMailParams: IStudentUpdateMailParams = {
  mail: '',
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

export const StudentUpdateMailOptions: IStudentUpdateMailOptions = {
  mail: {
    label: 'メールアドレス',
    rules: {
      required: true,
      email: true,
    },
  } as ITextField,
}
