import { setup, refresh, setSafetyMode } from '~~/test/helpers/store-helper'
import { UserStore } from '~/store'
import { ApiError } from '~/types/exception'
import { ErrorResponse } from '~/types/api/exception'
import {
  TeacherEditRoleForm,
  TeacherEditRoleOptions,
  TeacherEditSubjectForHighSchoolOptions,
  TeacherEditSubjectForm,
  TeacherNewForm,
  TeacherNewOptions,
} from '~/types/form'
import { Role } from '~/types/store'

describe('store/user', () => {
  beforeEach(() => {
    setup()
  })

  afterEach(() => {
    refresh()
  })

  describe('getters', () => {
    it('getStudents', () => {
      expect(UserStore.getStudents).toEqual([])
    })

    it('getStudentMap', () => {
      expect(UserStore.getStudentMap).toEqual({})
    })

    it('getTeacher', () => {
      expect(UserStore.getTeacher).toEqual({
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
        createdAt: '',
        updatedAt: '',
      })
    })

    it('getTeachers', () => {
      expect(UserStore.getTeachers).toEqual([])
    })

    it('getTeacherMap', () => {
      expect(UserStore.getTeacherMap).toEqual({})
    })

    it('getTeachersTotal', () => {
      expect(UserStore.getTeachersTotal).toBe(0)
    })
  })

  describe('actions', () => {
    describe('listTeachers', () => {
      describe('success', () => {
        beforeEach(() => {
          setSafetyMode(true)
        })

        it('changing state when limit 0 and offset 0', async () => {
          await UserStore.listTeachers({ limit: 0, offset: 0 })
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
              role: 1,
              subjects: {
                小学校: [],
                中学校: [],
                高校: [],
                その他: [],
              },
              createdAt: '2021-12-02T18:30:00+09:00',
              updatedAt: '2021-12-02T18:30:00+09:00',
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
              role: 1,
              subjects: {
                小学校: [],
                中学校: [],
                高校: [],
                その他: [],
              },
              createdAt: '2021-12-02T18:30:00+09:00',
              updatedAt: '2021-12-02T18:30:00+09:00',
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
              role: 2,
              subjects: {
                小学校: [],
                中学校: [],
                高校: [],
                その他: [],
              },
              createdAt: '2021-12-02T18:30:00+09:00',
              updatedAt: '2021-12-02T18:30:00+09:00',
            },
          ])
          expect(UserStore.getTeachersTotal).toBe(3)
        })

        it('changing state when limit 20 and offset 2', async () => {
          await UserStore.listTeachers({ limit: 20, offset: 2 })
          expect(UserStore.getTeachers).toEqual([
            {
              id: '000000000000000000003',
              name: '鈴木 小太郎',
              nameKana: 'すずき こたろう',
              lastName: '鈴木',
              firstName: '小太郎',
              lastNameKana: 'すずき',
              firstNameKana: 'こたろう',
              mail: 'teacher-003@calmato.jp',
              role: 2,
              subjects: {
                小学校: [],
                中学校: [],
                高校: [],
                その他: [],
              },
              createdAt: '2021-12-02T18:30:00+09:00',
              updatedAt: '2021-12-02T18:30:00+09:00',
            },
          ])
          expect(UserStore.getTeachersTotal).toBe(3)
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
          await expect(UserStore.listTeachers({ limit: 0, offset: 0 })).rejects.toThrow(err)
        })
      })
    })

    describe('showTeacher', () => {
      describe('success', () => {
        beforeEach(() => {
          setSafetyMode(true)
        })

        it('changing state', async () => {
          await UserStore.showTeacher({ teacherId: '000000000000000000001' })
          expect(UserStore.getTeacher).toEqual({
            id: '000000000000000000001',
            name: '中村 太郎',
            nameKana: 'なかむら たろう',
            lastName: '中村',
            firstName: '太郎',
            lastNameKana: 'なかむら',
            firstNameKana: 'たろう',
            mail: 'teacher-001@calmato.jp',
            role: 1,
            subjects: {
              小学校: [
                {
                  id: 1,
                  name: '国語',
                  fullname: '小学校国語',
                  color: '#F8BBD0',
                  schoolType: '小学校',
                  createdAt: '',
                  updatedAt: '',
                },
              ],
              中学校: [
                {
                  id: 2,
                  name: '数学',
                  fullname: '中学校数学',
                  color: '#BBDEFB',
                  schoolType: '中学校',
                  createdAt: '',
                  updatedAt: '',
                },
              ],
              高校: [
                {
                  id: 3,
                  name: '英語',
                  fullname: '高校英語',
                  color: '#FEE6C9',
                  schoolType: '高校',
                  createdAt: '',
                  updatedAt: '',
                },
              ],
              その他: [],
            },
            createdAt: '2021-12-02T18:30:00+09:00',
            updatedAt: '2021-12-02T18:30:00+09:00',
          })
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
          await expect(UserStore.showTeacher({ teacherId: '000000000000000000001' })).rejects.toThrow(err)
        })
      })
    })

    describe('createTeacher', () => {
      let form: TeacherNewForm
      beforeEach(() => {
        form = {
          params: {
            lastName: '中村',
            firstName: '太郎',
            lastNameKana: 'なかむら',
            firstNameKana: 'たろう',
            mail: 'teacher-001@calmato.jp',
            password: '12345678',
            passwordConfirmation: '12345678',
            role: 1,
          },
          options: TeacherNewOptions,
        }
      })

      describe('success', () => {
        beforeEach(() => {
          setSafetyMode(true)
        })

        it('changing state', async () => {
          await UserStore.createTeacher({ form })
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
              role: 1,
              subjects: {
                小学校: [
                  {
                    id: 1,
                    name: '国語',
                    fullname: '小学校国語',
                    color: '#F8BBD0',
                    schoolType: '小学校',
                    createdAt: '',
                    updatedAt: '',
                  },
                ],
                中学校: [
                  {
                    id: 2,
                    name: '数学',
                    fullname: '中学校数学',
                    color: '#BBDEFB',
                    schoolType: '中学校',
                    createdAt: '',
                    updatedAt: '',
                  },
                ],
                高校: [
                  {
                    id: 3,
                    name: '英語',
                    fullname: '高校英語',
                    color: '#FEE6C9',
                    schoolType: '高校',
                    createdAt: '',
                    updatedAt: '',
                  },
                ],
                その他: [],
              },
              createdAt: '2021-12-02T18:30:00+09:00',
              updatedAt: '2021-12-02T18:30:00+09:00',
            },
          ])
          expect(UserStore.getTeachersTotal).toBe(1)
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
          await expect(UserStore.createTeacher({ form })).rejects.toThrow(err)
        })
      })
    })

    describe('updateTeacherSubjects', () => {
      let form: TeacherEditSubjectForm
      beforeEach(() => {
        form = {
          params: {
            schoolType: 3,
            subjectIds: [1, 2],
          },
          options: TeacherEditSubjectForHighSchoolOptions,
        }
      })

      describe('success', () => {
        beforeEach(() => {
          setSafetyMode(true)
        })

        it('return resolve', async () => {
          await expect(
            UserStore.updateTeacherSubjects({ teacherId: '000000000000000000001', form })
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
          await expect(UserStore.updateTeacherSubjects({ teacherId: '000000000000000000001', form })).rejects.toThrow(
            err
          )
        })
      })
    })

    describe('updateTeacherRole', () => {
      let form: TeacherEditRoleForm
      beforeEach(() => {
        form = {
          params: { role: Role.ADMINISTRATOR },
          options: TeacherEditRoleOptions,
        }
      })

      describe('success', () => {
        beforeEach(() => {
          setSafetyMode(true)
        })

        it('return resolve', async () => {
          await expect(
            UserStore.updateTeacherRole({ teacherId: '000000000000000000001', form })
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
          await expect(UserStore.updateTeacherRole({ teacherId: '000000000000000000001', form })).rejects.toThrow(err)
        })
      })
    })

    describe('deleteTeacher', () => {
      describe('success', () => {
        beforeEach(() => {
          setSafetyMode(true)
        })

        it('return resolve', async () => {
          await expect(UserStore.deleteTeacher({ teacherId: '000000000000000000001' })).resolves.toBeUndefined()
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
          await expect(UserStore.deleteTeacher({ teacherId: '000000000000000000001' })).rejects.toThrow(err)
        })
      })
    })
  })
})
