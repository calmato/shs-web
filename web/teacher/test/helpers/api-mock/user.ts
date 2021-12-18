import { TeacherResponse, TeachersResponse } from '~/types/api/v1'

export const listTeachers: { [key: string]: TeachersResponse } = {
  '/v1/teachers': {
    teachers: [
      {
        id: '000000000000000000001',
        lastName: '中村',
        firstName: '太郎',
        lastNameKana: 'なかむら',
        firstNameKana: 'たろう',
        mail: 'teacher-001@calmato.jp',
        role: 1,
        createdAt: '2021-12-02T18:30:00+09:00',
        updatedAt: '2021-12-02T18:30:00+09:00',
      },
      {
        id: '000000000000000000002',
        lastName: '西山',
        firstName: '幸子',
        lastNameKana: 'にしやま',
        firstNameKana: 'さちこ',
        mail: 'teacher-002@calmato.jp',
        role: 1,
        createdAt: '2021-12-02T18:30:00+09:00',
        updatedAt: '2021-12-02T18:30:00+09:00',
      },
      {
        id: '000000000000000000003',
        lastName: '鈴木',
        firstName: '小太郎',
        lastNameKana: 'すずき',
        firstNameKana: 'こたろう',
        mail: 'teacher-003@calmato.jp',
        role: 2,
        createdAt: '2021-12-02T18:30:00+09:00',
        updatedAt: '2021-12-02T18:30:00+09:00',
      },
    ],
    total: 3,
  },
  '/v1/teachers?limit=20&offset=2': {
    teachers: [
      {
        id: '000000000000000000003',
        lastName: '鈴木',
        firstName: '小太郎',
        lastNameKana: 'すずき',
        firstNameKana: 'こたろう',
        mail: 'teacher-003@calmato.jp',
        role: 2,
        createdAt: '2021-12-02T18:30:00+09:00',
        updatedAt: '2021-12-02T18:30:00+09:00',
      },
    ],
    total: 3,
  },
}

export const createTeacher: { [key: string]: TeacherResponse } = {
  '/v1/teachers': {
    id: '000000000000000000001',
    lastName: '中村',
    firstName: '太郎',
    lastNameKana: 'なかむら',
    firstNameKana: 'たろう',
    mail: 'teacher-001@calmato.jp',
    role: 1,
    createdAt: '2021-12-02T18:30:00+09:00',
    updatedAt: '2021-12-02T18:30:00+09:00',
  },
}
