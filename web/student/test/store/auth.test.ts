import { setup, setSafetyMode, refresh } from '~~/test/helpers/store-helper'
import { AuthStore } from '~/store'
import { ApiError } from '~/types/exception'
import { ErrorResponse } from '~/types/api/exception'

describe('store/auth', () => {
  beforeEach(() => {
    setup()
  })

  afterEach(() => {
    refresh()
  })

  describe('getters', () => {
    it('getUid', () => {
      expect(AuthStore.getUid).toBe('')
    })

    it('getToken', () => {
      expect(AuthStore.getToken).toBe('')
    })

    it('getEmailVerified', () => {
      expect(AuthStore.getEmailVerified).toBeFalsy()
    })

    it('getAuth', () => {
      expect(AuthStore.getAuth).toEqual({
        id: '',
        name: '',
        nameKana: '',
        lastName: '',
        firstName: '',
        lastNameKana: '',
        firstNameKana: '',
        mail: '',
        schoolType: 'その他',
        grade: 0,
        subjects: [],
      })
    })
  })

  describe('actions', () => {
    describe('showAuth', () => {
      describe('success', () => {
        beforeEach(() => {
          setSafetyMode(true)
        })

        it('stateが更新されていること', async () => {
          await AuthStore.showAuth()
          expect(AuthStore.getAuth).toEqual({
            id: 'kSByoE6FetnPs5Byk3a9Zx',
            name: '中村 広大',
            nameKana: 'なかむら こうだい',
            lastName: '中村',
            firstName: '広大',
            lastNameKana: 'なかむら',
            firstNameKana: 'こうだい',
            mail: 'teacher-test001@calmato.jp',
            schoolType: '高校',
            grade: 2,
            subjects: [],
          })
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
          await expect(AuthStore.showAuth()).rejects.toThrow(err)
        })
      })
    })
  })
})
