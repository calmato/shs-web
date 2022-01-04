import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/test/helpers/component-helper'
import TheSidebar from '~/components/organisms/TheSidebar.vue'
import { Menu } from '~/types/props/menu'

describe('components/organisms/TheSidebar', () => {
  let wrapper: any
  beforeEach(() => {
    wrapper = shallowMount(TheSidebar, {
      ...Options,
    })
  })

  describe('script', () => {
    describe('data', () => {
      it('selectedItem', () => {
        expect(wrapper.vm.selectedItem).toBe(-1)
      })
    })

    describe('props', () => {
      describe('current', () => {
        it('初期値', () => {
          expect(wrapper.props().current).toBe('')
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ current: '/test' })
          expect(wrapper.props().current).toBe('/test')
        })
      })

      describe('items', () => {
        it('初期値', () => {
          expect(wrapper.props().items).toEqual([])
        })

        it('値が代入されること', async () => {
          const items: Menu[] = [
            {
              name: 'テスト',
              icon: 'mdi-home',
              path: '/test',
            },
          ]
          await wrapper.setProps({ items })
          expect(wrapper.props().items).toBe(items)
        })
      })
    })

    describe('methods', () => {
      describe('onClick', () => {
        it('emitted', async () => {
          const item: Menu = { name: 'テスト', icon: 'mdi-home', path: '/test' }
          await wrapper.vm.onClick(item)
          expect(wrapper.emitted('click')).toBeTruthy()
          expect(wrapper.emitted('click')[0][0]).toBe(item)
        })
      })
    })
  })
})
