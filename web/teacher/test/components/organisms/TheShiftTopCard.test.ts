import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/test/helpers/component-helper'
import TheShiftTopCard from '~/components/organisms/TheShiftTopCard.vue'
import { ShiftStatus, ShiftSummary } from '~/types/store'

describe('components/organisms/TheShiftTopCard', () => {
  let wrapper: any
  beforeEach(() => {
    wrapper = shallowMount(TheShiftTopCard, {
      ...Options,
    })
  })

  describe('script', () => {
    describe('props', () => {
      describe('summary', () => {
        it('初期値', () => {
          expect(wrapper.props().summary).toEqual({
            id: 0,
            year: 0,
            month: 0,
            status: ShiftStatus.UNKNOWN,
            openAt: '',
            endAt: '',
            createdAt: '',
            updatedAt: '',
          })
        })

        it('値が代入されること', async () => {
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
          await wrapper.setProps({ summary })
          expect(wrapper.props().summary).toBe(summary)
        })
      })
    })

    describe('methods', () => {
      describe('getTitle', () => {
        it('return', async () => {
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
          await wrapper.setProps({ summary })
          expect(wrapper.vm.getTitle()).toBe('2021年2月のシフト')
        })
      })

      describe('getDate', () => {
        it('return date is null', () => {
          expect(wrapper.vm.getDate()).toBe('YYYY/MM/DD')
        })

        it('return date is null', () => {
          expect(wrapper.vm.getDate('2021-01-01T00:00:00+09:00')).toBe('2021/01/01')
        })
      })

      describe('isEnabledCreateLesson', () => {
        it('return to shift status is finished', async () => {
          const summary: ShiftSummary = {
            id: 1,
            year: 2021,
            month: 2,
            status: ShiftStatus.FINISHED,
            openAt: '2021-01-01T00:00:00+09:00',
            endAt: '2021-01-15T00:00:00+09:00',
            createdAt: '2021-12-30T19:25:57+09:00',
            updatedAt: '2021-12-30T19:25:57+09:00',
          }
          await wrapper.setProps({ summary })
          expect(wrapper.vm.isEnabledCreateLesson()).toBeTruthy()
        })

        it('return to shift status is accepting', async () => {
          const summary: ShiftSummary = {
            id: 1,
            year: 2021,
            month: 2,
            status: ShiftStatus.ACCEPTING,
            openAt: '2021-01-01T00:00:00+09:00',
            endAt: '2021-01-15T00:00:00+09:00',
            createdAt: '2021-12-30T19:25:57+09:00',
            updatedAt: '2021-12-30T19:25:57+09:00',
          }
          await wrapper.setProps({ summary })
          expect(wrapper.vm.isEnabledCreateLesson()).toBeTruthy()
        })

        it('return to shift status is other', async () => {
          const summary: ShiftSummary = {
            id: 1,
            year: 2021,
            month: 2,
            status: ShiftStatus.WAITING,
            openAt: '2021-01-01T00:00:00+09:00',
            endAt: '2021-01-15T00:00:00+09:00',
            createdAt: '2021-12-30T19:25:57+09:00',
            updatedAt: '2021-12-30T19:25:57+09:00',
          }
          await wrapper.setProps({ summary })
          expect(wrapper.vm.isEnabledCreateLesson()).toBeFalsy()
        })
      })

      describe('onClickEdit', () => {
        it('emitted', async () => {
          const summary: ShiftSummary = {
            id: 1,
            year: 2021,
            month: 2,
            status: ShiftStatus.WAITING,
            openAt: '2021-01-01T00:00:00+09:00',
            endAt: '2021-01-15T00:00:00+09:00',
            createdAt: '2021-12-30T19:25:57+09:00',
            updatedAt: '2021-12-30T19:25:57+09:00',
          }
          await wrapper.setProps({ summary })
          await wrapper.vm.onClickEdit()
          expect(wrapper.emitted('click:edit')).toBeTruthy()
          expect(wrapper.emitted('click:edit')[0][0]).toBe(summary)
        })
      })

      describe('onClickNew', () => {
        it('emitted', async () => {
          const summary: ShiftSummary = {
            id: 1,
            year: 2021,
            month: 2,
            status: ShiftStatus.WAITING,
            openAt: '2021-01-01T00:00:00+09:00',
            endAt: '2021-01-15T00:00:00+09:00',
            createdAt: '2021-12-30T19:25:57+09:00',
            updatedAt: '2021-12-30T19:25:57+09:00',
          }
          await wrapper.setProps({ summary })
          await wrapper.vm.onClickNew()
          expect(wrapper.emitted('click:new')).toBeTruthy()
          expect(wrapper.emitted('click:new')[0][0]).toBe(summary)
        })
      })
    })
  })
})
