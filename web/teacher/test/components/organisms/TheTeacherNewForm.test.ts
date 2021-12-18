import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/test/helpers/component-helper'
import TheTeacherNewForm from '~/components/organisms/TheTeacherNewForm.vue'
import { TeacherNewForm, TeacherNewOptions, TeacherNewParams } from '~/types/form'
import { Role } from '~/types/store'

describe('components/organisms/TheTeacherNewForm', () => {
  let wrapper: any
  beforeEach(() => {
    wrapper = shallowMount(TheTeacherNewForm, {
      ...Options,
    })
  })

  describe('script', () => {
    describe('props', () => {
      describe('form', () => {
        it('初期値', () => {
          expect(wrapper.props().form).toEqual({
            params: TeacherNewParams,
            options: TeacherNewOptions,
          })
        })

        it('値が代入されること', async () => {
          const form: TeacherNewForm = {
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
          await wrapper.setProps({ form })
          expect(wrapper.props().form).toBe(form)
        })
      })
    })

    describe('data', () => {
      it('roleItems', () => {
        expect(wrapper.vm.roleItems).toEqual([
          { text: '--- 役職を選択してください -------', value: Role.UNKNOWN },
          { text: '講師', value: Role.TEACHER },
          { text: '管理者', value: Role.ADMINISTRATOR },
        ])
      })
    })
  })
})
