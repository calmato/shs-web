import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/test/helpers/component-helper'
import TheShiftNewCard from '~/components/organisms/TheShiftNewCard.vue'
import { ShiftsNewForm, ShiftsNewOptions, ShiftsNewParams } from '~/types/form'

describe('components/organisms/TheShiftNewCard', () => {
  let wrapper: any
  beforeEach(() => {
    wrapper = shallowMount(TheShiftNewCard, {
      ...Options,
    })
  })

  describe('script', () => {
    describe('props', () => {
      describe('form', () => {
        describe('form', () => {
          it('初期値', () => {
            expect(wrapper.props().form).toEqual({
              params: ShiftsNewParams,
              options: ShiftsNewOptions,
            })
          })

          it('値が代入されること', async () => {
            const form: ShiftsNewForm = {
              params: {
                yearMonth: '2022-02',
                openDate: '2022-01-01',
                endDate: '2022-01-15',
                closedDates: [],
              },
              options: ShiftsNewOptions,
            }
            await wrapper.setProps({ form })
            expect(wrapper.props().form).toBe(form)
          })
        })

        describe('loading', () => {
          it('初期値', () => {
            expect(wrapper.props().loading).toBeFalsy()
          })

          it('値が代入されること', async () => {
            await wrapper.setProps({ loading: true })
            expect(wrapper.props().loading).toBeTruthy()
          })
        })
      })

      describe('methods', () => {
        describe('addClosedDate', () => {
          it('emitが実行されること', async () => {
            await wrapper.vm.addClosedDate()
            expect(wrapper.emitted('click:add')).toBeTruthy()
          })
        })

        describe('removeClosedDate', () => {
          it('emitが実行されること', async () => {
            await wrapper.vm.removeClosedDate()
            expect(wrapper.emitted('click:remove')).toBeTruthy()
          })
        })

        describe('onSubmit', () => {
          it('emitが実行されること', async () => {
            await wrapper.vm.onSubmit()
            expect(wrapper.emitted('click:submit')).toBeTruthy()
          })
        })

        describe('onClose', () => {
          it('emitが実行されること', async () => {
            await wrapper.vm.onClose()
            expect(wrapper.emitted('click:close')).toBeTruthy()
          })
        })
      })
    })
  })
})
