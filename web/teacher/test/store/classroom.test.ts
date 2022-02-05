import { async } from '@firebase/util'
import { refresh, setSafetyMode, setup } from '~~/test//helpers/store-helper'
import { ClassroomStore } from '~/store'
import { ErrorResponse } from '~/types/api/exception'
import { ApiError } from '~/types/exception'

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
      })
    })
  })
})
