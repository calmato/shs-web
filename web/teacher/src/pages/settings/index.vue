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
    @click="handleClick"
  />
</template>

<script lang="ts">
import { computed, defineComponent, reactive, SetupContext } from '@nuxtjs/composition-api'
import TheSettingTop from '~/components/templates/TheSettingTop.vue'
import { AuthStore } from '~/store'
import { SubjectUpdateForm } from '~/types/form'
import { Menu } from '~/types/props/setting'
import { AuthState, SchoolType, Subject } from '~/types/store'

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
      subjectIDs: [],
    })

    const juniorHighSchoolSubjectForm = reactive<SubjectUpdateForm>({
      schoolType: SchoolType.JUNIOR_HIGH_SCHOOL,
      subjectIDs: [],
    })

    const highSchoolSubjectForm = reactive<SubjectUpdateForm>({
      schoolType: SchoolType.HIGH_SCHOOL,
      subjectIDs: [],
    })

    const auth = computed<AuthState>(() => store.getters['auth/getAuth'])

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

    return {
      userItems,
      systemItems,
      handleClick,
      auth,
      elementarySchoolSubjects,
      juniorHighSchoolSubjects,
      highSchoolSubjects,
      elementarySchoolSubjectForm,
      juniorHighSchoolSubjectForm,
      highSchoolSubjectForm,
    }
  },
})
</script>
