import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/test/helpers/component-helper'
import TheSubmissionTop from '~/components/templates/TheSubmissionTop.vue'
import { ShiftStatus, SubmissionStatus, TeacherShiftSummary } from '~/types/store'

describe('components/templates/TheSubmissionTop', () => {
  let wrapper: any
  beforeEach(() => {
    wrapper = shallowMount(TheSubmissionTop, {
      ...Options,
    })
  })

  describe('script', () => {
    describe('props', () => {
      describe('summaries', () => {
        it('初期値', () => {
          expect(wrapper.props().summaries).toEqual([])
        })

        it('値が代入されること', async () => {
          const summaries: TeacherShiftSummary[] = [
            {
              id: 1,
              year: 2021,
              month: 2,
              shiftStatus: ShiftStatus.FINISHED,
              submissionStatus: SubmissionStatus.SUBMITTED,
              openAt: '2021-01-01T00:00:00+09:00',
              endAt: '2021-01-15T00:00:00+09:00',
              createdAt: '2021-12-30T19:25:57+09:00',
              updatedAt: '2021-12-30T19:25:57+09:00',
            },
          ]
          await wrapper.setProps({ summaries })
          expect(wrapper.props().summaries).toBe(summaries)
        })
      })
    })

    describe('methods', () => {
      describe('getTitle', () => {
        it('return', () => {
          expect(wrapper.vm.getTitle(2022, 2)).toBe('2022年2月 シフト希望')
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

      describe('getSubmissionStatus', () => {
        it('status is waiting', () => {
          expect(wrapper.vm.getSubmissionStatus(SubmissionStatus.WAITING)).toBe('未提出')
        })

        it('status is submitted', () => {
          expect(wrapper.vm.getSubmissionStatus(SubmissionStatus.SUBMITTED)).toBe('提出済み')
        })

        it('invalid status', () => {
          expect(wrapper.vm.getSubmissionStatus(-1)).toBe('')
        })
      })

      describe('getSubmissionStatusColor', () => {
        it('status is waiting', () => {
          expect(wrapper.vm.getSubmissionStatusColor(SubmissionStatus.WAITING)).toBe('red')
        })

        it('status is submitted', () => {
          expect(wrapper.vm.getSubmissionStatusColor(SubmissionStatus.SUBMITTED)).toBe('primary')
        })

        it('invalid status', () => {
          expect(wrapper.vm.getSubmissionStatusColor(-1)).toBe('')
        })
      })

      describe('getDetailButtonMessage', () => {
        it('status is waiting', () => {
          expect(wrapper.vm.getDetailButtonMessage(SubmissionStatus.WAITING)).toBe('入力')
        })

        it('status is submitted', () => {
          expect(wrapper.vm.getDetailButtonMessage(SubmissionStatus.SUBMITTED)).toBe('修正')
        })

        it('invalid status', () => {
          expect(wrapper.vm.getDetailButtonMessage(-1)).toBe('')
        })
      })

      describe('getDetailButtonColor', () => {
        it('status is waiting', () => {
          expect(wrapper.vm.getDetailButtonColor(SubmissionStatus.WAITING)).toBe('primary')
        })

        it('status is submitted', () => {
          expect(wrapper.vm.getDetailButtonColor(SubmissionStatus.SUBMITTED)).toBe('secondary')
        })

        it('invalid status', () => {
          expect(wrapper.vm.getDetailButtonColor(-1)).toBe('')
        })
      })

      describe('onClickShow', () => {
        it('emitが実行されること', async () => {
          const summary: TeacherShiftSummary = {
            id: 1,
            year: 2021,
            month: 2,
            shiftStatus: 3,
            submissionStatus: 2,
            openAt: '2021-01-01T00:00:00+09:00',
            endAt: '2021-01-15T00:00:00+09:00',
            createdAt: '2021-12-30T19:25:57+09:00',
            updatedAt: '2021-12-30T19:25:57+09:00',
          }
          await wrapper.vm.onClickShow(summary)
          expect(wrapper.emitted('click:show')).toBeTruthy()
          expect(wrapper.emitted('click:show')[0][0]).toBe(summary)
        })
      })

      describe('onClickTop', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onClickTop()
          expect(wrapper.emitted('click:top')).toBeTruthy()
        })
      })
    })
  })
})
