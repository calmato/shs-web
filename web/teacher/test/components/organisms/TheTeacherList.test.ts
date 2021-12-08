import { mount } from '@vue/test-utils'
import * as Options from '~~/test/helpers/component-helper'
import TheTeacherList from '~/components/organisms/TheTeacherList.vue'
import { Role, Teacher } from '~/types/store'

describe('components/organisms/TheTeacherList', () => {
  let wrapper: any
  beforeEach(() => {
    wrapper = mount(TheTeacherList, {
      ...Options,
    })
  })

  describe('script', () => {
    describe('props', () => {
      describe('items', () => {
        it('初期値', () => {
          expect(wrapper.props().items).toEqual([])
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
          await wrapper.setProps({ items: teachers })
          expect(wrapper.props().items).toBe(teachers)
        })
      })

      describe('loading', () => {
        it('初期値', () => {
          expect(wrapper.props().loading).toBeFalsy()
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ loading: true })
          expect(wrapper.props().loading).toBeTruthy()
        })
      })
    })

    describe('data', () => {
      it('headers', () => {
        expect(wrapper.vm.headers).toEqual([
          { text: '講師名', value: 'name', sortable: false },
          { text: '役職', value: 'role', sortable: false },
        ])
      })
    })

    describe('methods', () => {
      describe('getRole', () => {
        it('role is teacher', () => {
          expect(wrapper.vm.getRole(Role.TEACHER)).toBe('講師')
        })

        it('role is administrator', () => {
          expect(wrapper.vm.getRole(Role.ADMINISTRATOR)).toBe('管理者')
        })

        it('invalid role', () => {
          expect(wrapper.vm.getRole(-1)).toBe('不明')
        })
      })

      describe('getRoleColor', () => {
        it('role is teacher', () => {
          expect(wrapper.vm.getRoleColor(Role.TEACHER)).toBe('primary')
        })

        it('role is administrator', () => {
          expect(wrapper.vm.getRoleColor(Role.ADMINISTRATOR)).toBe('secondary')
        })

        it('invalid type', () => {
          expect(wrapper.vm.getRoleColor(-1)).toBe('')
        })
      })
    })
  })
})
