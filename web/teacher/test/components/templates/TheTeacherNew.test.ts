import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/test/helpers/component-helper'
import TheTeacherNew from '~/components/templates/TheTeacherNew.vue'
import { TeacherNewForm, TeacherNewOptions, TeacherNewParams } from '~/types/form'

describe('components/templates/TheTeacherNew', () => {
  let wrapper: any
  beforeEach(() => {
    wrapper = shallowMount(TheTeacherNew, {
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

      describe('loading', () => {
        it('初期値', () => {
          expect(wrapper.props().loading).toBeFalsy()
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ loading: true })
          expect(wrapper.props().loading).toBeTruthy()
        })
      })
    })

    describe('methods', () => {
      describe('onSubmit', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onSubmit()
          expect(wrapper.emitted('submit')).toBeTruthy()
        })
      })

      describe('onCancel', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onCancel()
          expect(wrapper.emitted('cancel')).toBeTruthy()
        })
      })
    })
  })
})
