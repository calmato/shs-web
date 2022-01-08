<template>
  <the-user-top
    :loading="loading"
    :is-admin="isAdmin"
    :subjects="subjects"
    :students="students"
    :teacher="teacher"
    :teacher-edit-dialog="teacherDialog"
    :teachers="teachers"
    :teachers-total="teachersTotal"
    :teachers-page.sync="teachersPage"
    :teachers-items-per-page.sync="teachersItemsPerPage"
    :edit-teacher-elementary-school-form="editTeacherElementarySchoolForm"
    :edit-teacher-junior-high-school-form="editTeacherJuniorHighSchoolForm"
    :edit-teacher-high-school-form="editTeacherHighSchoolForm"
    @click:new="handleClickNew"
    @click:show-teacher="handleClickShowTeacher"
    @click:close-teacher="handleCloseTeacherDialog"
    @submit:teacher-elementary-school="handleSubmitTeacherElementarySchool"
    @submit:teacher-junior-high-school="handleSubmitTeacherJuniorHighSchool"
    @submit:teacher-high-school="handleSubmitTeacherHighSchool"
    @submit:teacher-role="handleSubmitTeacherRole"
    @submit:teacher-delete="handleSubmitDeleteTeacher"
  />
</template>

<script lang="ts">
import { computed, defineComponent, reactive, ref, SetupContext, useAsync, watch } from '@nuxtjs/composition-api'
import TheUserTop from '~/components/templates/TheUserTop.vue'
import { CommonStore, UserStore } from '~/store'
import {
  TeacherEditSubjectForm,
  TeacherEditSubjectForElementarySchoolParams,
  TeacherEditSubjectForHighSchoolParams,
  TeacherEditSubjectForJuniorHighSchoolParams,
  TeacherEditRoleForm,
  TeacherEditRoleParams,
  TeacherEditRoleOptions,
  TeacherEditSubjectForHighSchoolOptions,
  TeacherEditSubjectForJuniorHighSchoolOptions,
  TeacherEditSubjectForElementarySchoolOptions,
} from '~/types/form'
import { PromiseState, Role, Student, Subject, SubjectsMap, Teacher } from '~/types/store'

