import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/test/helpers/component-helper'
import TheCalendarDetail from '~/components/organisms/TheCalendarDetail.vue'
import { EventDetail } from '~/types/props/calendar'

describe('components/organisms/TheCalendarDetail', () => {
  let wrapper: any
  beforeEach(() => {
    wrapper = shallowMount(TheCalendarDetail, {
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
            start: '2021-11-25T18:30:00+09:00',
            end: '2021-11-25T20:00:00+09:00',
            remark: '漢字テスト3~4ページを行う',
          }
          await wrapper.setProps({ detail })
          expect(wrapper.props().detail).toBe(detail)
        })
      })
    })

    describe('methods', () => {
      describe('getLessonDate', () => {
        it('return', async () => {
          const detail: EventDetail = {
            lessonId: 1,
            subject: '国語',
            teacher: '中村 一郎',
            student: '市川 二郎',
            start: '2021-11-25T18:30:00+09:00',
            end: '2021-11-25T20:00:00+09:00',
            remark: '漢字テスト3~4ページを行う',
          }
          await wrapper.setProps({ detail })
          expect(wrapper.vm.getLessonDate()).toBe('11月25日')
        })
      })

      describe('getLessonTime', () => {
        it('return', async () => {
          const detail: EventDetail = {
            lessonId: 1,
            subject: '国語',
            teacher: '中村 一郎',
            student: '市川 二郎',
            start: '2021-11-25T18:30:00+09:00',
            end: '2021-11-25T20:00:00+09:00',
            remark: '漢字テスト3~4ページを行う',
          }
          await wrapper.setProps({ detail })
          expect(wrapper.vm.getLessonTime()).toBe('18:30 ~ 20:00')
        })
      })

      describe('onClose', () => {
        it('emitted', async () => {
          await wrapper.vm.onClose()
          expect(wrapper.emitted('click:close')).toBeTruthy()
        })
      })
    })
  })
})
