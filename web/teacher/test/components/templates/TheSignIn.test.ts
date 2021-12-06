import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/test/helpers/component-helper'
import TheSignIn from '~/components/templates/TheSignIn.vue'

describe('components/templates/TheSignIn', () => {
  let wrapper: any
  beforeEach(() => {
    wrapper = shallowMount(TheSignIn, {
      ...Options,
    })
  })

  describe('script', () => {
    describe('props', () => {
      describe('message', () => {
        it('初期値', () => {
          expect(wrapper.props().message).toEqual('')
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ message: 'test message' })
          expect(wrapper.props().message).toEqual('test message')
        })
      })
    })

    describe('methods', () => {
      describe('onClick', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onClick()
          expect(wrapper.emitted('click')).toBeTruthy()
        })
      })
    })
  })
})
