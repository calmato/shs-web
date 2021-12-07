import { shallowMount } from '@vue/test-utils'
import * as Options from '~~/test/helpers/component-helper'
import TheUserTop from '~/components/templates/TheUserTop.vue'
import { Student, Teacher } from '~/types/store'

describe('components/templates/TheUserTop', () => {
  let wrapper: any
  beforeEach(() => {
    wrapper = shallowMount(TheUserTop, {
      ...Options,
    })
  })

  describe('script', () => {
    describe('props', () => {
      describe('loading', () => {
        it('初期値', () => {
          expect(wrapper.props().loading).toBeFalsy()
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ loading: true })
          expect(wrapper.props().loading).toBeTruthy()
        })
      })

      describe('students', () => {
        it('初期値', () => {
          expect(wrapper.props().students).toEqual([])
        })

        it('値が代入されること', async () => {
          const students: Student[] = [
            {
              id: '123456789012345678901',
              name: '浜田 二郎',
              lastname: '浜田',
              firstname: '二郎',
              type: 1,
              grade: 2,
              createdAt: '',
              updatedAt: '',
            },
          ]
          await wrapper.setProps({ students })
          expect(wrapper.props().students).toBe(students)
        })
      })

      describe('teachers', () => {
        it('初期値', () => {
          expect(wrapper.props().teachers).toEqual([])
        })

        it('値が代入されること', async () => {
          const teachers: Teacher[] = [
            {
              id: '000000000000000000001',
              name: '中村 太郎',
              lastname: '中村',
              firstname: '太郎',
              role: 0,
              createdAt: '',
              updatedAt: '',
            },
          ]
          await wrapper.setProps({ teachers })
          expect(wrapper.props().teachers).toBe(teachers)
        })
      })
    })

    describe('data', () => {
      it('actors', () => {
        expect(wrapper.vm.actors).toEqual([
          { name: '講師', value: 'teachers' },
          { name: '生徒', value: 'students' },
        ])
      })

      it('selector', () => {
        expect(wrapper.vm.selector).toBe('teachers')
      })
    })
  })
})
