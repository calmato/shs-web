import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/test/helpers/component-helper'
import PcShiftTop from '~/components/templates/PcShiftTop.vue'
import { ShiftSummary } from '~/types/store'
import { ShiftsNewForm, ShiftsNewOptions, ShiftsNewParams } from '~/types/form'

describe('components/templates/PcShiftTop', () => {
  let wrapper: any
  beforeEach(() => {
    wrapper = shallowMount(PcShiftTop, {
      ...Options,
    })
  })

  describe('script', () => {
    describe('props', () => {
      describe('newDialog', () => {
        it('初期値', () => {
          expect(wrapper.props().newDialog).toBeFalsy()
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ newDialog: true })
          expect(wrapper.props().newDialog).toBeTruthy()
        })
      })

      describe('newForm', () => {
        it('初期値', () => {
          expect(wrapper.props().newForm).toEqual({
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
          await wrapper.setProps({ newForm: form })
          expect(wrapper.props().newForm).toBe(form)
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

      describe('acceptingSummaries', () => {
        it('初期値', () => {
          expect(wrapper.props().acceptingSummaries).toEqual([])
        })

        it('値が代入されること', async () => {
          const summaries: ShiftSummary[] = [
            {
              id: 1,
              year: 2021,
              month: 2,
              status: 3,
              openAt: '2021-01-01T00:00:00+09:00',
              endAt: '2021-01-15T00:00:00+09:00',
              createdAt: '2021-12-30T19:25:57+09:00',
              updatedAt: '2021-12-30T19:25:57+09:00',
            },
          ]
          await wrapper.setProps({ acceptingSummaries: summaries })
          expect(wrapper.props().acceptingSummaries).toBe(summaries)
        })
      })

      describe('finishedSummaries', () => {
        it('初期値', () => {
          expect(wrapper.props().finishedSummaries).toEqual([])
        })

        it('値が代入されること', async () => {
          const summaries: ShiftSummary[] = [
            {
              id: 1,
              year: 2021,
              month: 2,
              status: 3,
              openAt: '2021-01-01T00:00:00+09:00',
              endAt: '2021-01-15T00:00:00+09:00',
              createdAt: '2021-12-30T19:25:57+09:00',
              updatedAt: '2021-12-30T19:25:57+09:00',
            },
          ]
          await wrapper.setProps({ finishedSummaries: summaries })
          expect(wrapper.props().finishedSummaries).toBe(summaries)
        })
      })

      describe('waitingSummaries', () => {
        it('初期値', () => {
          expect(wrapper.props().waitingSummaries).toEqual([])
        })

        it('値が代入されること', async () => {
          const summaries: ShiftSummary[] = [
            {
              id: 1,
              year: 2021,
              month: 2,
              status: 3,
              openAt: '2021-01-01T00:00:00+09:00',
              endAt: '2021-01-15T00:00:00+09:00',
              createdAt: '2021-12-30T19:25:57+09:00',
              updatedAt: '2021-12-30T19:25:57+09:00',
            },
          ]
          await wrapper.setProps({ waitingSummaries: summaries })
          expect(wrapper.props().waitingSummaries).toBe(summaries)
        })
      })
    })

    describe('data', () => {
      it('deleteDialog', () => {
        expect(wrapper.vm.deleteDialog).toBeFalsy()
      })
    })

    describe('methods', () => {
      describe('toggleNewDialog', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.toggleNewDialog()
          expect(wrapper.emitted('toggle:new-dialog')).toBeTruthy()
        })
      })

      describe('toggleEditDialog', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.toggleEditDialog()
          expect(wrapper.emitted('toggle:edit-dialog')).toBeTruthy()
        })

        it('deleteDialogがfalseであること', async () => {
          await wrapper.vm.toggleEditDialog()
          expect(wrapper.vm.deleteDialog).toBeFalsy()
        })
      })

      describe('onClickNewShift', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onClickNewShift()
          expect(wrapper.emitted('click:new-shift')).toBeTruthy()
        })
      })

      describe('onClickEditshift', () => {
        it('emitが実行されること', async () => {
          const summary: ShiftSummary = {
            id: 1,
            year: 2021,
            month: 2,
            status: 3,
            openAt: '2021-01-01T00:00:00+09:00',
            endAt: '2021-01-15T00:00:00+09:00',
            createdAt: '2021-12-30T19:25:57+09:00',
            updatedAt: '2021-12-30T19:25:57+09:00',
          }
          await wrapper.vm.onClickEditShift(summary)
          expect(wrapper.emitted('click:edit-shift')).toBeTruthy()
          expect(wrapper.emitted('click:edit-shift')[0][0]).toBe(summary)
        })
      })

      describe('onClickDeleteShift', () => {
        it('deleteDialogがtrueであること', async () => {
          await wrapper.vm.onClickDeleteShift()
          expect(wrapper.vm.deleteDialog).toBeTruthy()
        })
      })

      describe('onClickDeleteAccept', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onClickDeleteAccept()
          expect(wrapper.emitted('submit:delete')).toBeTruthy()
        })

        it('deleteDialogがfalseであること', async () => {
          await wrapper.vm.onClickDeleteAccept()
          expect(wrapper.vm.deleteDialog).toBeFalsy()
        })
      })

      describe('onClickDeleteCancel', () => {
        it('deleteDialogがfalseであること', async () => {
          await wrapper.vm.onClickDeleteCancel()
          expect(wrapper.vm.deleteDialog).toBeFalsy()
        })
      })

      describe('onClickNewLesson', () => {
        it('emitが実行されること', async () => {
          const summary: ShiftSummary = {
            id: 1,
            year: 2021,
            month: 2,
            status: 3,
            openAt: '2021-01-01T00:00:00+09:00',
            endAt: '2021-01-15T00:00:00+09:00',
            createdAt: '2021-12-30T19:25:57+09:00',
            updatedAt: '2021-12-30T19:25:57+09:00',
          }
          await wrapper.vm.onClickNewLesson(summary)
          expect(wrapper.emitted('click:new-lesson')).toBeTruthy()
          expect(wrapper.emitted('click:new-lesson')[0][0]).toBe(summary)
        })
      })

      describe('onClickAddClosedDate', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onClickAddClosedDate()
          expect(wrapper.emitted('click:add-closed-date')).toBeTruthy()
        })
      })

      describe('onClickRemoveClosedDate', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onClickRemoveClosedDate()
          expect(wrapper.emitted('click:remove-closed-date')).toBeTruthy()
        })
      })

      describe('onSubmitNew', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onSubmitNew()
          expect(wrapper.emitted('submit:new')).toBeTruthy()
        })
      })

      describe('onSubmitEdit', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onSubmitEdit()
          expect(wrapper.emitted('submit:edit')).toBeTruthy()
        })
      })
    })
  })
})
