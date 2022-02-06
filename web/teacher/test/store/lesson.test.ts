import { setup, refresh, setSafetyMode, setIsAxiosMockValue } from '~~/test/helpers/store-helper'
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
      expect(LessonStore.getLessons).toEqual([])
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
              fullname: '小学校国語',
              color: '#F8BBD0',
              schoolType: '小学校',
              createdAt: '',
              updatedAt: '',
            },
            {
              id: 2,
              name: '数学',
              fullname: '中学校数学',
              color: '#BBDEFB',
              schoolType: '中学校',
              createdAt: '',
              updatedAt: '',
            },
            {
              id: 3,
              name: '英語',
              fullname: '高校英語',
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
          setIsAxiosMockValue(true)
        })

        it('return reject', async () => {
          const err = new ApiError(400, 'api error', {
            status: 400,
            message: 'api error',
            details: 'some error',
          } as ErrorResponse)
          await expect(LessonStore.getAllSubjects()).rejects.toThrow(err)
        })

        it('throw internal server error', async () => {
          setIsAxiosMockValue(false)

          const err = new Error('internal server error')
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
          setIsAxiosMockValue(true)
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

        it('not called getAllSubject', async () => {
          const mockGetAllSubjects = jest.spyOn(LessonStore, 'getAllSubjects')

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
            expect(mockGetAllSubjects).toBeCalledTimes(0)
          }
        })

        it('throw internal server error', async () => {
          setIsAxiosMockValue(false)

          const invalidPayload: SubjectNewForm = {
            name: '',
            color: '',
            schoolType: '小学校',
          }

          const err = new Error('internal server error')

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
          setIsAxiosMockValue(true)
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
            await LessonStore.editSubject(invalidPayload)
          } catch (e) {
            expect(e).toEqual(err)
          }
        })

        it('notcalled getAllSubject', async () => {
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
            await LessonStore.editSubject(invalidPayload)
          } catch (e) {
            expect(e).toEqual(err)
          }
        })

        it('throw internal server error', async () => {
          setIsAxiosMockValue(false)

          const invalidPayload: SubjectEditForm = {
            subjectId: 1,
            name: '算数',
            color: '#DBD0E6',
            schoolType: '小学校',
          }

          const err = new Error('internal server error')

          try {
            await LessonStore.editSubject(invalidPayload)
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
          setIsAxiosMockValue(true)
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

        it('not called getAllSubject', async () => {
          const mockGetAllSubjects = jest.spyOn(LessonStore, 'getAllSubjects')

          const err = new ApiError(400, 'api error', {
            status: 400,
            message: 'api error',
            details: 'some error',
          } as ErrorResponse)

          try {
            await LessonStore.deleteSubject(-1)
          } catch (e) {
            expect(e).toEqual(err)
            expect(mockGetAllSubjects).toBeCalledTimes(0)
          }
        })

        it('throw internal server error', async () => {
          setIsAxiosMockValue(false)

          const err = new Error('internal server error')

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
