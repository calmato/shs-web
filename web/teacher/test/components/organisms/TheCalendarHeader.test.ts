import { shallowMount } from '@vue/test-utils'
import dayjs from 'dayjs'
import * as Options from '~~/test/helpers/component-helper'
import TheCalendarHeader from '~/components/organisms/TheCalendarHeader.vue'
import { Date } from '~/types/props/calendar'

describe('components/organisms/TheCalendarHeader', () => {
  let wrapper: any
  beforeEach(() => {
    wrapper = shallowMount(TheCalendarHeader, {
      ...Options,
    })
  })

  describe('script', () => {
    describe('props', () => {
      describe('end', () => {
        it('初期値', () => {
          expect(wrapper.props().end).toBeNull()
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
          expect(wrapper.props().start).toBeNull()
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
    })

    describe('computed', () => {
      describe('calendarType', () => {
        it('getter', () => {
          expect(wrapper.vm.calendarType).toBe('month')
        })
      })
    })

    describe('methods', () => {
      describe('getTitle', () => {
        it('return: props.now is null', async () => {
          await wrapper.setProps({ now: null })
          expect(wrapper.vm.getTitle()).toBe('')
        })

        it('return: props.now is not null', async () => {
          const now = dayjs()
          await wrapper.setProps({ now })
          expect(wrapper.vm.getTitle()).toBe(now.format('YYYY年MM月'))
        })

        it('return: props.start and props.end are not null', async () => {
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
          await wrapper.setProps({ start, end })
          expect(wrapper.vm.getTitle()).toBe(dayjs('2021-08-02').format('YYYY年MM月'))
        })
      })

      describe('getTypeName', () => {
        it('return: type is unmatch', () => {
          const type: string = 'year'
          expect(wrapper.vm.getTypeName(type)).toBe('')
        })

        it('return: type is match', () => {
          const type: string = 'month'
          expect(wrapper.vm.getTypeName(type)).toBe('月')
        })
      })

      describe('onClickToday', () => {
        it('emitted', async () => {
          await wrapper.vm.onClickToday()
          expect(wrapper.emitted('click:today')).toBeTruthy()
        })
      })

      describe('onClickPrev', () => {
        it('emitted', async () => {
          await wrapper.vm.onClickPrev()
          expect(wrapper.emitted('click:prev')).toBeTruthy()
        })
      })

      describe('onClickNext', () => {
        it('emitted', async () => {
          await wrapper.vm.onClickNext()
          expect(wrapper.emitted('click:next')).toBeTruthy()
        })
      })
    })
  })
})
