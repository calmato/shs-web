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
              nameKana: 'なかむら たろう',
              lastName: '中村',
              firstName: '太郎',
              lastNameKana: 'なかむら',
              firstNameKana: 'たろう',
              mail: 'teacher-001@calmato.jp',
              role: 0,
              subjects: {},
              createdAt: '',
              updatedAt: '',
            },
          ]
          await wrapper.setProps({ items: teachers })
          expect(wrapper.props().items).toBe(teachers)
        })
      })

      describe('itemsPerPage', () => {
        it('初期値', () => {
          expect(wrapper.props().itemsPerPage).toBe(10)
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ itemsPerPage: 50 })
          expect(wrapper.props().itemsPerPage).toBe(50)
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

      describe('page', () => {
        it('初期値', () => {
          expect(wrapper.props().page).toBe(1)
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ page: 2 })
          expect(wrapper.props().page).toBe(2)
        })
      })

      describe('total', () => {
        it('初期値', () => {
          expect(wrapper.props().total).toBe(0)
        })

        it('値が代入されること', async () => {
          await wrapper.setProps({ total: 100 })
          expect(wrapper.props().total).toBe(100)
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

      it('footer', () => {
        expect(wrapper.vm.footer).toEqual({
          itemsPerPageText: '表示件数',
          itemsPerPageOptions: [10, 20, 30, 50],
        })
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

      describe('onClick', () => {
        it('emitted', async () => {
          const teacher: Teacher = {
            id: '000000000000000000001',
            name: '中村 太郎',
            nameKana: 'なかむら たろう',
            lastName: '中村',
            firstName: '太郎',
            lastNameKana: 'なかむら',
            firstNameKana: 'たろう',
            mail: 'teacher-001@calmato.jp',
            role: 0,
            subjects: {},
            createdAt: '',
            updatedAt: '',
          }
          await wrapper.vm.onClick(teacher)
          expect(wrapper.emitted('click')).toBeTruthy()
          expect(wrapper.emitted('click')[0][0]).toBe(teacher)
        })
      })
    })
  })
})
