import { TeacherResponse, TeachersResponse } from '~/types/api/v1'
import { SchoolType } from '~/types/store'

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

export const showTeacher: { [key: string]: TeacherResponse } = {
  '/v1/teachers/000000000000000000001': {
    id: '000000000000000000001',
    lastName: '中村',
    firstName: '太郎',
    lastNameKana: 'なかむら',
    firstNameKana: 'たろう',
    mail: 'teacher-001@calmato.jp',
    role: 1,
    subjects: {
      [SchoolType.ELEMENTARY_SCHOOL]: [
        {
          id: 1,
          name: '国語',
          color: '#F8BBD0',
          schoolType: SchoolType.ELEMENTARY_SCHOOL,
          createdAt: '',
          updatedAt: '',
        },
      ],
      [SchoolType.JUNIOR_HIGH_SCHOOL]: [
        {
          id: 2,
          name: '数学',
          color: '#BBDEFB',
          schoolType: SchoolType.JUNIOR_HIGH_SCHOOL,
          createdAt: '',
          updatedAt: '',
        },
      ],
      [SchoolType.HIGH_SCHOOL]: [
        {
          id: 3,
          name: '英語',
          color: '#FEE6C9',
          schoolType: SchoolType.HIGH_SCHOOL,
          createdAt: '',
          updatedAt: '',
        },
      ],
    },
    createdAt: '2021-12-02T18:30:00+09:00',
    updatedAt: '2021-12-02T18:30:00+09:00',
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
    subjects: {
      [SchoolType.ELEMENTARY_SCHOOL]: [
        {
          id: 1,
          name: '国語',
          color: '#F8BBD0',
          schoolType: SchoolType.ELEMENTARY_SCHOOL,
          createdAt: '',
          updatedAt: '',
        },
      ],
      [SchoolType.JUNIOR_HIGH_SCHOOL]: [
        {
          id: 2,
          name: '数学',
          color: '#BBDEFB',
          schoolType: SchoolType.JUNIOR_HIGH_SCHOOL,
          createdAt: '',
          updatedAt: '',
        },
      ],
      [SchoolType.HIGH_SCHOOL]: [
        {
          id: 3,
          name: '英語',
          color: '#FEE6C9',
          schoolType: SchoolType.HIGH_SCHOOL,
          createdAt: '',
          updatedAt: '',
        },
      ],
    },
    createdAt: '2021-12-02T18:30:00+09:00',
    updatedAt: '2021-12-02T18:30:00+09:00',
  },
}

export const updateTeacherSubjects: { [key: string]: {} } = {
  '/v1/teachers/000000000000000000001/subjects': {},
}

export const updateTeacherRole: { [key: string]: {} } = {
  '/v1/teachers/000000000000000000001/role': {},
}

export const deleteTeacher: { [key: string]: {} } = {
  '/v1/teachers/000000000000000000001': {},
}
