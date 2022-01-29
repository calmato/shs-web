import { AuthResponse, SubjectResponse } from '~/types/api/v1'
import { Auth, SchoolType, Subject } from '~/types/store'

/**
 * AuthResponseをAuthに変換する関数
 * @param authResponse
 * @returns
 */
export function authResponse2Auth(authResponse: AuthResponse): Auth {
  const auth: Auth = {
    ...authResponse,
    subjects: {
      小学校: authResponse.subjects[1].map((item) => subjectResponse2Subject(item)),
      中学校: authResponse.subjects[2].map((item) => subjectResponse2Subject(item)),
      高校: authResponse.subjects[3].map((item) => subjectResponse2Subject(item)),
    },
  }
  return auth
}

/**
 * SubjectResponseをSubjectに変換する関数
 * @param subjectResponse
 * @returns
 */
export function subjectResponse2Subject(subjectResponse: SubjectResponse): Subject {
  const subject: Subject = {
    ...subjectResponse,
    schoolType: schoolTypeNum2schoolTypeString(subjectResponse.schoolType),
  }
  return subject
}

/**
 * 数値をSchoolTypeに変換する関数
 * @param schoolType
 * @returns
 */
export function schoolTypeNum2schoolTypeString(schoolType: number): SchoolType {
  switch (schoolType) {
    case 1:
      return '小学校'
    case 2:
      return '中学校'
    case 3:
      return '高校'
    default:
      return 'その他'
  }
}

/**
 * SchoolTypeを数値に変換する関数
 * @param schoolType
 * @returns
 */
export function schoolTypeString2schoolTypeNum(schoolType: SchoolType) {
  switch (schoolType) {
    case '小学校':
      return 1
    case '中学校':
      return 2
    case '高校':
      return 3
    case 'その他':
    default:
      return 0
  }
}
