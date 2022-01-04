import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/test/helpers/component-helper'
import TheDatePicker from '~/components/molecules/TheDatePicker.vue'

describe('components/molecules/TheDatePicker', () => {
  let wrapper: any
  beforeEach(() => {
    wrapper = shallowMount(TheDatePicker, {
      ...Options,
    })
  })

  describe('script', () => {
    describe('props', () => {
      describe('label', () => {
        it('初期値', () => {
          expect(wrapper.props().label).toBe('')
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ label: 'テスト' })
          expect(wrapper.props().label).toBe('テスト')
        })
      })

      describe('locale', () => {
        it('初期値', () => {
          expect(wrapper.props().locale).toBe('ja')
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ locale: 'us' })
          expect(wrapper.props().locale).toBe('us')
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

      describe('rules', () => {
        it('初期値', () => {
          expect(wrapper.props().rules).toEqual({})
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ rules: { required: true } })
          expect(wrapper.props().rules).toEqual({ required: true })
        })
      })

      describe('type', () => {
        it('初期値', () => {
          expect(wrapper.props().type).toBe('date')
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ type: 'month' })
          expect(wrapper.props().type).toBe('month')
        })
      })

      describe('value', () => {
        it('初期値', () => {
          expect(wrapper.props().value).toBe('')
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ value: 'test' })
          expect(wrapper.props().value).toBe('test')
        })
      })
    })

    describe('computed', () => {
      describe('formData', () => {
        it('getter', () => {
          expect(wrapper.vm.formData).toBe('')
        })

        it('setter', async () => {
          await wrapper.setData({ formData: '2022-01-01' })
          expect(wrapper.emitted('update:value')).toBeTruthy()
          expect(wrapper.emitted('update:value')[0][0]).toBe('2022-01-01')
        })
      })
    })
  })
})
