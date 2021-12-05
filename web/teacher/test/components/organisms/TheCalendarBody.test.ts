import { mount } from '@vue/test-utils'
import dayjs from 'dayjs'
import * as Options from '~~/test/helpers/component-helper'
import TheCalendarBody from '~/components/organisms/TheCalendarBody.vue'
import { Date, Event } from '~/types/props/calendar'

describe('components/organisms/TheCalendarBody', () => {
  let wrapper: any
  beforeEach(() => {
    wrapper = mount(TheCalendarBody, {
      ...Options,
    })
  })

  describe('script', () => {
    describe('props', () => {
      describe('end', () => {
        it('初期値', () => {
          expect(wrapper.props().end).toEqual({
            date: '2021-07-31',
            time: '23:59:59',
            year: 2021,
            month: 8,
            weekday: 6,
            day: 31,
            hour: 23,
            minute: 59,
            future: false,
            hasDay: false,
            hasTime: false,
            past: false,
            present: false,
          })
        })

        it('値が代入されること', async () => {
          const end: Date = {
            date: '2021-08-02',
            time: '18:30:00',
            year: 2021,
            month: 8,
            weekday: 1,
            day: 2,
            hour: 18,
            minute: 30,
            future: false,
            hasDay: false,
            hasTime: false,
            past: false,
            present: false,
          }
          await wrapper.setProps({ end })
          expect(wrapper.props().end).toBe(end)
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
              start: '2021-11-25T09:00:00+09:00',
              end: '2021-11-26T17:00:00+09:00',
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

      describe('focus', () => {
        it('初期値', () => {
          expect(wrapper.props().focus).toBe('')
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ focus: '2021-08-02' })
          expect(wrapper.props().focus).toBe('2021-08-02')
        })
      })

      describe('now', () => {
        it('初期値', () => {
          expect(wrapper.props().now).toBeNull()
        })

        it('値が代入されること', async () => {
          const now = dayjs()
          await wrapper.setProps({ now })
          expect(wrapper.props().now).toBe(now)
        })
      })

      describe('start', () => {
        it('初期値', () => {
          expect(wrapper.props().start).toEqual({
            date: '2021-07-01',
            time: '00:00:00',
            year: 2021,
            month: 8,
            weekday: 4,
            day: 1,
            hour: 0,
            minute: 0,
            future: false,
            hasDay: false,
            hasTime: false,
            past: false,
            present: false,
          })
        })

        it('値が代入されること', async () => {
          const start: Date = {
            date: '2021-08-02',
            time: '18:30:00',
            year: 2021,
            month: 8,
            weekday: 1,
            day: 2,
            hour: 18,
            minute: 30,
            future: false,
            hasDay: false,
            hasTime: false,
            past: false,
            present: false,
          }
          await wrapper.setProps({ start })
          expect(wrapper.props().start).toBe(start)
        })
      })

      describe('type', () => {
        it('初期値', () => {
          expect(wrapper.props().type).toBe('month')
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ type: 'week' })
          expect(wrapper.props().type).toBe('week')
        })
      })

      describe('types', () => {
        it('初期値', () => {
          expect(wrapper.props().types).toEqual([
            { name: '月', value: 'month' },
            { name: '週', value: 'week' },
            { name: '日', value: 'day' },
          ])
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ types: [] })
          expect(wrapper.props().types).toEqual([])
        })
      })

      describe('weekdays', () => {
        it('初期値', () => {
          expect(wrapper.props().weekdays).toEqual([0, 1, 2, 3, 4, 5, 6])
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ weekdays: [1, 2, 3, 4, 5, 6, 0] })
          expect(wrapper.props().weekdays).toEqual([1, 2, 3, 4, 5, 6, 0])
        })
      })
    })

    describe('computed', () => {
      describe('currentTime', () => {
        it('getter', () => {
          expect(wrapper.vm.currentTime).toBe('2021-07-23 20:00:00')
        })
      })

      describe('calendarFocus', () => {
        it('getter', () => {
          expect(wrapper.vm.calendarFocus).toBe('')
        })
      })

      describe('calendarType', () => {
        it('getter', () => {
          expect(wrapper.vm.calendarType).toBe('month')
        })
      })
    })

    describe('methods', () => {
      describe('getParseEvents', () => {
        it('return', async () => {
          const events: Event[] = [
            {
              lessonId: 1,
              name: '月次報告',
              start: '2021-11-25T09:00:00+09:00',
              end: '2021-11-26T17:00:00+09:00',
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
          expect(wrapper.vm.getParseEvents()).toEqual([
            {
              lessonId: 1,
              name: '月次報告',
              start: '2021-11-25 09:00:00',
              end: '2021-11-26 17:00:00',
              color: 'primary',
            },
            {
              lessonId: 2,
              name: '出張',
              start: '2021-11-26 09:00:00',
              end: '2021-11-28 09:00:00',
              color: 'secondary',
            },
          ])
        })
      })

      describe('getEventColor', () => {
        it('return: event is null', () => {
          expect(wrapper.vm.getEventColor()).toBe('primary')
        })

        it('return: event is not null', () => {
          const event: Event = {
            lessonId: 1,
            name: '出張',
            start: '2021-11-26',
            end: '2021-11-28',
            color: 'secondary',
          }
          expect(wrapper.vm.getEventColor(event)).toBe('secondary')
        })
      })

      describe('viewDay', () => {
        it('emitted', async () => {
          const date: Date = {
            date: '2021-08-02',
            time: '18:30:00',
            year: 2021,
            month: 8,
            weekday: 1,
            day: 2,
            hour: 18,
            minute: 30,
            future: false,
            hasDay: false,
            hasTime: false,
            past: false,
            present: false,
          }
          await wrapper.vm.viewDay(date)
          expect(wrapper.emitted('update:focus')).toBeTruthy()
          expect(wrapper.emitted('update:type')).toBeTruthy()
          expect(wrapper.emitted('update:type')[0][0]).toBe('day')
        })
      })

      describe('showEvent', () => {
        it('emitted', async () => {
          const event: Event = {
            lessonId: 1,
            name: '出張',
            start: '2021-11-26',
            end: '2021-11-28',
            color: 'secondary',
          }
          await wrapper.vm.showEvent({ event })
          expect(wrapper.emitted('click')).toBeTruthy()
          expect(wrapper.emitted('click')[0][0]).toBe(event)
        })
      })

      describe('updateRange', () => {
        it('emitted', async () => {
          const start: Date = {
            date: '2021-08-02',
            time: '18:30:00',
            year: 2021,
            month: 8,
            weekday: 1,
            day: 2,
            hour: 18,
            minute: 30,
            future: false,
            hasDay: false,
            hasTime: false,
            past: false,
            present: false,
          }
          const end: Date = {
            date: '2021-08-02',
            time: '18:30:00',
            year: 2021,
            month: 8,
            weekday: 1,
            day: 2,
            hour: 18,
            minute: 30,
            future: false,
            hasDay: false,
            hasTime: false,
            past: false,
            present: false,
          }
          await wrapper.vm.updateRange({ start, end })
          expect(wrapper.emitted('update:start')).toBeTruthy()
          expect(wrapper.emitted('update:end')).toBeTruthy()
        })
      })
    })
  })
})
