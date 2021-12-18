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
              nameKana: 'はまだ じろう',
              lastName: '浜田',
              firstName: '二郎',
              lastNameKana: 'はまだ',
              firstNameKana: 'じろう',
              mail: 'student-001@calmato.jp',
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
              nameKana: 'なかむら たろう',
              lastName: '中村',
              firstName: '太郎',
              lastNameKana: 'なかむら',
              firstNameKana: 'たろう',
              mail: 'teacher-001@calmato.jp',
              role: 0,
              createdAt: '',
              updatedAt: '',
            },
          ]
          await wrapper.setProps({ teachers })
          expect(wrapper.props().teachers).toBe(teachers)
        })
      })

      describe('teachersTotal', () => {
        it('初期値', () => {
          expect(wrapper.props().teachersTotal).toBe(0)
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ teachersTotal: 100 })
          expect(wrapper.props().teachersTotal).toBe(100)
        })
      })

      describe('teachersPage', () => {
        it('初期値', () => {
          expect(wrapper.props().teachersPage).toBe(1)
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ teachersPage: 2 })
          expect(wrapper.props().teachersPage).toBe(2)
        })
      })

      describe('teachersItemsPerPage', () => {
        it('初期値', () => {
          expect(wrapper.props().teachersItemsPerPage).toBe(10)
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ teachersItemsPerPage: 50 })
          expect(wrapper.props().teachersItemsPerPage).toBe(50)
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

    describe('methods', () => {
      describe('onClickNew', () => {
        it('emitが実行されること', async () => {
          await wrapper.vm.onClickNew('teachers')
          expect(wrapper.emitted('click:new')).toBeTruthy()
          expect(wrapper.emitted('click:new')[0][0]).toBe('teachers')
        })
      })
    })
  })
})
