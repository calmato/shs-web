import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/test/helpers/component-helper'
import TheHeader from '~/components/organisms/TheHeader.vue'

describe('components/organisms/TheHeader', () => {
  let wrapper: any
  beforeEach(() => {
    wrapper = shallowMount(TheHeader, {
      ...Options,
    })
  })

  describe('script', () => {
    describe('props', () => {
      describe('overlay', () => {
        it('初期値', () => {
          expect(wrapper.props().overlay).toBeFalsy()
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ overlay: true })
          expect(wrapper.props().overlay).toBeTruthy()
        })
      })
    })

    describe('methods', () => {
      describe('onClick', () => {
        it('emitted', async () => {
          await wrapper.vm.onClick()
          expect(wrapper.emitted('click')).toBeTruthy()
        })
      })
    })
  })
})
