import { mount } from '@vue/test-utils'
import * as Options from '~~/test/helpers/component-helper'
import TheStudentList from '~/components/organisms/TheStudentList.vue'
import { SchoolType, Student } from '~/types/store'

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
              lastname: '浜田',
              firstname: '二郎',
              type: 1,
              grade: 2,
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

    describe('methods', () => {
      describe('getSchoolType', () => {
        it('type is elementary school', () => {
          expect(wrapper.vm.getSchoolType(SchoolType.ElementarySchool)).toBe('小学校')
        })

        it('type is junior high school', () => {
          expect(wrapper.vm.getSchoolType(SchoolType.JuniorHighSchool)).toBe('中学校')
        })

        it('type is high school', () => {
          expect(wrapper.vm.getSchoolType(SchoolType.HighSchool)).toBe('高等学校')
        })

        it('invalid type', () => {
          expect(wrapper.vm.getSchoolType(-1)).toBe('')
        })
      })

      describe('getSchoolTypeColor', () => {
        it('type is elementary school', () => {
          expect(wrapper.vm.getSchoolTypeColor(SchoolType.ElementarySchool)).toBe('primary')
        })

        it('type is junior high school', () => {
          expect(wrapper.vm.getSchoolTypeColor(SchoolType.JuniorHighSchool)).toBe('secondary')
        })

        it('type is high school', () => {
          expect(wrapper.vm.getSchoolTypeColor(SchoolType.HighSchool)).toBe('info')
        })

        it('invalid type', () => {
          expect(wrapper.vm.getSchoolTypeColor(-1)).toBe('')
        })
      })

      describe('getSchoolYear', () => {
        it('return', () => {
          expect(wrapper.vm.getSchoolYear(3)).toBe('3年')
        })
      })
    })
  })
})