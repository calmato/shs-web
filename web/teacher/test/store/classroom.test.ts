import { refresh, setIsAxiosMockValue, setSafetyMode, setup } from '~~/test//helpers/store-helper'
import { ClassroomStore } from '~/store'
import { ErrorResponse } from '~/types/api/exception'
import { ApiError } from '~/types/exception'
import { Schedule } from '~/types/api/v1'
import { HourForm } from '~/types/form'

describe('store/classroom', () => {
  beforeEach(() => {
    setup()
  })

  afterEach(() => {
    refresh()
  })

  describe('getters', () => {
    it('getTotalRooms', () => {
      expect(ClassroomStore.getTotalRooms).toBe(0)
    })

    it('getSchedules', () => {
      expect(ClassroomStore.getSchedules).toEqual([])
    })
    it('weekdayHourFormValue', () => {
      expect(ClassroomStore.weekdayHourFormValue).toEqual([])
    })

    it('holidayHourFormValue', () => {
      expect(ClassroomStore.holidayHourFormValue).toEqual([])
    })

    it('regularHoliday', () => {
      expect(ClassroomStore.regularHoliday).toEqual([])
    })
  })

  describe('actions', () => {
    describe('getTotalRoomsByApi', () => {
      describe('success', () => {
        beforeEach(() => {
          setSafetyMode(true)
        })

        it('return resolve', async () => {
          await expect(ClassroomStore.getTotalRoomsByApi()).resolves.toBeUndefined()
        })

        it('change state', async () => {
          await ClassroomStore.getTotalRoomsByApi()
          expect(ClassroomStore.getTotalRooms).toBe(5)
        })
      })

      describe('failure', () => {
        beforeEach(() => {
          setSafetyMode(false)
          setIsAxiosMockValue(true)
        })

        it('return reject', async () => {
          const err = new ApiError(503, 'api error', {
            status: 503,
            message: 'api error',
            details: 'some error',
          } as ErrorResponse)
          try {
            await ClassroomStore.getTotalRoomsByApi()
          } catch (e) {
            expect(e).toEqual(err)
          }
        })

        it('throw internal server error', async () => {
          setIsAxiosMockValue(false)

          const err = new Error('internal server error')

          try {
            await ClassroomStore.getTotalRoomsByApi()
          } catch (e) {
            expect(e).toEqual(err)
          }
        })
      })
    })

    describe('getSchedulesByApi', () => {
      describe('success', () => {
        beforeEach(() => {
          setSafetyMode(true)
        })

        it('return resolve', async () => {
          await expect(ClassroomStore.getSchedulesByApi()).resolves.toBeUndefined()
        })

        it('change state', async () => {
          const expectedSchedules: Schedule[] = [
            {
              weekday: 0,
              isClosed: false,
              lessons: [
                { startTime: '1530', endTime: '1700' },
                { startTime: '1700', endTime: '1830' },
                { startTime: '1830', endTime: '2000' },
                { startTime: '2000', endTime: '2130' },
              ],
            },
            {
              weekday: 1,
              isClosed: false,
              lessons: [
                { startTime: '1700', endTime: '1830' },
                { startTime: '1830', endTime: '2000' },
                { startTime: '2000', endTime: '2130' },
              ],
            },
            {
              weekday: 2,
              isClosed: false,
              lessons: [
                { startTime: '1700', endTime: '1830' },
                { startTime: '1830', endTime: '2000' },
                { startTime: '2000', endTime: '2130' },
              ],
            },
            {
              weekday: 3,
              isClosed: false,
              lessons: [
                { startTime: '1700', endTime: '1830' },
                { startTime: '1830', endTime: '2000' },
                { startTime: '2000', endTime: '2130' },
              ],
            },
            {
              weekday: 4,
              isClosed: false,
              lessons: [
                { startTime: '1700', endTime: '1830' },
                { startTime: '1830', endTime: '2000' },
                { startTime: '2000', endTime: '2130' },
              ],
            },
            {
              weekday: 5,
              isClosed: false,
              lessons: [
                { startTime: '1700', endTime: '1830' },
                { startTime: '1830', endTime: '2000' },
                { startTime: '2000', endTime: '2130' },
              ],
            },
            {
              weekday: 6,
              isClosed: false,
              lessons: [
                { startTime: '1530', endTime: '1700' },
                { startTime: '1700', endTime: '1830' },
                { startTime: '1830', endTime: '2000' },
                { startTime: '2000', endTime: '2130' },
              ],
            },
          ]

          const expectedWeekdayHourFormValue: HourForm[] = [
            { startAt: '17:00', endAt: '18:30' },
            { startAt: '18:30', endAt: '20:00' },
            { startAt: '20:00', endAt: '21:30' },
          ]

          const expectedHolidayHourFormValue: HourForm[] = [
            { startAt: '15:30', endAt: '17:00' },
            { startAt: '17:00', endAt: '18:30' },
            { startAt: '18:30', endAt: '20:00' },
            { startAt: '20:00', endAt: '21:30' },
          ]

          const expectedRegularHoliday: number[] = []

          await ClassroomStore.getSchedulesByApi()
          expect(ClassroomStore.getSchedules).toEqual(expectedSchedules)
          expect(ClassroomStore.weekdayHourFormValue.length).toBe(expectedWeekdayHourFormValue.length)
          expect(ClassroomStore.holidayHourFormValue.length).toEqual(expectedHolidayHourFormValue.length)
          expect(ClassroomStore.regularHoliday).toEqual(expectedRegularHoliday)
        })
      })

      describe('failure', () => {
        beforeEach(() => {
          setSafetyMode(false)
          setIsAxiosMockValue(true)
        })

        it('return reject', async () => {
          const err = new ApiError(503, 'api error', {
            status: 503,
            message: 'api error',
            details: 'some error',
          } as ErrorResponse)
          try {
            await ClassroomStore.getSchedulesByApi()
          } catch (e) {
            expect(e).toEqual(err)
          }
        })

        it('throw internal server error', async () => {
          setIsAxiosMockValue(false)

          const err = new Error('internal server error')

          try {
            await ClassroomStore.getSchedulesByApi()
          } catch (e) {
            expect(e).toEqual(err)
          }
        })
      })
    })
  })
})
