import { setup, refresh, setSafetyMode } from '~~/test/helpers/store-helper'
import { LessonStore } from '~/store'
import { SchoolType } from '~/types/store'
import { ErrorResponse } from '~/types/api/exception'
import { ApiError } from '~/types/exception'

describe('store/lesson', () => {
  beforeEach(() => {
    setup()
  })

  afterEach(() => {
    refresh()
  })

  describe('getters', () => {
    it('getSubjcets', () => {
      expect(LessonStore.getSubjects).toEqual([])
    })

    it('getSubjectMap', () => {
      expect(LessonStore.getSubjectMap).toEqual({})
    })

    it('getSubjectsMap', () => {
      expect(LessonStore.getSubjectsMap).toEqual({
        [SchoolType.ELEMENTARY_SCHOOL]: [],
        [SchoolType.JUNIOR_HIGH_SCHOOL]: [],
        [SchoolType.HIGH_SCHOOL]: [],
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

  describe('actions', () => {
    describe('getAllSubjects', () => {
      describe('success', () => {
        beforeEach(() => {
          setSafetyMode(true)
        })

        it('return resolve', async () => {
          await expect(LessonStore.getAllSubjects()).resolves.toBeUndefined()
        })

        it('changing state', async () => {
          await LessonStore.getAllSubjects()
          expect(LessonStore.getSubjects).toEqual([
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
          ])
        })
      })

      describe('failure', () => {
        beforeEach(() => {
          setSafetyMode(false)
        })

        it('return reject', async () => {
          const err = new ApiError(400, 'api error', {
            status: 400,
            message: 'api error',
            details: 'some error',
          } as ErrorResponse)
          await expect(LessonStore.getAllSubjects()).rejects.toThrow(err)
        })
      })
    })
  })
})
