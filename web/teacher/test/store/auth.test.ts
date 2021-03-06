import { setup, setSafetyMode, refresh } from '~~/test/helpers/store-helper'
import { AuthStore } from '~/store'
import { ApiError } from '~/types/exception'
import { ErrorResponse } from '~/types/api/exception'
import { Role } from '~/types/store'

const subjects = {
  小学校: [],
  中学校: [],
  高校: [
    {
      id: 1,
      name: '国語',
      fullname: '高校国語',
      color: '#F8BBD0',
      schoolType: '高校',
      createdAt: '2021-12-02T18:30:00+09:00',
      updatedAt: '2021-12-02T18:30:00+09:00',
    },
  ],
  その他: [],
}

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
        lastName: '',
        firstName: '',
        lastNameKana: '',
        firstNameKana: '',
        mail: '',
        role: Role.TEACHER,
        subjects: {
          小学校: [],
          中学校: [],
          高校: [],
          その他: [],
        },
      })
    })

    it('getRole', () => {
      expect(AuthStore.getRole).toBe(Role.TEACHER)
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
            lastName: '中村',
            firstName: '広大',
            lastNameKana: 'なかむら',
            firstNameKana: 'こうだい',
            mail: 'teacher-test001@calmato.jp',
            role: Role.TEACHER,
            subjects,
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
