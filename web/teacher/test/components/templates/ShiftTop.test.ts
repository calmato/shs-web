import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/test/helpers/component-helper'
import ShiftTop from '~/components/templates/ShiftTop.vue'

describe('components/templates/ShiftTop', () => {
  let wrapper: any
  beforeEach(() => {
    wrapper = shallowMount(ShiftTop, {
      ...Options,
    })
  })

  it('test is none', () => {
    expect(wrapper).not.toBeNaN()
  })
})