export default defineComponent({
  components: {
    TheUserTop,
  },

  setup(_, { root }: SetupContext) {
    const router = root.$router
    const store = root.$store

    const teacherDialog = ref<boolean>(false)
    const teachersPage = ref<number>(1)
    const teachersItemsPerPage = ref<number>(10)

    const editTeacherElementarySchoolForm = reactive<TeacherEditSubjectForm>({
      params: TeacherEditSubjectForElementarySchoolParams,
      options: TeacherEditSubjectForElementarySchoolOptions,
    })

    const editTeacherJuniorHighSchoolForm = reactive<TeacherEditSubjectForm>({
      params: TeacherEditSubjectForJuniorHighSchoolParams,
      options: TeacherEditSubjectForJuniorHighSchoolOptions,
    })

    const editTeacherHighSchoolForm = reactive<TeacherEditSubjectForm>({
      params: TeacherEditSubjectForHighSchoolParams,
      options: TeacherEditSubjectForHighSchoolOptions,
    })

    const editTeacherRoleForm = reactive<TeacherEditRoleForm>({
      params: TeacherEditRoleParams,
      options: TeacherEditRoleOptions,
    })

    const loading = computed<boolean>(() => store.getters['common/getPromiseState'] === PromiseState.LOADING)
    const isAdmin = computed<boolean>(() => store.getters['auth/getRole'] === Role.ADMINISTRATOR)
    const subjects = computed<SubjectsMap>(() => store.getters['lesson/getSubjectsMap'])
    const students = computed<Student[]>(() => store.getters['user/getStudents'])
    const teacher = computed<Teacher>(() => store.getters['user/getTeacher'])
    const teachers = computed<Teacher[]>(() => store.getters['user/getTeachers'])
    const teachersTotal = computed<number>(() => store.getters['user/getTeachersTotal'])

    watch(teachersPage, async () => {
      await listTeachers()
    })

    watch(teachersItemsPerPage, async () => {
      await listTeachers()
    })

    useAsync(async () => {
      await listTeachers()
    })

    async function listTeachers(): Promise<void> {
      CommonStore.startConnection()

      const limit: number = teachersItemsPerPage.value
      const offset: number = (teachersPage.value - 1) * limit

      await UserStore.listTeachers({ limit, offset })
        .catch((err: Error) => {
          console.log('feiled to list teachers', err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    const getSubjectIds = (subjects: Subject[] | undefined): number[] => {
      const ids: number[] | undefined = subjects?.map((subject: Subject) => {
        return subject.id
      })
      return ids || []
    }

    const handleClickNew = (actor: string): void => {
      router.push(`/users/${actor}/new`)
    }

    const handleClickShowTeacher = async (teacher: Teacher): Promise<void> => {
      CommonStore.startConnection()

      const teacherId: string = teacher.id

      await UserStore.showTeacher({ teacherId })
        .then(() => {
          const teacher: Teacher = store.getters['user/getTeacher']
          const subjects: SubjectsMap = teacher.subjects
          editTeacherElementarySchoolForm.params.subjectIds = getSubjectIds(subjects['小学校'])
          editTeacherJuniorHighSchoolForm.params.subjectIds = getSubjectIds(subjects['中学校'])
          editTeacherHighSchoolForm.params.subjectIds = getSubjectIds(subjects['高校'])
          editTeacherRoleForm.params.role = teacher.role
          teacherDialog.value = true
        })
        .catch((err: Error) => {
          console.log('feiled to show teacher', err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    const handleSubmitTeacherElementarySchool = async (): Promise<void> => {
      CommonStore.startConnection()

      const teacherId: string = teacher.value.id

      await UserStore.updateTeacherSubjects({ teacherId, form: editTeacherElementarySchoolForm })
        .catch((err: Error) => {
          console.log('feiled to update teacher subjects', err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    const handleSubmitTeacherJuniorHighSchool = async (): Promise<void> => {
      CommonStore.startConnection()

      const teacherId: string = teacher.value.id

      await UserStore.updateTeacherSubjects({ teacherId, form: editTeacherJuniorHighSchoolForm })
        .catch((err: Error) => {
          console.log('feiled to update teacher subjects', err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    const handleSubmitTeacherHighSchool = async (): Promise<void> => {
      CommonStore.startConnection()

      const teacherId: string = teacher.value.id

      await UserStore.updateTeacherSubjects({ teacherId, form: editTeacherHighSchoolForm })
        .catch((err: Error) => {
          console.log('feiled to update teacher subjects', err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    const handleSubmitTeacherRole = async (): Promise<void> => {
      CommonStore.startConnection()

      const teacherId: string = teacher.value.id

      await UserStore.updateTeacherRole({ teacherId, form: editTeacherRoleForm })
        .catch((err: Error) => {
          console.log('feiled to update teacher role', err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    const handleSubmitDeleteTeacher = async (): Promise<void> => {
      CommonStore.startConnection()

      const teacherId: string = teacher.value.id

      await UserStore.deleteTeacher({ teacherId })
        .then(() => {
          teacherDialog.value = false
        })
        .catch((err: Error) => {
          console.log('feiled to delete teacher', err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    const handleCloseTeacherDialog = (): void => {
      teacherDialog.value = false
    }

    return {
      loading,
      isAdmin,
      subjects,
      students,
      teacher,
      teacherDialog,
      teachers,
      teachersTotal,
      teachersPage,
      teachersItemsPerPage,
      editTeacherElementarySchoolForm,
      editTeacherJuniorHighSchoolForm,
      editTeacherHighSchoolForm,
      editTeacherRoleForm,
      handleClickNew,
      handleClickShowTeacher,
      handleSubmitTeacherElementarySchool,
      handleSubmitTeacherJuniorHighSchool,
      handleSubmitTeacherHighSchool,
      handleSubmitTeacherRole,
      handleSubmitDeleteTeacher,
      handleCloseTeacherDialog,
    }
  },
})
</script>
