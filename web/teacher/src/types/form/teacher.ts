import { IHidden, ISelect, ITextField } from './util'

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

export interface ITeacherUpdateMailParams {
  mail: string
}

export interface ITeacherEditSubjectParams {
  schoolType: 1 | 2 | 3
  subjectIds: number[]
}

export interface ITeacherEditRoleParams {
  role: number
}

export interface ITeacherUpdatePasswordParams {
  password: string
  passwordConfirmaion: string
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

export interface ITeacherEditSubjectOptions {
  schoolType: IHidden
  subjectIds: ISelect
}

export interface ITeacherEditRoleOptions {
  role: ISelect
}

export interface ITeacherUpdateMailOptions {
  mail: ITextField
}

export interface ITeacherUpdatePasswordOptions {
  password: ITextField
  passwordConfirmation: ITextField
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

export interface TeacherEditSubjectForm {
  params: ITeacherEditSubjectParams
  options: ITeacherEditSubjectOptions
}

export interface TeacherEditRoleForm {
  params: ITeacherEditRoleParams
  options: ITeacherEditRoleOptions
}

export interface TeacherUpdateMailForm {
  params: ITeacherUpdateMailParams
  options: ITeacherUpdateMailOptions
}

export interface TeacherUpdatePasswordForm {
  params: ITeacherUpdatePasswordParams
  options: ITeacherUpdatePasswordOptions
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

export const TeacherEditSubjectForElementarySchoolParams: ITeacherEditSubjectParams = {
  schoolType: 1,
  subjectIds: [],
}

export const TeacherEditSubjectForJuniorHighSchoolParams: ITeacherEditSubjectParams = {
  schoolType: 2,
  subjectIds: [],
}

export const TeacherEditSubjectForHighSchoolParams: ITeacherEditSubjectParams = {
  schoolType: 3,
  subjectIds: [],
}

export const TeacherEditRoleParams: ITeacherEditRoleParams = {
  role: 0,
}

export const TeacherUpdateMailParams: ITeacherUpdateMailParams = {
  mail: '',
}

export const TeacherUpdatePasswordParams: ITeacherUpdatePasswordParams = {
  password: '',
  passwordConfirmaion: '',
}

/**
 * ---------------------------
 * const - options
 * ---------------------------
 */
export const TeacherNewOptions: ITeacherNewOptions = {
  lastName: {
    label: '???',
    rules: {
      required: true,
      max: 16,
    },
  } as ITextField,
  firstName: {
    label: '???',
    rules: {
      required: true,
      max: 16,
    },
  } as ITextField,
  lastNameKana: {
    label: '??? (??????)',
    rules: {
      required: true,
      hiragana: true,
      max: 32,
    },
  } as ITextField,
  firstNameKana: {
    label: '??? (??????)',
    rules: {
      required: true,
      hiragana: true,
      max: 32,
    },
  } as ITextField,
  mail: {
    label: '?????????????????????',
    rules: {
      required: true,
      email: true,
    },
  } as ITextField,
  password: {
    label: '???????????????',
    rules: {
      required: true,
      password: true,
      min: 6,
      max: 32,
    },
  } as ITextField,
  passwordConfirmation: {
    label: '???????????????(?????????)',
    rules: {
      required: true,
      confirmed: '???????????????',
    },
  } as ITextField,
  role: {
    label: '??????',
    rules: {
      required: true,
    },
  } as ISelect,
}

export const TeacherUpdateMailOptions: ITeacherUpdateMailOptions = {
  mail: {
    label: '?????????????????????',
    rules: {
      required: true,
      email: true,
    },
  } as ITextField,
}

export const TeacherUpdatePasswordOptions: ITeacherUpdatePasswordOptions = {
  password: {
    label: '????????????????????????',
    rules: {
      required: true,
      password: true,
      min: 6,
      max: 32,
    },
  } as ITextField,
  passwordConfirmation: {
    label: '???????????????(?????????)',
    rules: {
      required: true,
      confirmed: '????????????????????????',
    },
  } as ITextField,
}

export const TeacherEditSubjectForElementarySchoolOptions: ITeacherEditSubjectOptions = {
  schoolType: {} as IHidden,
  subjectIds: {
    label: '???????????? (?????????)',
  } as ISelect,
}

export const TeacherEditSubjectForJuniorHighSchoolOptions: ITeacherEditSubjectOptions = {
  schoolType: {} as IHidden,
  subjectIds: {
    label: '???????????? (?????????)',
  } as ISelect,
}

export const TeacherEditSubjectForHighSchoolOptions: ITeacherEditSubjectOptions = {
  schoolType: {} as IHidden,
  subjectIds: {
    label: '???????????? (????????????)',
  } as ISelect,
}

export const TeacherEditRoleOptions: ITeacherEditRoleOptions = {
  role: {
    label: '??????',
    rules: {
      required: true,
    },
  } as ISelect,
}
