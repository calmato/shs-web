<template>
  <the-setting-top
    class="px-2"
    :user-items="userItems"
    :system-items="systemItems"
    :user="auth"
    :elementary-school-subjects="elementarySchoolSubjects"
    :junior-high-school-subjects="juniorHighSchoolSubjects"
    :high-school-subjects="highSchoolSubjects"
    :elementary-school-subjects-form-value.sync="elementarySchoolSubjectForm.subjectIDs"
    :junior-high-school-subjects-form-value.sync="juniorHighSchoolSubjectForm.subjectIDs"
    :high-school-subjects-form-value.sync="highSchoolSubjectForm.subjectIDs"
    @handleElementarySchoolSubjectsBlur="handleElementarySchoolSubjectsBlur"
    @handleJuniorHighSchoolSubjectsBlur="handleJuniorHighSchoolSubjectsBlur"
    @handleHighSchoolSubjectsBlur="handleHighSchoolSubjectsBlur"
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
import { Auth, SchoolTypeArray, Subject } from '~/types/store'

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
        title: '定休日/コマ/ブース 設定',
        path: '/settings/classroom',
      },
    ]

    const elementarySchoolSubjectForm = reactive<SubjectUpdateForm>({
      schoolType: 1,
      subjectIDs: [],
    })

    const juniorHighSchoolSubjectForm = reactive<SubjectUpdateForm>({
      schoolType: 2,
      subjectIDs: [],
    })

    const highSchoolSubjectForm = reactive<SubjectUpdateForm>({
      schoolType: 3,
      subjectIDs: [],
    })

    const auth = computed<Auth>(() => store.getters['auth/getAuth'])

    const subjects = computed<Subject[]>(() => store.getters['lesson/getSubjects'])
    const elementarySchoolSubjects = computed<Subject[]>(() =>
      subjects.value.filter((item) => item.schoolType === '小学校')
    )
    const juniorHighSchoolSubjects = computed<Subject[]>(() =>
      subjects.value.filter((item) => item.schoolType === '中学校')
    )
    const highSchoolSubjects = computed<Subject[]>(() => subjects.value.filter((item) => item.schoolType === '高校'))

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

    const handleElementarySchoolSubjectsBlur = (_val: number[]) => {
      AuthStore.updateOwnSubjects(elementarySchoolSubjectForm)
    }

    const handleJuniorHighSchoolSubjectsBlur = (_val: number[]) => {
      AuthStore.updateOwnSubjects(juniorHighSchoolSubjectForm)
    }

    const handleHighSchoolSubjectsBlur = (_val: number[]) => {
      AuthStore.updateOwnSubjects(highSchoolSubjectForm)
    }

    onMounted(() => {
      const defaultSubjects = auth.value.subjects
      SchoolTypeArray.forEach((schoolType) => {
        if (schoolType === 'その他') {
          return
        }
        const values: number[] = defaultSubjects[schoolType].map((itme) => itme.id)
        switch (schoolType) {
          case '小学校':
            elementarySchoolSubjectForm.subjectIDs = values
            break
          case '中学校':
            juniorHighSchoolSubjectForm.subjectIDs = values
            break
          case '高校':
            highSchoolSubjectForm.subjectIDs = values
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
      handleElementarySchoolSubjectsBlur,
      handleJuniorHighSchoolSubjectsBlur,
      handleHighSchoolSubjectsBlur,
    }
  },
})
</script>
