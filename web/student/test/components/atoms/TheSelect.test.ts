import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/test/helpers/component-helper'
import TheSelect from '~/components/atoms/TheSelect.vue'

describe('components/atoms/TheSelect', () => {
  let wrapper: any
  beforeEach(() => {
    wrapper = shallowMount(TheSelect, {
      ...Options,
    })
  })

  describe('script', () => {
    describe('props', () => {
      describe('autofocus', () => {
        it('初期値', () => {
          expect(wrapper.props().autofocus).toBeFalsy()
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ autofocus: true })
          expect(wrapper.props().autofocus).toBeTruthy()
        })
      })

      describe('label', () => {
        it('初期値', () => {
          expect(wrapper.props().label).toBe('')
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ label: 'テスト' })
          expect(wrapper.props().label).toBe('テスト')
        })
      })

      describe('name', () => {
        it('初期値', () => {
          expect(wrapper.props().name).toBe('')
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ name: 'test' })
          expect(wrapper.props().name).toBe('test')
        })
      })

      describe('outlined', () => {
        it('初期値', () => {
          expect(wrapper.props().outlined).toBeFalsy()
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ outlined: true })
          expect(wrapper.props().outlined).toBeTruthy()
        })
      })

      describe('rules', () => {
        it('初期値', () => {
          expect(wrapper.props().rules).toEqual({})
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ rules: { required: true } })
          expect(wrapper.props().rules).toEqual({ required: true })
        })
      })

      describe('chips', () => {
        it('初期値', () => {
          expect(wrapper.props().chips).toEqual(false)
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ chips: { chips: true } })
          expect(wrapper.props().chips).toEqual({ chips: true })
        })
      })

      describe('appendOuterIcon', () => {
        it('初期値', () => {
          expect(wrapper.props().appendOuterIcon).toBeUndefined()
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ appendOuterIcon: 'mdi-map' })
          expect(wrapper.props().appendOuterIcon).toBe('mdi-map')
        })
      })

      describe('items', () => {
        it('初期値', () => {
          expect(wrapper.props().items).toEqual([])
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ items: [{ text: 'test', value: 0 }] })
          expect(wrapper.props().items).toEqual([{ text: 'test', value: 0 }])
        })
      })

      describe('value', () => {
        it('初期値', () => {
          expect(wrapper.props().value).toBeUndefined()
        })

        it('値が代入されること（文字列型）', async () => {
          await wrapper.setProps({ value: 'test' })
          expect(wrapper.props().value).toBe('test')
        })

        it('値が代入されること（数値型）', async () => {
          await wrapper.setProps({ value: 1 })
          expect(wrapper.props().value).toBe(1)
        })
      })
    })

    describe('computed', () => {
      describe('formData', () => {
        it('getter', () => {
          expect(wrapper.vm.formData).toBeUndefined()
        })

        it('setter', async () => {
          await wrapper.setData({ formData: 'test' })
          expect(wrapper.emitted('update:value')).toBeTruthy()
          expect(wrapper.emitted('update:value')[0][0]).toBe('test')
        })
      })
    })
  })
})
