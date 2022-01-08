import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/test/helpers/component-helper'
import TheUserTop from '~/components/templates/TheUserTop.vue'
import { Role, SchoolType, Student, SubjectsMap, Teacher } from '~/types/store'
import {
  TeacherEditRoleForm,
  TeacherEditRoleOptions,
  TeacherEditRoleParams,
  TeacherEditSubjectForElementarySchoolOptions,
  TeacherEditSubjectForElementarySchoolParams,
  TeacherEditSubjectForHighSchoolOptions,
  TeacherEditSubjectForHighSchoolParams,
  TeacherEditSubjectForJuniorHighSchoolOptions,
  TeacherEditSubjectForJuniorHighSchoolParams,
  TeacherEditSubjectForm,
} from '~/types/form'

describe('components/templates/TheUserTop', () => {
  let wrapper: any
  beforeEach(() => {
    wrapper = shallowMount(TheUserTop, {
      ...Options,
    })
  })

  describe('script', () => {
    describe('props', () => {
      describe('loading', () => {
        it('初期値', () => {
          expect(wrapper.props().loading).toBeFalsy()
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ loading: true })
          expect(wrapper.props().loading).toBeTruthy()
        })
      })

      describe('isAdmin', () => {
        it('初期値', () => {
          expect(wrapper.props().isAdmin).toBeFalsy()
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ isAdmin: true })
          expect(wrapper.props().isAdmin).toBeTruthy()
        })
      })

      describe('subjects', () => {
        it('初期値', () => {
          expect(wrapper.props().subjects).toEqual({
            [SchoolType.ELEMENTARY_SCHOOL]: [],
            [SchoolType.JUNIOR_HIGH_SCHOOL]: [],
            [SchoolType.HIGH_SCHOOL]: [],
          })
        })
        it('値が代入されること', async () => {
          const subjects: SubjectsMap = {
            [SchoolType.ELEMENTARY_SCHOOL]: [
              {
                id: 1,
                name: '国語',
                color: '#F8BBD0',
                schoolType: SchoolType.ELEMENTARY_SCHOOL,
                createdAt: '',
                updatedAt: '',
              },
            ],
            [SchoolType.JUNIOR_HIGH_SCHOOL]: [
              {
                id: 2,
                name: '数学',
                color: '#BBDEFB',
                schoolType: SchoolType.JUNIOR_HIGH_SCHOOL,
                createdAt: '',
                updatedAt: '',
              },
            ],
            [SchoolType.HIGH_SCHOOL]: [
              {
                id: 3,
                name: '英語',
                color: '#FEE6C9',
                schoolType: SchoolType.HIGH_SCHOOL,
                createdAt: '',
                updatedAt: '',
              },
            ],
          }
          await wrapper.setProps({ subjects })
          expect(wrapper.props().subjects).toEqual({
            [SchoolType.ELEMENTARY_SCHOOL]: [
              {
                id: 1,
                name: '国語',
                color: '#F8BBD0',
                schoolType: SchoolType.ELEMENTARY_SCHOOL,
                createdAt: '',
                updatedAt: '',
              },
            ],
            [SchoolType.JUNIOR_HIGH_SCHOOL]: [
              {
                id: 2,
                name: '数学',
                color: '#BBDEFB',
                schoolType: SchoolType.JUNIOR_HIGH_SCHOOL,
                createdAt: '',
                updatedAt: '',
              },
            ],
            [SchoolType.HIGH_SCHOOL]: [
              {
                id: 3,
                name: '英語',
                color: '#FEE6C9',
                schoolType: SchoolType.HIGH_SCHOOL,
                createdAt: '',
                updatedAt: '',
              },
            ],
          })
        })
      })

      describe('students', () => {
        it('初期値', () => {
          expect(wrapper.props().students).toEqual([])
        })

        it('値が代入されること', async () => {
          const students: Student[] = [
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
          ]
          await wrapper.setProps({ students })
          expect(wrapper.props().students).toBe(students)
        })
      })

      describe('teachers', () => {
        it('初期値', () => {
          expect(wrapper.props().teacher).toEqual({
            id: '',
            lastName: '',
            firstName: '',
            lastNameKana: '',
            firstNameKana: '',
            mail: '',
            role: Role.TEACHER,
            subjects: {},
            createdAt: '',
            updatedAt: '',
          })
        })

        it('値が代入されること', async () => {
          const teacher: Teacher = {
            id: '000000000000000000001',
            name: '中村 太郎',
            nameKana: 'なかむら たろう',
            lastName: '中村',
            firstName: '太郎',
            lastNameKana: 'なかむら',
            firstNameKana: 'たろう',
            mail: 'teacher-001@calmato.jp',
            role: 0,
            subjects: {},
            createdAt: '',
            updatedAt: '',
          }
          await wrapper.setProps({ teacher })
          expect(wrapper.props().teacher).toBe(teacher)
        })
      })

      describe('teacherDialog', () => {
        it('初期値', () => {
          expect(wrapper.props().teacherDialog).toBeFalsy()
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ teacherDialog: true })
          expect(wrapper.props().teacherDialog).toBeTruthy()
        })
      })

      describe('teachers', () => {
        it('初期値', () => {
          expect(wrapper.props().teachers).toEqual([])
        })

        it('値が代入されること', async () => {
          const teachers: Teacher[] = [
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
              subjects: {},
              createdAt: '',
              updatedAt: '',
            },
          ]
          await wrapper.setProps({ teachers })
          expect(wrapper.props().teachers).toBe(teachers)
        })
      })

      describe('teachersTotal', () => {
        it('初期値', () => {
          expect(wrapper.props().teachersTotal).toBe(0)
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ teachersTotal: 100 })
          expect(wrapper.props().teachersTotal).toBe(100)
        })
      })

      describe('teachersPage', () => {
        it('初期値', () => {
          expect(wrapper.props().teachersPage).toBe(1)
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ teachersPage: 2 })
          expect(wrapper.props().teachersPage).toBe(2)
        })
      })

      describe('teachersItemsPerPage', () => {
        it('初期値', () => {
          expect(wrapper.props().teachersItemsPerPage).toBe(10)
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ teachersItemsPerPage: 50 })
          expect(wrapper.props().teachersItemsPerPage).toBe(50)
        })
      })

      describe('editTeacherElementarySchoolForm', () => {
        it('初期値', () => {
          expect(wrapper.props().editTeacherElementarySchoolForm).toEqual({
            params: TeacherEditSubjectForElementarySchoolParams,
            options: TeacherEditSubjectForElementarySchoolOptions,
          })
        })

        it('値が代入されること', async () => {
          const form: TeacherEditSubjectForm = {
            params: {
              schoolType: SchoolType.ELEMENTARY_SCHOOL,
              subjectIds: [1, 2],
            },
            options: TeacherEditSubjectForElementarySchoolOptions,
          }
          await wrapper.setProps({ editTeacherElementarySchoolForm: form })
          expect(wrapper.props().editTeacherElementarySchoolForm).toBe(form)
        })
      })

      describe('editTeacherJuniorHighSchoolForm', () => {
        it('初期値', () => {
          expect(wrapper.props().editTeacherJuniorHighSchoolForm).toEqual({
            params: TeacherEditSubjectForJuniorHighSchoolParams,
            options: TeacherEditSubjectForJuniorHighSchoolOptions,
          })
        })

        it('値が代入されること', async () => {
          const form: TeacherEditSubjectForm = {
            params: {
              schoolType: SchoolType.JUNIOR_HIGH_SCHOOL,
              subjectIds: [1, 2],
            },
            options: TeacherEditSubjectForJuniorHighSchoolOptions,
          }
          await wrapper.setProps({ editTeacherJuniorHighSchoolForm: form })
          expect(wrapper.props().editTeacherJuniorHighSchoolForm).toBe(form)
        })
      })

      describe('editTeacherHighSchoolForm', () => {
        it('初期値', () => {
          expect(wrapper.props().editTeacherHighSchoolForm).toEqual({
            params: TeacherEditSubjectForHighSchoolParams,
            options: TeacherEditSubjectForHighSchoolOptions,
          })
        })

        it('値が代入されること', async () => {
          const form: TeacherEditSubjectForm = {
            params: {
              schoolType: SchoolType.HIGH_SCHOOL,
              subjectIds: [1, 2],
            },
            options: TeacherEditSubjectForHighSchoolOptions,
          }
          await wrapper.setProps({ editTeacherHighSchoolForm: form })
          expect(wrapper.props().editTeacherHighSchoolForm).toBe(form)
        })
      })

      describe('editTeacherRoleForm', () => {
        it('初期値', () => {
          expect(wrapper.props().editTeacherRoleForm).toEqual({
            params: TeacherEditRoleParams,
            options: TeacherEditRoleOptions,
          })
        })

        it('値が代入されること', async () => {
          const form: TeacherEditRoleForm = {
            params: {
              role: Role.ADMINISTRATOR,
            },
            options: TeacherEditRoleOptions,
          }
          await wrapper.setProps({ editTeacherRoleForm: form })
          expect(wrapper.props().editTeacherRoleForm).toBe(form)
        })
      })
    })

    describe('data', () => {
      it('actors', () => {
        expect(wrapper.vm.actors).toEqual([
          { name: '講師', value: 'teachers' },
          { name: '生徒', value: 'students' },
        ])
      })

      it('selector', () => {
        expect(wrapper.vm.selector).toBe('teachers')
      })
    })

    describe('methods', () => {
      describe('onClickNew', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onClickNew('teachers')
          expect(wrapper.emitted('click:new')).toBeTruthy()
          expect(wrapper.emitted('click:new')[0][0]).toBe('teachers')
        })
      })

      describe('onClickShowTeacher', () => {
        it('emitが実行されること', async () => {
          const teacher: Teacher = {
            id: '000000000000000000001',
            name: '中村 太郎',
            nameKana: 'なかむら たろう',
            lastName: '中村',
            firstName: '太郎',
            lastNameKana: 'なかむら',
            firstNameKana: 'たろう',
            mail: 'teacher-001@calmato.jp',
            role: 0,
            subjects: {},
            createdAt: '',
            updatedAt: '',
          }
          await wrapper.vm.onClickShowTeacher(teacher)
          expect(wrapper.emitted('click:show-teacher')).toBeTruthy()
          expect(wrapper.emitted('click:show-teacher')[0][0]).toBe(teacher)
        })
      })

      describe('onCloseTeacherDialog', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onCloseTeacherDialog()
          expect(wrapper.emitted('click:close-teacher')).toBeTruthy()
        })
      })

      describe('onSubmitTeacherElementarySchool', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onSubmitTeacherElementarySchool()
          expect(wrapper.emitted('submit:teacher-elementary-school')).toBeTruthy()
        })
      })

      describe('onSubmitTeacherJuniorHighSchool', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onSubmitTeacherJuniorHighSchool()
          expect(wrapper.emitted('submit:teacher-junior-high-school')).toBeTruthy()
        })
      })

      describe('onSubmitTeacherHighSchool', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onSubmitTeacherHighSchool()
          expect(wrapper.emitted('submit:teacher-high-school')).toBeTruthy()
        })
      })

      describe('onSubmitTeacherRole', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onSubmitTeacherRole()
          expect(wrapper.emitted('submit:teacher-role')).toBeTruthy()
        })
      })
    })
  })
})
