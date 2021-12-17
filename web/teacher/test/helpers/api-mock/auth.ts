import { AuthResponse } from '~/types/api/v1'
import { Role } from '~/types/store'

export const showAuth: { [key: string]: AuthResponse } = {
  '/v1/me': {
    id: 'kSByoE6FetnPs5Byk3a9Zx',
    lastName: '中村',
    firstName: '広大',
    lastNameKana: 'なかむら',
    firstNameKana: 'こうだい',
    mail: 'teacher-test001@calmato.jp',
    role: Role.TEACHER,
  },
}
