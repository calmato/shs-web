import { mount } from '@vue/test-utils'
import * as Options from '~~/test/helpers/component-helper'
import TheTeacherEditCard from '~/components/organisms/TheTeacherEditCard.vue'
import { Role, SubjectsMap, Teacher } from '~/types/store'
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

describe('components/organisms/TheTeacherEditCard', () => {
  let wrapper: any
  beforeEach(() => {
    wrapper = mount(TheTeacherEditCard, {
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

      describe('deleteDialog', () => {
        it('初期値', () => {
          expect(wrapper.props().deleteDialog).toBeFalsy()
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ deleteDialog: true })
          expect(wrapper.props().deleteDialog).toBeTruthy()
        })
      })

      describe('teacher', () => {
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
            subjects: {
              小学校: [],
              中学校: [],
              高校: [],
              その他: [],
            },
            createdAt: '',
            updatedAt: '',
          }
          await wrapper.setProps({ teacher })
          expect(wrapper.props().teacher).toBe(teacher)
        })
      })

      describe('subjects', () => {
        it('初期値', () => {
          expect(wrapper.props().subjects).toEqual({
            小学校: [],
            中学校: [],
            高校: [],
            その他: [],
          })
        })

        it('値が代入されること', async () => {
          const subjects: SubjectsMap = {
            小学校: [],
            中学校: [],
            高校: [],
            その他: [],
          }
          await wrapper.setProps({ subjects })
          expect(wrapper.props().subjects).toBe(subjects)
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
              schoolType: 1,
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
              schoolType: 2,
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
              schoolType: 3,
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
      it('roleItems', () => {
        expect(wrapper.vm.roleItems).toEqual([
          { text: '講師', value: Role.TEACHER },
          { text: '管理者', value: Role.ADMINISTRATOR },
        ])
      })
    })

    describe('methods', () => {
      describe('getTeacherName', () => {
        it('return', async () => {
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
            subjects: {
              小学校: [],
              中学校: [],
              高校: [],
              その他: [],
            },
            createdAt: '',
            updatedAt: '',
          }
          await wrapper.setProps({ teacher })
          expect(wrapper.vm.getTeacherName()).toBe('中村 太郎 (なかむら たろう)')
        })
      })

      describe('getElementarySchoolSubjects', () => {
        it('return', async () => {
          const subjects: SubjectsMap = {
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
          }
          await wrapper.setProps({ subjects })
          expect(wrapper.vm.getElementarySchoolSubjects()).toEqual([
            {
              id: 1,
              name: '国語',
              color: '#F8BBD0',
              schoolType: '小学校',
              createdAt: '',
              updatedAt: '',
            },
          ])
        })
      })

      describe('getJuniorHighSchoolSubjects', () => {
        it('return', async () => {
          const subjects: SubjectsMap = {
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
          }
          await wrapper.setProps({ subjects })
          expect(wrapper.vm.getJuniorHighSchoolSubjects()).toEqual([
            {
              id: 2,
              name: '数学',
              color: '#BBDEFB',
              schoolType: '中学校',
              createdAt: '',
              updatedAt: '',
            },
          ])
        })
      })

      describe('getHighSchoolSubjects', () => {
        it('return', async () => {
          const subjects: SubjectsMap = {
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
          }
          await wrapper.setProps({ subjects })
          expect(wrapper.vm.getHighSchoolSubjects()).toEqual([
            {
              id: 3,
              name: '英語',
              color: '#FEE6C9',
              schoolType: '高校',
              createdAt: '',
              updatedAt: '',
            },
          ])
        })
      })

      describe('onClose', () => {
        it('emitted', async () => {
          await wrapper.vm.onClose()
          expect(wrapper.emitted('click:close')).toBeTruthy()
        })
      })

      describe('onDelete', () => {
        it('emitted', async () => {
          await wrapper.vm.onDelete()
          expect(wrapper.emitted('click:delete')).toBeTruthy()
        })
      })

      describe('onDeleteAccept', () => {
        it('emitted', async () => {
          await wrapper.vm.onDeleteAccept()
          expect(wrapper.emitted('click:delete-accept')).toBeTruthy()
        })
      })

      describe('onDeleteCancel', () => {
        it('emitted', async () => {
          await wrapper.vm.onDeleteCancel()
          expect(wrapper.emitted('click:delete-cancel')).toBeTruthy()
        })
      })

      describe('onSubmitRole', () => {
        it('emitted', async () => {
          await wrapper.vm.onSubmitRole()
          expect(wrapper.emitted('submit:role')).toBeTruthy()
        })
      })

      describe('onSubmitElemntarySchool', () => {
        it('emitted', async () => {
          await wrapper.vm.onSubmitElementarySchool()
          expect(wrapper.emitted('submit:elementary-school')).toBeTruthy()
        })
      })

      describe('onSubmitJuniorHighSchool', () => {
        it('emitted', async () => {
          await wrapper.vm.onSubmitJuniorHighSchool()
          expect(wrapper.emitted('submit:junior-high-school')).toBeTruthy()
        })
      })

      describe('onSubmitHighSchool', () => {
        it('emitted', async () => {
          await wrapper.vm.onSubmitHighSchool()
          expect(wrapper.emitted('submit:high-school')).toBeTruthy()
        })
      })
    })
  })
})
