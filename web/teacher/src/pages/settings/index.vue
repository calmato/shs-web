<template>
  <the-setting-top
    class="px-2"
    :user-items="userItems"
    :system-items="systemItems"
    :user="auth"
    :elementary-school-subjects="elementarySchoolSubjects"
    :junior-high-school-subjects="juniorHighSchoolSubjects"
    :high-school-subjects="highSchoolSubjects"
    :elementary-school-subjects-form-value.sync="elementarySchoolSubjectForm.subjectIds"
    :junior-high-school-subjects-form-value.sync="juniorHighSchoolSubjectForm.subjectIds"
    :high-school-subjects-form-value.sync="highSchoolSubjectForm.subjectIds"
    @handleElementarySchoolSubjectsChange="handleElementarySchoolSubjectsChange"
    @handleJuniorHighSchoolSubjectsChange="handleJuniorHighSchoolSubjectsChange"
    @handleHighSchoolSubjectsChange="handleHighSchoolSubjectsChange"
    @click="handleClick"
    @onClickBackButton="handleBackButton"
  />
</template>

<script lang="ts">
import { computed, defineComponent, onMounted, reactive, SetupContext } from '@nuxtjs/composition-api'
import TheSettingTop from '~/components/templates/TheSettingTop.vue'
import { AuthStore } from '~/store'
import { SubjectUpdateForm } from '~/types/form'
import { Menu } from '~/types/props/setting'
import { Auth, SchoolType, Subject } from '~/types/store'

export default defineComponent({
  components: {
    TheSettingTop,
  },

  setup(_, { root }: SetupContext) {
    const router = root.$router
    const store = root.$store

    const userItems: Menu[] = [
      {
        title: 'パスワードの変更',
        path: '#パスワードの変更',
      },
      {
        title: 'サインアウト',
        path: '/signout',
      },
    ]
    const systemItems: Menu[] = [
      {
        title: '教室・科目設定',
        path: '#教室・科目設定',
      },
      {
        title: 'コマ設定',
        path: '#コマ設定',
      },
    ]

    const elementarySchoolSubjectForm = reactive<SubjectUpdateForm>({
      schoolType: SchoolType.ELEMENTARY_SCHOOL,
      subjectIds: [],
    })

    const juniorHighSchoolSubjectForm = reactive<SubjectUpdateForm>({
      schoolType: SchoolType.JUNIOR_HIGH_SCHOOL,
      subjectIds: [],
    })

    const highSchoolSubjectForm = reactive<SubjectUpdateForm>({
      schoolType: SchoolType.HIGH_SCHOOL,
      subjectIds: [],
    })

    const auth = computed<Auth>(() => store.getters['auth/getAuth'])

    const subjects = computed<Subject[]>(() => store.getters['lesson/getSubjects'])
    const elementarySchoolSubjects = computed<Subject[]>(() => subjects.value.filter((item) => item.schoolType === 1))
    const juniorHighSchoolSubjects = computed<Subject[]>(() => subjects.value.filter((item) => item.schoolType === 2))
    const highSchoolSubjects = computed<Subject[]>(() => subjects.value.filter((item) => item.schoolType === 3))

    const handleClick = (item: Menu): void => {
      if (item.path === '/signout') {
        AuthStore.signOut()
        router.push('/signin')
        return
      }
      router.push(item.path)
    }

    const handleBackButton = (): void => {
      router.back()
    }

    const handleElementarySchoolSubjectsChange = (_val: number[]) => {
      AuthStore.updateOwnSubjects(elementarySchoolSubjectForm)
    }

    const handleJuniorHighSchoolSubjectsChange = (_val: number[]) => {
      AuthStore.updateOwnSubjects(juniorHighSchoolSubjectForm)
    }

    const handleHighSchoolSubjectsChange = (_val: number[]) => {
      AuthStore.updateOwnSubjects(highSchoolSubjectForm)
    }

    onMounted(() => {
      const defaultSubjects = auth.value.subjects
      Object.keys(defaultSubjects).forEach((schoolTypeString: string) => {
        const schoolType = Number(schoolTypeString) as
          | SchoolType.ELEMENTARY_SCHOOL
          | SchoolType.JUNIOR_HIGH_SCHOOL
          | SchoolType.HIGH_SCHOOL
        const _v = [1, 2, 3].includes(schoolType) ? defaultSubjects[schoolType] : undefined
        const value = typeof _v !== 'undefined' ? _v.map((item) => item.id) : []
        switch (schoolType) {
          case SchoolType.ELEMENTARY_SCHOOL:
            elementarySchoolSubjectForm.subjectIds = value
            break
          case SchoolType.JUNIOR_HIGH_SCHOOL:
            juniorHighSchoolSubjectForm.subjectIds = value
            break
          case SchoolType.HIGH_SCHOOL:
            highSchoolSubjectForm.subjectIds = value
            break
          default:
            break
        }
      })
    })

    return {
      userItems,
      systemItems,
      handleClick,
      handleBackButton,
      auth,
      elementarySchoolSubjects,
      juniorHighSchoolSubjects,
      highSchoolSubjects,
      elementarySchoolSubjectForm,
      juniorHighSchoolSubjectForm,
      highSchoolSubjectForm,
      handleElementarySchoolSubjectsChange,
      handleJuniorHighSchoolSubjectsChange,
      handleHighSchoolSubjectsChange,
    }
  },
})
</script>
