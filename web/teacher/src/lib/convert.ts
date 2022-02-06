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
      その他: [],
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
  const schoolType: SchoolType = schoolTypeNum2schoolTypeString(subjectResponse.schoolType)
  const subject: Subject = {
    ...subjectResponse,
    fullname: `${schoolType}${subjectResponse.name}`,
    schoolType,
  }
  return subject
}

export function subjectResponses2Subjects(responses: SubjectResponse[]): Subject[] {
  const subjects: Subject[] | undefined = responses?.map((response: SubjectResponse) => {
    const schoolType: SchoolType = schoolTypeNum2schoolTypeString(response.schoolType)
    const subject: Subject = {
      ...response,
      fullname: `${schoolType}${response.name}`,
      schoolType,
    }
    return subject
  })
  return subjects || []
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
export function schoolTypeString2schoolTypeNum(schoolType: SchoolType): number {
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
