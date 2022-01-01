import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/test/helpers/component-helper'
import TheShiftEditCard from '~/components/organisms/TheShiftEditCard.vue'
import {
  ShiftSummaryEditScheduleForm,
  ShiftSummaryEditScheduleOptions,
  ShiftSummaryEditScheduleParams,
} from '~/types/form'

describe('components/organisms/TheShiftNewCard', () => {
  let wrapper: any
  beforeEach(() => {
    wrapper = shallowMount(TheShiftEditCard, {
      ...Options,
    })
  })

  describe('script', () => {
    describe('props', () => {
      describe('form', () => {
        describe('form', () => {
          it('初期値', () => {
            expect(wrapper.props().form).toEqual({
              params: ShiftSummaryEditScheduleParams,
              options: ShiftSummaryEditScheduleOptions,
            })
          })

          it('値が代入されること', async () => {
            const form: ShiftSummaryEditScheduleForm = {
              params: {
                summaryId: 1,
                openDate: '2022-01-01',
                endDate: '2022-01-15',
              },
              options: ShiftSummaryEditScheduleOptions,
            }
            await wrapper.setProps({ form })
            expect(wrapper.props().form).toBe(form)
          })
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

      describe('deleteDialog', () => {
        it('初期値', () => {
          expect(wrapper.props().deleteDialog).toBeFalsy()
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ deleteDialog: true })
          expect(wrapper.props().deleteDialog).toBeTruthy()
        })
      })
    })

    describe('methods', () => {
      describe('onDelete', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onDelete()
          expect(wrapper.emitted('click:delete')).toBeTruthy()
        })
      })

      describe('onDeleteAccept', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onDeleteAccept()
          expect(wrapper.emitted('click:delete-accept')).toBeTruthy()
        })
      })

      describe('onDeleteCancel', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onDeleteCancel()
          expect(wrapper.emitted('click:delete-cancel')).toBeTruthy()
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
