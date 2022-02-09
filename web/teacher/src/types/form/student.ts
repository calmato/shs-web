import { SchoolType } from '../store'
import { IHidden, ISelect, ITextField } from './util'

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

export interface IStudentEditSubjectParams {
  schoolType: 1 | 2 | 3
  subjectIds: number[]
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

export interface IStudentEditSubjectOptions {
  schoolType: IHidden
  subjectIds: ISelect
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

export interface StudentEditSubjectForm {
  params: IStudentEditSubjectParams
  options: IStudentEditSubjectOptions
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
  schoolType: 'その他',
  grade: '',
}

export const StudentEditSubjectForElementarySchoolParams: IStudentEditSubjectParams = {
  schoolType: 1,
  subjectIds: [],
}

export const StudentEditSubjectForJuniorHighSchoolParams: IStudentEditSubjectParams = {
  schoolType: 2,
  subjectIds: [],
}

export const StudentEditSubjectForHighSchoolParams: IStudentEditSubjectParams = {
  schoolType: 3,
  subjectIds: [],
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

export const StudentEditSubjectForElementarySchoolOptions: IStudentEditSubjectOptions = {
  schoolType: {} as IHidden,
  subjectIds: {
    label: '受講科目 (小学校)',
  } as ISelect,
}

export const StudentEditSubjectForJuniorHighSchoolOptions: IStudentEditSubjectOptions = {
  schoolType: {} as IHidden,
  subjectIds: {
    label: '受講科目 (中学校)',
  } as ISelect,
}

export const StudentEditSubjectForHighSchoolOptions: IStudentEditSubjectOptions = {
  schoolType: {} as IHidden,
  subjectIds: {
    label: '受講科目 (高等学校)',
  } as ISelect,
}
