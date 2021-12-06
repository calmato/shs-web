import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/test/helpers/component-helper'
import SettingTop from '~/components/templates/SettingTop.vue'
import { Menu } from '~/types/props/setting'

describe('components/templates/SettingTop', () => {
  let wrapper: any
  beforeEach(() => {
    wrapper = shallowMount(SettingTop, {
      ...Options,
    })
  })

  describe('script', () => {
    describe('props', () => {
      describe('userItems', () => {
        it('初期値', () => {
          expect(wrapper.props().userItems).toEqual([])
        })

        it('値が代入されること', async () => {
          const items: Menu[] = [{ title: 'テスト', path: '/test' }]
          await wrapper.setProps({ userItems: items })
          expect(wrapper.props().userItems).toEqual(items)
        })
      })

      describe('systemItems', () => {
        it('初期値', () => {
          expect(wrapper.props().systemItems).toEqual([])
        })

        it('値が代入されること', async () => {
          const items: Menu[] = [{ title: 'テスト', path: '/test' }]
          await wrapper.setProps({ systemItems: items })
          expect(wrapper.props().systemItems).toEqual(items)
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
