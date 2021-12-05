import { setup, refresh } from '~~/test/helpers/store-helper'
import { LessonStore } from '~/store'

describe('store/lesson', () => {
  beforeEach(() => {
    setup()
  })

  afterEach(() => {
    refresh()
  })

  describe('getters', () => {
    it('getSubjcets', () => {
      expect(LessonStore.getSubjects).toEqual([
        {
          id: 1,
          name: '国語',
          color: '#F8BBD0',
          createdAt: '',
          updatedAt: '',
        },
        {
          id: 2,
          name: '数学',
          color: '#BBDEFB',
          createdAt: '',
          updatedAt: '',
        },
        {
          id: 3,
          name: '英語',
          color: '#FEE6C9',
          createdAt: '',
          updatedAt: '',
        },
      ])
    })

    it('getSubjectMap', () => {
      expect(LessonStore.getSubjectMap).toEqual({
        1: {
          id: 1,
          name: '国語',
          color: '#F8BBD0',
          createdAt: '',
          updatedAt: '',
        },
        2: {
          id: 2,
          name: '数学',
          color: '#BBDEFB',
          createdAt: '',
          updatedAt: '',
        },
        3: {
          id: 3,
          name: '英語',
          color: '#FEE6C9',
          createdAt: '',
          updatedAt: '',
        },
      })
    })

    it('getLessons', () => {
      expect(LessonStore.getLessons).toEqual([
        {
          id: 1,
          teacherId: '000000000000000000001',
          studentId: '123456789012345678901',
          subjectId: 1,
          startAt: '2021-12-10T18:30:00+09:00',
          endAt: '2021-12-10T20:00:00+09:00',
          createdAt: '',
          updatedAt: '',
        },
      ])
    })
  })
})
