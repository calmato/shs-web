import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/test/helpers/component-helper'
import Top from '~/components/templates/Top.vue'

describe('components/templates/SignIn', () => {
  let wrapper: any
  beforeEach(() => {
    wrapper = shallowMount(Top, {
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
