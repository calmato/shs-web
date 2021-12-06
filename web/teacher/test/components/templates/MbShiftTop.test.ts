import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/test/helpers/component-helper'
import MbShiftTop from '~/components/templates/MbShiftTop.vue'

describe('components/templates/MbShiftTop', () => {
  let wrapper: any
  beforeEach(() => {
    wrapper = shallowMount(MbShiftTop, {
      ...Options,
    })
  })

  describe('script', () => {
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
