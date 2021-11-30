import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/test/helpers/component-helper'
import Top from '~/components/templates/Top.vue'
import { Event } from '~/types/props/calendar'

describe('components/templates/SignIn', () => {
  let wrapper: any
  beforeEach(() => {
    wrapper = shallowMount(Top, {
      ...Options,
    })
  })

  describe('script', () => {
    describe('props', () => {
      describe('events', () => {
        it('初期値', () => {
          expect(wrapper.props().events).toEqual([])
        })

        it('値が代入されること', async () => {
          const events: Event[] = [
            {
              name: '月次報告',
              start: '2021-11-25 09:00',
              end: '2021-11-26 17:00',
              color: 'primary',
            },
            {
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
    })

    describe('methods', () => {
      describe('onClick', () => {
        it('updated focus', async () => {
          await wrapper.setData({ focus: '' })
          await wrapper.vm.setToday()
          expect(wrapper.vm.focus).not.toBe('')
        })
      })
    })
  })
})
