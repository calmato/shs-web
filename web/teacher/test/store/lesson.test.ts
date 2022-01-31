import { setup, refresh, setSafetyMode } from '~~/test/helpers/store-helper'
import { LessonStore } from '~/store'
import { ErrorResponse } from '~/types/api/exception'
import { ApiError } from '~/types/exception'
import { SubjectEditForm, SubjectNewForm } from '~/types/form'

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
        小学校: [],
        中学校: [],
        高校: [],
        その他: [],
      })
    })

    it('getLessons', () => {
      expect(LessonStore.getLessons).toEqual([
        {
          id: 1,
          shiftId: 0,
          room: 1,
          teacherId: '000000000000000000001',
          studentId: '123456789012345678901',
          subjectId: 1,
          startAt: '2021-12-10T18:30:00+09:00',
          endAt: '2021-12-10T20:00:00+09:00',
          notes: '',
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
              schoolType: '小学校',
              createdAt: '',
              updatedAt: '',
            },
            {
              id: 2,
              name: '数学',
              color: '#BBDEFB',
              schoolType: '中学校',
              createdAt: '',
              updatedAt: '',
            },
            {
              id: 3,
              name: '英語',
              color: '#FEE6C9',
              schoolType: '高校',
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

    describe('createSubject', () => {
      describe('success', () => {
        beforeEach(() => {
          setSafetyMode(true)
        })

        afterEach(() => {
          jest.clearAllMocks()
        })

        it('return resolve', async () => {
          const payload: SubjectNewForm = {
            name: '算数',
            color: '#DBD0E6',
            schoolType: '小学校',
          }
          await expect(LessonStore.createSubject(payload)).resolves.toBeUndefined()
        })

        it('called getAllSubject', async () => {
          const mockGetAllSubjects = jest.spyOn(LessonStore, 'getAllSubjects')

          const payload: SubjectNewForm = {
            name: '算数',
            color: '#DBD0E6',
            schoolType: '小学校',
          }
          await expect(LessonStore.createSubject(payload)).resolves.toBeUndefined()
          expect(mockGetAllSubjects).toBeCalled()
          expect(mockGetAllSubjects).toBeCalledTimes(1)
        })
      })

      describe('failure', () => {
        beforeEach(() => {
          setSafetyMode(false)
        })

        it('return reject', async () => {
          const invalidPayload: SubjectNewForm = {
            name: '',
            color: '',
            schoolType: '小学校',
          }

          const err = new ApiError(400, 'api error', {
            status: 400,
            message: 'api error',
            details: 'some error',
          } as ErrorResponse)

          try {
            await LessonStore.createSubject(invalidPayload)
          } catch (e) {
            expect(e).toEqual(err)
          }
        })
      })
    })

    describe('editSubject', () => {
      describe('success', () => {
        beforeEach(() => {
          setSafetyMode(true)
        })

        afterEach(() => {
          jest.clearAllMocks()
        })

        it('return resolve', async () => {
          const payload: SubjectEditForm = {
            subjectId: 1,
            name: '算数',
            color: '#DBD0E6',
            schoolType: '小学校',
          }

          await expect(LessonStore.editSubject(payload)).resolves.toBeUndefined()
        })

        it('called getAllSubject', async () => {
          const mockGetAllSubjects = jest.spyOn(LessonStore, 'getAllSubjects')

          const payload: SubjectEditForm = {
            subjectId: 1,
            name: '算数',
            color: '#DBD0E6',
            schoolType: '小学校',
          }

          await expect(LessonStore.editSubject(payload)).resolves.toBeUndefined()

          expect(mockGetAllSubjects).toBeCalled()
          expect(mockGetAllSubjects).toBeCalledTimes(1)
        })
      })

      describe('failure', () => {
        beforeEach(() => {
          setSafetyMode(false)
        })

        it('return reject', async () => {
          const invalidPayload: SubjectEditForm = {
            subjectId: 1,
            name: '算数',
            color: '#DBD0E6',
            schoolType: '小学校',
          }

          const err = new ApiError(400, 'api error', {
            status: 400,
            message: 'api error',
            details: 'some error',
          } as ErrorResponse)

          try {
            await LessonStore.createSubject(invalidPayload)
          } catch (e) {
            expect(e).toEqual(err)
          }
        })
      })
    })

    describe('deleteSubject', () => {
      describe('success', () => {
        beforeEach(() => {
          setSafetyMode(true)
        })

        afterEach(() => {
          jest.clearAllMocks()
        })

        it('return resolve', async () => {
          await expect(LessonStore.deleteSubject(1)).resolves.toBeUndefined()
        })

        it('called getAllSubject', async () => {
          const mockGetAllSubjects = jest.spyOn(LessonStore, 'getAllSubjects')

          await expect(LessonStore.deleteSubject(1)).resolves.toBeUndefined()
          expect(mockGetAllSubjects).toBeCalled()
          expect(mockGetAllSubjects).toBeCalledTimes(1)
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

          try {
            await LessonStore.deleteSubject(-1)
          } catch (e) {
            expect(e).toEqual(err)
          }
        })
      })
    })
  })
})
