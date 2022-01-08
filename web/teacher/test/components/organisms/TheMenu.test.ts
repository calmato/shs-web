import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/test/helpers/component-helper'
import TheMenu from '~/components/organisms/TheMenu.vue'
import { Menu } from '~/types/props/menu'

describe('components/organisms/TheMenu', () => {
  let wrapper: any
  beforeEach(() => {
    wrapper = shallowMount(TheMenu, {
      ...Options,
    })
  })

  describe('script', () => {
    describe('props', () => {
      describe('absolute', () => {
        it('初期値', () => {
          expect(wrapper.props().absolute).toBeFalsy()
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ absolute: true })
          expect(wrapper.props().absolute).toBeTruthy()
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
              filter: 'all',
            },
          ]
          await wrapper.setProps({ items })
          expect(wrapper.props().items).toBe(items)
        })
      })

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
      describe('onClickItem', () => {
        it('emitted', async () => {
          const item: Menu = { name: 'テスト', icon: 'mdi-home', path: '/test', filter: 'all' }
          await wrapper.vm.onClickItem(item)
          expect(wrapper.emitted('click:item')).toBeTruthy()
          expect(wrapper.emitted('click:item')[0][0]).toBe(item)
        })
      })

      describe('onClickClose', () => {
        it('emitted', async () => {
          await wrapper.vm.onClickClose()
          expect(wrapper.emitted('click:close')).toBeTruthy()
        })
      })
    })
  })
})
