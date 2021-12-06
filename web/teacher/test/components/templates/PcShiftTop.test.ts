import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/test/helpers/component-helper'
import PcShiftTop from '~/components/templates/PcShiftTop.vue'

describe('components/templates/PcShiftTop', () => {
  let wrapper: any
  beforeEach(() => {
    wrapper = shallowMount(PcShiftTop, {
      ...Options,
    })
  })

  it('test is none', () => {
    expect(wrapper).not.toBeNaN()
  })
})
