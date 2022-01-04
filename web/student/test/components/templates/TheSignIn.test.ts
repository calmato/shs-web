import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/test/helpers/component-helper'
import TheSignIn from '~/components/templates/TheSignIn.vue'
import { SignInForm } from '~/types/form'

describe('components/templates/TheSignIn', () => {
  let wrapper: any
  beforeEach(() => {
    wrapper = shallowMount(TheSignIn, {
      ...Options,
    })
  })

  describe('script', () => {
    describe('props', () => {
      describe('form', () => {
        it('初期値', () => {
          expect(wrapper.props().form).toEqual({ mail: '', password: '' })
        })

        it('値だ代入されること', async () => {
          const form: SignInForm = {
            mail: 'teacher-test01@calmato.jp',
            password: '12345678',
          }
          await wrapper.setProps({ form })
          expect(wrapper.props().form).toBe(form)
        })
      })

      describe('hasError', () => {
        it('初期値', () => {
          expect(wrapper.props().hasError).toBeFalsy()
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ hasError: true })
          expect(wrapper.props().hasError).toBeTruthy()
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

    describe('computed', () => {
      describe('signInForm', () => {
        it('getter', () => {
          expect(wrapper.vm.signInForm).toEqual({ mail: '', password: '' })
        })
      })
    })

    describe('methods', () => {
      describe('onSubmit', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onSubmit()
          expect(wrapper.emitted('click')).toBeTruthy()
        })
      })
    })
  })
})
