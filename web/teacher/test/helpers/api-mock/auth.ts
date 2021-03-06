import { AuthResponse } from '~/types/api/v1'
import { Role } from '~/types/store'

const subjects: AuthResponse['subjects'] = {
  '1': [],
  '2': [],
  '3': [
    {
      id: 1,
      name: '国語',
      color: '#F8BBD0',
      schoolType: 3,
      createdAt: '2021-12-02T18:30:00+09:00',
      updatedAt: '2021-12-02T18:30:00+09:00',
    },
  ],
}

export const showAuth: { [key: string]: AuthResponse } = {
  '/v1/me': {
    id: 'kSByoE6FetnPs5Byk3a9Zx',
    lastName: '中村',
    firstName: '広大',
    lastNameKana: 'なかむら',
    firstNameKana: 'こうだい',
    mail: 'teacher-test001@calmato.jp',
    role: Role.TEACHER,
    subjects,
  },
}
