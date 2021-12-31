import { setup, refresh, setSafetyMode } from '~~/test/helpers/store-helper'
import { ShiftStore } from '~/store'
import { ApiError } from '~/types/exception'
import { ErrorResponse } from '~/types/api/exception'
import { ShiftDetail, ShiftStatus } from '~/types/store'
import { ShiftsNewForm, ShiftsNewOptions } from '~/types/form'

describe('store/shift', () => {
  beforeEach(() => {
    setup()
  })

  afterEach(() => {
    refresh()
  })

  describe('getters', () => {
    it('getSummaries', () => {
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
      expect(ShiftStore.getDetails).toEqual(new Map<string, ShiftDetail>())
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
          expect(ShiftStore.getDetails).toEqual(
            new Map([
              [
                '20210201',
                {
                  isClosed: false,
                  lessons: [
                    { id: 1, startTime: '1700', endTime: '1830' },
                    { id: 2, startTime: '1830', endTime: '2000' },
                  ],
                },
              ],
              [
                '20210202',
                {
                  isClosed: true,
                  lessons: [],
                },
              ],
            ])
          )
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
  })
})
