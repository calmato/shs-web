<template>
  <the-setting-top
    class="px-2 mt-4"
    :menu-items="menuItems"
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
  />
</template>

<script lang="ts">
import { computed, defineComponent, onMounted, reactive, useRouter, useStore } from '@nuxtjs/composition-api'
import TheSettingTop from '~/components/templates/TheSettingTop.vue'
import { AuthStore } from '~/store'
import { SubjectUpdateForm } from '~/types/form'
import { Menu } from '~/types/props/setting'
import { Auth, SchoolTypeArray, Subject } from '~/types/store'

export default defineComponent({
  components: {
    TheSettingTop,
  },

  setup() {
    const router = useRouter()
    const store = useStore()

    const menuItems: Menu[] = [
      {
        title: 'メールアドレス変更',
        path: '/settings/mail',
      },
      {
        title: 'パスワードの変更',
        path: '/settings/password',
      },
      {
        title: '開講科目設定',
        path: '/settings/subjects',
      },
      {
        title: '定休日/コマ/ブース 設定',
        path: '/settings/classroom',
      },
      {
        title: 'サインアウト',
        path: '/signout',
        textColor: 'error',
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
      menuItems,
      handleClick,
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
