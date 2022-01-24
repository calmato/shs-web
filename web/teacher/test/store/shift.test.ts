import { setup, refresh, setSafetyMode } from '~~/test/helpers/store-helper'
import { ShiftStore } from '~/store'
import { ApiError } from '~/types/exception'
import { ErrorResponse } from '~/types/api/exception'
import { ShiftStatus } from '~/types/store'
import {
  ShiftsNewForm,
  ShiftsNewOptions,
  ShiftSummaryEditScheduleForm,
  ShiftSummaryEditScheduleOptions,
} from '~/types/form'

describe('store/shift', () => {
  beforeEach(() => {
    setup()
  })

  afterEach(() => {
    refresh()
  })

  describe('getters', () => {
    it('getSummary', () => {
      expect(ShiftStore.getSummary).toEqual({
        id: 0,
        year: 0,
        month: 0,
        status: ShiftStatus.UNKNOWN,
        openAt: '',
        endAt: '',
        createdAt: '',
        updatedAt: '',
      })
    })

    it('getSummaries', () => {
      expect(ShiftStore.getSummaries).toEqual([])
    })

    it('getDetails', () => {
      expect(ShiftStore.getDetails).toEqual([])
    })

    it('getRooms', () => {
      expect(ShiftStore.getRooms).toBe(4)
    })

    it('getTeachers', () => {
      expect(ShiftStore.getTeachers).toEqual([])
    })

    it('getStudents', () => {
      expect(ShiftStore.getStudents).toEqual([])
    })

    it('getLessons', () => {
      expect(ShiftStore.getLessons).toEqual([])
    })
  })

  describe('actions', () => {
    describe('listShiftSummaries', () => {
      describe('success', () => {
        beforeEach(() => {
          setSafetyMode(true)
        })

        it('changing state when params is zero', async () => {
          await ShiftStore.listShiftSummaries({ limit: 0, offset: 0, status: ShiftStatus.UNKNOWN })
          expect(ShiftStore.getSummaries).toEqual([
            {
              id: 1,
              year: 2021,
              month: 2,
              status: ShiftStatus.FINISHED,
              openAt: '2021-01-01T00:00:00+09:00',
              endAt: '2021-01-15T00:00:00+09:00',
              createdAt: '2021-12-30T19:25:57+09:00',
              updatedAt: '2021-12-30T19:25:57+09:00',
            },
            {
              id: 2,
              year: 2022,
              month: 2,
              status: ShiftStatus.WAITING,
              openAt: '2021-01-01T00:00:00+09:00',
              endAt: '2021-01-15T00:00:00+09:00',
              createdAt: '2021-12-30T19:25:57+09:00',
              updatedAt: '2021-12-30T19:25:57+09:00',
            },
          ])
        })

        it('changing state when limit 20 and status 1', async () => {
          await ShiftStore.listShiftSummaries({ limit: 20, offset: 0, status: ShiftStatus.WAITING })
          expect(ShiftStore.getSummaries).toEqual([
            {
              id: 2,
              year: 2022,
              month: 2,
              status: ShiftStatus.WAITING,
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
          await expect(
            ShiftStore.listShiftSummaries({ limit: 0, offset: 0, status: ShiftStatus.UNKNOWN })
          ).rejects.toThrow(err)
        })
      })
    })

    describe('updateShiftSummarySchedule', () => {
      let form: ShiftSummaryEditScheduleForm
      beforeEach(() => {
        form = {
          params: {
            summaryId: 1,
            openDate: '2021-01-01',
            endDate: '2021-01-15',
          },
          options: ShiftSummaryEditScheduleOptions,
        }
      })

      describe('success', () => {
        beforeEach(() => {
          setSafetyMode(true)
        })

        it('return', async () => {
          await expect(ShiftStore.updateShiftSummarySchedule({ form })).resolves.toBeUndefined()
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
          await expect(ShiftStore.updateShiftSummarySchedule({ form })).rejects.toThrow(err)
        })
      })
    })

    describe('createShifts', () => {
      let form: ShiftsNewForm
      beforeEach(() => {
        form = {
          params: {
            yearMonth: '2022-02',
            openDate: '2021-01-01',
            endDate: '2021-01-15',
            closedDates: ['2021-02-02'],
          },
          options: ShiftsNewOptions,
        }
      })

      describe('success', () => {
        beforeEach(() => {
          setSafetyMode(true)
        })

        it('changing state', async () => {
          await ShiftStore.createShifts({ form })
          expect(ShiftStore.getSummary).toEqual({
            id: 1,
            year: 2022,
            month: 2,
            status: ShiftStatus.FINISHED,
            openAt: '2021-01-01T00:00:00+09:00',
            endAt: '2021-01-15T00:00:00+09:00',
            createdAt: '2021-12-30T19:25:57+09:00',
            updatedAt: '2021-12-30T19:25:57+09:00',
          })
          expect(ShiftStore.getDetails).toEqual([
            {
              date: '20210201',
              isClosed: false,
              lessons: [
                { id: 1, startTime: '1700', endTime: '1830' },
                { id: 2, startTime: '1830', endTime: '2000' },
              ],
            },
            {
              date: '20210202',
              isClosed: true,
              lessons: [],
            },
          ])
          expect(ShiftStore.getSummaries).toEqual([
            {
              id: 1,
              year: 2022,
              month: 2,
              status: ShiftStatus.FINISHED,
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
          await expect(ShiftStore.createShifts({ form })).rejects.toThrow(err)
        })
      })
    })

    describe('deleteShifts', () => {
      describe('success', () => {
        beforeEach(() => {
          setSafetyMode(true)
        })

        it('return', async () => {
          await expect(ShiftStore.deleteShifts({ summaryId: 1 })).resolves.toBeUndefined()
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
          await expect(ShiftStore.deleteShifts({ summaryId: 1 })).rejects.toThrow(err)
        })
      })
    })

    describe('listShiftDetails', () => {
      describe('success', () => {
        beforeEach(() => {
          setSafetyMode(true)
        })

        it('changing state', async () => {
          await ShiftStore.listShiftDetails({ summaryId: 1 })
          expect(ShiftStore.getSummary).toEqual({
            id: 1,
            year: 2022,
            month: 2,
            status: ShiftStatus.FINISHED,
            openAt: '2021-01-01T00:00:00+09:00',
            endAt: '2021-01-15T00:00:00+09:00',
            createdAt: '2021-12-30T19:25:57+09:00',
            updatedAt: '2021-12-30T19:25:57+09:00',
          })
          expect(ShiftStore.getDetails).toEqual([
            {
              date: '20210201',
              isClosed: false,
              lessons: [
                { id: 1, startTime: '1700', endTime: '1830' },
                { id: 2, startTime: '1830', endTime: '2000' },
              ],
            },
            {
              date: '20210202',
              isClosed: true,
              lessons: [],
            },
          ])
          expect(ShiftStore.getTeachers).toEqual([
            {
              id: '000000000000000000001',
              name: '中村 太郎',
              nameKana: 'なかむら たろう',
              lessonTotal: 0,
            },
            {
              id: '000000000000000000002',
              name: '西山 幸子',
              nameKana: 'にしやま さちこ',
              lessonTotal: 2,
            },
            {
              id: '000000000000000000003',
              name: '鈴木 小太郎',
              nameKana: 'すずき こたろう',
              lessonTotal: 1,
            },
          ])
          expect(ShiftStore.getStudents).toEqual([
            {
              id: '100000000000000000001',
              name: '中村 太郎',
              nameKana: 'なかむら たろう',
              lessonTotal: 0,
              suggestedLessonsTotal: 0,
            },
            {
              id: '100000000000000000002',
              name: '西山 幸子',
              nameKana: 'にしやま さちこ',
              lessonTotal: 2,
              suggestedLessonsTotal: 0,
            },
          ])
          expect(ShiftStore.getLessons).toEqual([
            {
              id: 1,
              shiftId: 1,
              subjectId: 1,
              room: 1,
              teacherId: '000000000000000000001',
              studentId: '100000000000000000001',
              startAt: '2022-02-01T17:00:00+09:00',
              endAt: '2022-02-01T18:30:00+09:00',
              notes: '',
              createdAt: '2021-12-02T18:30:00+09:00',
              updatedAt: '2021-12-02T18:30:00+09:00',
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
          await expect(ShiftStore.listShiftDetails({ summaryId: 1 })).rejects.toThrow(err)
        })
      })
    })
  })
})
