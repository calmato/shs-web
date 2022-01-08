import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/test/helpers/component-helper'
import TheSelectWithSlot from '~/components/atoms/TheSelectWithSlot.vue'

describe('components/atoms/TheSelectWithSlot', () => {
  let wrapper: any
  beforeEach(() => {
    wrapper = shallowMount(TheSelectWithSlot, {
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

      describe('chips', () => {
        it('初期値', () => {
          expect(wrapper.props().chips).toBeFalsy()
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ chips: true })
          expect(wrapper.props().chips).toBeTruthy()
        })
      })

      describe('multiple', () => {
        it('初期値', () => {
          expect(wrapper.props().multiple).toBeFalsy()
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ multiple: true })
          expect(wrapper.props().multiple).toBeTruthy()
        })
      })

      describe('rules', () => {
        it('初期値', () => {
          expect(wrapper.props().rules).toBeUndefined()
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ rules: { required: true } })
          expect(wrapper.props().rules).toEqual({ required: true })
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

      describe('itemText', () => {
        it('初期値', () => {
          expect(wrapper.props().itemText).toBeUndefined()
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ itemText: 'name' })
          expect(wrapper.props().itemText).toBe('name')
        })
      })

      describe('itemValue', () => {
        it('初期値', () => {
          expect(wrapper.props().itemValue).toBeUndefined()
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ itemValue: 'id' })
          expect(wrapper.props().itemValue).toBe('id')
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

        it('値が代入されること（配列型）', async () => {
          await wrapper.setProps({ value: [1, 2] })
          expect(wrapper.props().value).toEqual([1, 2])
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

      describe('isSuccess', () => {
        it('props.rules is undefinied', () => {
          expect(wrapper.vm.isSuccess(true)).toBeFalsy()
          expect(wrapper.vm.isSuccess(false)).toBeFalsy()
        })

        it('props.rules is not undefined', async () => {
          await wrapper.setProps({ rules: {} })
          expect(wrapper.vm.isSuccess(true)).toBeTruthy()
          expect(wrapper.vm.isSuccess(false)).toBeFalsy()
        })
      })

      describe('onBlur', () => {
        it('emitted', async () => {
          await wrapper.vm.onBlur()
          expect(wrapper.emitted('blur')).toBeTruthy()
        })
      })
    })
  })
})
