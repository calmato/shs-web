import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/test/helpers/component-helper'
import PcShiftTop from '~/components/templates/PcShiftTop.vue'
import { ShiftSummary } from '~/types/store'

describe('components/templates/PcShiftTop', () => {
  let wrapper: any
  beforeEach(() => {
    wrapper = shallowMount(PcShiftTop, {
      ...Options,
    })
  })

  describe('script', () => {
    describe('props', () => {
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

    describe('methods', () => {
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
    })
  })
})
