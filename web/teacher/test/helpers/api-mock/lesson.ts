import { SubjectsResponse } from '~/types/api/v1'
import { SchoolType } from '~/types/store'

export const listSubjects: { [key: string]: SubjectsResponse } = {
  '/v1/subjects': {
    subjects: [
      {
        id: 1,
        name: '国語',
        color: '#F8BBD0',
        schoolType: SchoolType.ELEMENTARY_SCHOOL,
        createdAt: '',
        updatedAt: '',
      },
      {
        id: 2,
        name: '数学',
        color: '#BBDEFB',
        schoolType: SchoolType.JUNIOR_HIGH_SCHOOL,
        createdAt: '',
        updatedAt: '',
      },
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
}
