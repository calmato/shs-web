import { mount } from '@vue/test-utils'
import * as Options from '~~/test/helpers/component-helper'
import TheStudentList from '~/components/organisms/TheStudentList.vue'
import { Student } from '~/types/store'

describe('components/organisms/TheStudentList', () => {
  let wrapper: any
  beforeEach(() => {
    wrapper = mount(TheStudentList, {
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
              schoolType: '小学校',
              grade: 2,
              subjects: [],
              createdAt: '',
              updatedAt: '',
            },
          ]
          await wrapper.setProps({ items: students })
          expect(wrapper.props().items).toBe(students)
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
          { text: '生徒名', value: 'name', sortable: false },
          { text: '校種', value: 'type', sortable: false },
          { text: '学年', value: 'grade', sortable: false },
        ])
      })
    })

    describe('getSchoolTypeColor', () => {
      it('type is elementary school', () => {
        expect(wrapper.vm.getSchoolTypeColor('小学校')).toBe('primary')
      })

      it('type is junior high school', () => {
        expect(wrapper.vm.getSchoolTypeColor('中学校')).toBe('secondary')
      })

      it('type is high school', () => {
        expect(wrapper.vm.getSchoolTypeColor('高校')).toBe('info')
      })

      it('invalid type', () => {
        expect(wrapper.vm.getSchoolTypeColor('その他')).toBe('')
      })
    })

    describe('getSchoolYear', () => {
      it('return', () => {
        expect(wrapper.vm.getSchoolYear(3)).toBe('3年')
      })
    })
  })
})
