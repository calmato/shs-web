import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/test/helpers/component-helper'
import Top from '~/components/templates/Top.vue'
import { Event, EventDetail } from '~/types/props/calendar'

describe('components/templates/SignIn', () => {
  let wrapper: any
  beforeEach(() => {
    wrapper = shallowMount(Top, {
      ...Options,
    })
  })

  describe('script', () => {
    describe('props', () => {
      describe('detail', () => {
        it('初期値', () => {
          expect(wrapper.props().detail).toEqual({})
        })

        it('値が代入されること', async () => {
          const detail: EventDetail = {
            lessonId: 1,
            subject: '国語',
            teacher: '中村 一郎',
            student: '市川 二郎',
            start: '2021-11-25 09:00:00',
            end: '2021-11-26 17:00:00',
            remark: '漢字テスト3~4ページを行う',
          }
          await wrapper.setProps({ detail })
          expect(wrapper.props().detail).toBe(detail)
        })
      })

      describe('events', () => {
        it('初期値', () => {
          expect(wrapper.props().events).toEqual([])
        })

        it('値が代入されること', async () => {
          const events: Event[] = [
            {
              lessonId: 1,
              name: '月次報告',
              start: '2021-11-25 09:00',
              end: '2021-11-26 17:00',
              color: 'primary',
            },
            {
              lessonId: 2,
              name: '出張',
              start: '2021-11-26',
              end: '2021-11-28',
              color: 'secondary',
            },
          ]
          await wrapper.setProps({ events })
          expect(wrapper.props().events).toBe(events)
        })
      })
    })

    describe('data', () => {
      it('weekdays', () => {
        expect(wrapper.vm.weekdays).toEqual([0, 1, 2, 3, 4, 5, 6])
      })

      it('types', () => {
        expect(wrapper.vm.types).toEqual([
          { name: '月', value: 'month' },
          { name: '週', value: 'week' },
          { name: '日', value: 'day' },
        ])
      })

      it('start', () => {
        expect(wrapper.vm.start).toBeUndefined()
      })

      it('end', () => {
        expect(wrapper.vm.end).toBeUndefined()
      })

      it('focus', () => {
        expect(wrapper.vm.focus).toBe('')
      })

      it('type', () => {
        expect(wrapper.vm.type).toBe('month')
      })

      it('dialog', () => {
        expect(wrapper.vm.dialog).toBeFalsy()
      })
    })

    describe('methods', () => {
      describe('toggleDialog', () => {
        it('updated dialog', async () => {
          await wrapper.setData({ dialog: false })
          await wrapper.vm.toggleDialog()
          expect(wrapper.vm.dialog).toBeTruthy()
        })
      })

      describe('showEvent', () => {
        let event: Event
        beforeEach(() => {
          event = {
            lessonId: 1,
            name: '月次報告',
            start: '2021-11-25 09:00',
            end: '2021-11-26 17:00',
            color: 'primary',
          }
        })

        it('updated dialog', async () => {
          await wrapper.setData({ dialog: false })
          await wrapper.vm.showEvent(event)
          expect(wrapper.vm.dialog).toBeTruthy()
        })

        it('emitted', async () => {
          await wrapper.vm.showEvent(event)
          expect(wrapper.emitted('click')).toBeTruthy()
          expect(wrapper.emitted('click')[0][0]).toBe(event)
        })
      })

      describe('setToday', () => {
        it('updated focus', async () => {
          await wrapper.setData({ focus: '' })
          await wrapper.vm.setToday()
          expect(wrapper.vm.focus).not.toBe('')
        })
      })
    })
  })
})
