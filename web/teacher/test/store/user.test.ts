import { setup, setSafetyMode, refresh } from '~~/test/helpers/store-helper'
import { UserStore } from '~/store'
import { ApiError } from '~/types/exception'
import { ErrorResponse } from '~/types/api/exception'

describe('store/user', () => {
  beforeEach(() => {
    setup()
  })

  afterEach(() => {
    refresh()
  })

  describe('getters', () => {
    it('getMessage', () => {
      expect(UserStore.getMessage).toBe('')
    })
  })

  describe('actions', () => {
    describe('hello', () => {
      describe('success', () => {
        beforeEach(() => {
          setSafetyMode(true)
        })

        it('resolveが返されること', async () => {
          await expect(UserStore.hello()).resolves.toBeUndefined()
        })

        it('stateが更新されていること', async () => {
          await UserStore.hello()
          expect(UserStore.getMessage).toBe('test message')
        })
      })

      describe('failure', () => {
        beforeEach(() => {
          setSafetyMode(false)
        })

        it('rejectが返されること', async () => {
          const err = new ApiError(400, 'api error', {
            status: 400,
            message: 'api error',
            details: 'some error',
          } as ErrorResponse)
          await expect(UserStore.hello()).rejects.toThrow(err)
        })
      })
    })
  })
})
