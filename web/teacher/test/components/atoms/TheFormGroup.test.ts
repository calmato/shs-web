import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/test/helpers/component-helper'
import TheFormGroup from '~/components/atoms/TheFormGroup.vue'

describe('components/atoms/TheFormGroup', () => {
  let wrapper: any
  beforeEach(() => {
    wrapper = shallowMount(TheFormGroup, {
      ...Options,
    })
  })

  it('test case is empty', () => {
    expect(wrapper).not.toBeUndefined()
  })
})
