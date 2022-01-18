import { setup, refresh, setSafetyMode } from '~~/test/helpers/store-helper'
import { SubmissionStore } from '~/store'
import { ApiError } from '~/types/exception'
import { ErrorResponse } from '~/types/api/exception'
import { ShiftStatus, SubmissionStatus } from '~/types/store'

describe('store/submission', () => {
  beforeEach(() => {
    setup()
  })

  afterEach(() => {
    refresh()
  })

  describe('getters', () => {
    it('getSummary', () => {
      expect(SubmissionStore.getSummary).toEqual({
        id: 0,
        year: 0,
        month: 0,
        shiftStatus: ShiftStatus.UNKNOWN,
        submissionStatus: SubmissionStatus.UNKNOWN,
        openAt: '',
        endAt: '',
        createdAt: '',
        updatedAt: '',
      })
    })

    it('getSummaries', () => {
      expect(SubmissionStore.getSummaries).toEqual([])
    })

    it('getShifts', () => {
      expect(SubmissionStore.getShifts).toEqual([])
    })
  })

  describe('actions', () => {
    describe('listTeacherSubmissions', () => {
      describe('success', () => {
        beforeEach(() => {
          setSafetyMode(true)
        })

        it('changing state', async () => {
          await SubmissionStore.listTeacherSubmissions({ teacherId: 'teacherid' })
          expect(SubmissionStore.getSummaries).toEqual([
            {
              id: 1,
              year: 2021,
              month: 2,
              shiftStatus: ShiftStatus.FINISHED,
              submissionStatus: SubmissionStatus.SUBMITTED,
              openAt: '2021-01-01T00:00:00+09:00',
              endAt: '2021-01-15T00:00:00+09:00',
              createdAt: '2021-12-30T19:25:57+09:00',
              updatedAt: '2021-12-30T19:25:57+09:00',
            },
            {
              id: 2,
              year: 2022,
              month: 2,
              shiftStatus: ShiftStatus.WAITING,
              submissionStatus: SubmissionStatus.WAITING,
              openAt: '2021-01-01T00:00:00+09:00',
              endAt: '2021-01-15T00:00:00+09:00',
              createdAt: '2021-12-30T19:25:57+09:00',
              updatedAt: '2021-12-30T19:25:57+09:00',
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
          await expect(SubmissionStore.listTeacherSubmissions({ teacherId: 'teacherid' })).rejects.toThrow(err)
        })
      })
    })

    describe('listShiftDetails', () => {
      describe('success', () => {
        beforeEach(() => {
          setSafetyMode(true)
        })

        it('changing state', async () => {
          await SubmissionStore.listTeacherShifts({ teacherId: 'teacherid', shiftId: 1 })
          expect(SubmissionStore.getSummary).toEqual({
            id: 1,
            year: 2022,
            month: 2,
            shiftStatus: ShiftStatus.FINISHED,
            submissionStatus: SubmissionStatus.SUBMITTED,
            openAt: '2021-01-01T00:00:00+09:00',
            endAt: '2021-01-15T00:00:00+09:00',
            createdAt: '2021-12-30T19:25:57+09:00',
            updatedAt: '2021-12-30T19:25:57+09:00',
          })
          expect(SubmissionStore.getShifts).toEqual([
            {
              date: '20210201',
              isClosed: false,
              lessons: [
                { id: 1, enabled: true, startTime: '1700', endTime: '1830' },
                { id: 2, enabled: false, startTime: '1830', endTime: '2000' },
              ],
            },
            {
              date: '20210202',
              isClosed: true,
              lessons: [],
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
          await expect(SubmissionStore.listTeacherShifts({ teacherId: 'teacherid', shiftId: 1 })).rejects.toThrow(err)
        })
      })
    })

    describe('submitTeacherShifts', () => {
      describe('success', () => {
        beforeEach(() => {
          setSafetyMode(true)
        })

        it('return resolve', async () => {
          await expect(
            SubmissionStore.submitTeacherShifts({ teacherId: 'teacherid', shiftId: 1, lessonIds: [1, 2] })
          ).resolves.toBeUndefined()
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
          await expect(
            SubmissionStore.submitTeacherShifts({ teacherId: 'teacherid', shiftId: 1, lessonIds: [1, 2] })
          ).rejects.toThrow(err)
        })
      })
    })
  })
})
