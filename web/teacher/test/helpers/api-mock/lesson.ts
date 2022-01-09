import { SubjectsResponse } from '~/types/api/v1'

export const listSubjects: { [key: string]: SubjectsResponse } = {
  '/v1/subjects': {
    subjects: [
      {
        id: 1,
        name: '国語',
        color: '#F8BBD0',
        schoolType: 1,
        createdAt: '',
        updatedAt: '',
      },
      {
        id: 2,
        name: '数学',
        color: '#BBDEFB',
        schoolType: 2,
        createdAt: '',
        updatedAt: '',
      },
      {
        id: 3,
        name: '英語',
        color: '#FEE6C9',
        schoolType: 3,
        createdAt: '',
        updatedAt: '',
      },
    ],
  },
}
