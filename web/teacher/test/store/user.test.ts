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

    it('getStudents', () => {
      expect(UserStore.getStudents).toEqual([
        {
          id: '123456789012345678901',
          name: '浜田 二郎',
          nameKana: 'はまだ じろう',
          lastName: '浜田',
          firstName: '二郎',
          lastNameKana: 'はまだ',
          firstNameKana: 'じろう',
          mail: 'student-001@calmato.jp',
          type: 1,
          grade: 2,
          createdAt: '',
          updatedAt: '',
        },
      ])
    })

    it('getStudentMap', () => {
      expect(UserStore.getStudentMap).toEqual({
        '123456789012345678901': {
          id: '123456789012345678901',
          name: '浜田 二郎',
          nameKana: 'はまだ じろう',
          lastName: '浜田',
          firstName: '二郎',
          lastNameKana: 'はまだ',
          firstNameKana: 'じろう',
          mail: 'student-001@calmato.jp',
          type: 1,
          grade: 2,
          createdAt: '',
          updatedAt: '',
        },
      })
    })

    it('getTeachers', () => {
      expect(UserStore.getTeachers).toEqual([
        {
          id: '000000000000000000001',
          name: '中村 太郎',
          nameKana: 'なかむら たろう',
          lastName: '中村',
          firstName: '太郎',
          lastNameKana: 'なかむら',
          firstNameKana: 'たろう',
          mail: 'teacher-001@calmato.jp',
          role: 0,
          createdAt: '',
          updatedAt: '',
        },
        {
          id: '000000000000000000002',
          name: '西山 幸子',
          nameKana: 'にしやま さちこ',
          lastName: '西山',
          firstName: '幸子',
          lastNameKana: 'にしやま',
          firstNameKana: 'さちこ',
          mail: 'teacher-002@calmato.jp',
          role: 0,
          createdAt: '',
          updatedAt: '',
        },
        {
          id: '000000000000000000003',
          name: '鈴木 小太郎',
          nameKana: 'すずき こたろう',
          lastName: '鈴木',
          firstName: '小太郎',
          lastNameKana: 'すずき',
          firstNameKana: 'こたろう',
          mail: 'teacher-003@calmato.jp',
          role: 1,
          createdAt: '',
          updatedAt: '',
        },
      ])
    })

    it('getTeacherMap', () => {
      expect(UserStore.getTeacherMap).toEqual({
        '000000000000000000001': {
          id: '000000000000000000001',
          name: '中村 太郎',
          nameKana: 'なかむら たろう',
          lastName: '中村',
          firstName: '太郎',
          lastNameKana: 'なかむら',
          firstNameKana: 'たろう',
          mail: 'teacher-001@calmato.jp',
          role: 0,
          createdAt: '',
          updatedAt: '',
        },
        '000000000000000000002': {
          id: '000000000000000000002',
          name: '西山 幸子',
          nameKana: 'にしやま さちこ',
          lastName: '西山',
          firstName: '幸子',
          lastNameKana: 'にしやま',
          firstNameKana: 'さちこ',
          mail: 'teacher-002@calmato.jp',
          role: 0,
          createdAt: '',
          updatedAt: '',
        },
        '000000000000000000003': {
          id: '000000000000000000003',
          name: '鈴木 小太郎',
          nameKana: 'すずき こたろう',
          lastName: '鈴木',
          firstName: '小太郎',
          lastNameKana: 'すずき',
          firstNameKana: 'こたろう',
          mail: 'teacher-003@calmato.jp',
          role: 1,
          createdAt: '',
          updatedAt: '',
        },
      })
    })
  })

  describe('actions', () => {
    describe('hello', () => {
      describe('success', () => {
        beforeEach(() => {
          setSafetyMode(true)
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
