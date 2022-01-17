import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/test/helpers/component-helper'
import TheSettingTop from '~/components/templates/TheSettingTop.vue'
import { Menu } from '~/types/props/setting'

describe('components/templates/TheSettingTop', () => {
  let wrapper: any
  beforeEach(() => {
    wrapper = shallowMount(TheSettingTop, {
      ...Options,
    })
  })

  describe('script', () => {
    describe('props', () => {
      describe('menuItems', () => {
        it('初期値', () => {
          expect(wrapper.props().menuItems).toEqual([])
        })

        it('値が代入されること', async () => {
          const items: Menu[] = [{ title: 'テスト', path: '/test' }]
          await wrapper.setProps({ menuItems: items })
          expect(wrapper.props().menuItems).toEqual(items)
        })
      })
    })

    describe('methods', () => {
      describe('onClick', () => {
        it('emitted', async () => {
          const item: Menu = { title: 'テスト', path: '/test' }
          await wrapper.vm.onClick(item)
          expect(wrapper.emitted('click')).toBeTruthy()
          expect(wrapper.emitted('click')[0][0]).toBe(item)
        })
      })
    })
  })
})
