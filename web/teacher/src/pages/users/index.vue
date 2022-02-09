<template>
  <the-user-top
    :loading="loading"
    :is-admin="isAdmin"
    :subjects="subjects"
    :student="student"
    :students="students"
    :student-edit-dialog="studentDialog"
    :students-total="studentsTotal"
    :students-page.sync="studentsPage"
    :students-items-per-page.sync="studentsItemsPerPage"
    :teacher="teacher"
    :teacher-edit-dialog="teacherDialog"
    :teachers="teachers"
    :teachers-total="teachersTotal"
    :teachers-page.sync="teachersPage"
    :teachers-items-per-page.sync="teachersItemsPerPage"
    :edit-student-elementary-school-form="editStudentElementarySchoolForm"
    :edit-student-junior-high-school-form="editStudentJuniorHighSchoolForm"
    :edit-student-high-school-form="editStudentHighSchoolForm"
    :edit-teacher-elementary-school-form="editTeacherElementarySchoolForm"
    :edit-teacher-junior-high-school-form="editTeacherJuniorHighSchoolForm"
    :edit-teacher-high-school-form="editTeacherHighSchoolForm"
    @click:new="handleClickNew"
    @click:show-student="handleClickShowStudent"
    @click:show-teacher="handleClickShowTeacher"
    @click:close-student="handleCloseStudentDialog"
    @click:close-teacher="handleCloseTeacherDialog"
    @submit:student-elementary-school="handleSubmitStudentElementarySchool"
    @submit:student-junior-high-school="handleSubmitStudentJuniorHighSchool"
    @submit:student-high-school="handleSubmitStudentHighSchool"
    @submit:student-delete="handleSubmitDeleteStudent"
    @submit:teacher-elementary-school="handleSubmitTeacherElementarySchool"
    @submit:teacher-junior-high-school="handleSubmitTeacherJuniorHighSchool"
    @submit:teacher-high-school="handleSubmitTeacherHighSchool"
    @submit:teacher-role="handleSubmitTeacherRole"
    @submit:teacher-delete="handleSubmitDeleteTeacher"
  />
</template>

<script lang="ts">
import { computed, defineComponent, reactive, ref, useAsync, watch, useRouter, useStore } from '@nuxtjs/composition-api'
import TheUserTop from '~/components/templates/TheUserTop.vue'
import { CommonStore, UserStore } from '~/store'
import {
  StudentEditSubjectForm,
  StudentEditSubjectForElementarySchoolOptions,
  StudentEditSubjectForElementarySchoolParams,
  StudentEditSubjectForJuniorHighSchoolOptions,
  StudentEditSubjectForJuniorHighSchoolParams,
  StudentEditSubjectForHighSchoolOptions,
  StudentEditSubjectForHighSchoolParams,
  TeacherEditSubjectForm,
  TeacherEditSubjectForElementarySchoolOptions,
  TeacherEditSubjectForElementarySchoolParams,
  TeacherEditSubjectForJuniorHighSchoolOptions,
  TeacherEditSubjectForJuniorHighSchoolParams,
  TeacherEditSubjectForHighSchoolOptions,
  TeacherEditSubjectForHighSchoolParams,
  TeacherEditRoleForm,
  TeacherEditRoleOptions,
  TeacherEditRoleParams,
} from '~/types/form'
import { PromiseState, Role, Student, Subject, SubjectsMap, Teacher } from '~/types/store'

export default defineComponent({
  components: {
    TheUserTop,
  },

  setup() {
    const router = useRouter()
    const store = useStore()

    const studentDialog = ref<boolean>(false)
    const studentsPage = ref<number>(1)
    const studentsItemsPerPage = ref<number>(10)
    const teacherDialog = ref<boolean>(false)
    const teachersPage = ref<number>(1)
    const teachersItemsPerPage = ref<number>(10)

    const editStudentElementarySchoolForm = reactive<StudentEditSubjectForm>({
      params: StudentEditSubjectForElementarySchoolParams,
      options: StudentEditSubjectForElementarySchoolOptions,
    })

    const editStudentJuniorHighSchoolForm = reactive<StudentEditSubjectForm>({
      params: StudentEditSubjectForJuniorHighSchoolParams,
      options: StudentEditSubjectForJuniorHighSchoolOptions,
    })

    const editStudentHighSchoolForm = reactive<StudentEditSubjectForm>({
      params: StudentEditSubjectForHighSchoolParams,
      options: StudentEditSubjectForHighSchoolOptions,
    })

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
    const student = computed<Student>(() => store.getters['user/getStudent'])
    const students = computed<Student[]>(() => store.getters['user/getStudents'])
    const studentsTotal = computed<number>(() => store.getters['user/getStudentsTotal'])
    const teacher = computed<Teacher>(() => store.getters['user/getTeacher'])
    const teachers = computed<Teacher[]>(() => store.getters['user/getTeachers'])
    const teachersTotal = computed<number>(() => store.getters['user/getTeachersTotal'])

    watch(studentsPage, async () => {
      await listStudents()
    })

    watch(studentsItemsPerPage, async () => {
      await listStudents()
    })

    watch(teachersPage, async () => {
      await listTeachers()
    })

    watch(teachersItemsPerPage, async () => {
      await listTeachers()
    })

    useAsync(async () => {
      await listStudents()
    })

    useAsync(async () => {
      await listTeachers()
    })

    async function listStudents(): Promise<void> {
      CommonStore.startConnection()

      const limit: number = studentsItemsPerPage.value
      const offset: number = (studentsPage.value - 1) * limit

      await UserStore.listStudents({ limit, offset })
        .catch((err: Error) => {
          console.log('feiled to list students', err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

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

    /**
     * 生徒関連
     */
    const handleClickShowStudent = async (student: Student): Promise<void> => {
      CommonStore.startConnection()

      const studentId: string = student.id

      await UserStore.showStudent({ studentId })
        .then(() => {
          const student: Student = store.getters['user/getStudent']
          const subjects: SubjectsMap = {
            小学校: [],
            中学校: [],
            高校: [],
            その他: [],
          }
          student.subjects.forEach((subject: Subject): void => {
            subjects[subject.schoolType].push(subject)
          })
          editStudentElementarySchoolForm.params.subjectIds = getSubjectIds(subjects['小学校'])
          editStudentJuniorHighSchoolForm.params.subjectIds = getSubjectIds(subjects['中学校'])
          editStudentHighSchoolForm.params.subjectIds = getSubjectIds(subjects['高校'])
          studentDialog.value = true
        })
        .catch((err: Error) => {
          console.log('feiled to show student', err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    const handleSubmitStudentElementarySchool = async (): Promise<void> => {
      CommonStore.startConnection()

      const studentId: string = student.value.id

      await UserStore.updateStudentSubjects({ studentId, form: editStudentElementarySchoolForm })
        .catch((err: Error) => {
          console.log('feiled to update student subjects', err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    const handleSubmitStudentJuniorHighSchool = async (): Promise<void> => {
      CommonStore.startConnection()

      const studentId: string = student.value.id

      await UserStore.updateStudentSubjects({ studentId, form: editStudentJuniorHighSchoolForm })
        .catch((err: Error) => {
          console.log('feiled to update student subjects', err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    const handleSubmitStudentHighSchool = async (): Promise<void> => {
      CommonStore.startConnection()

      const studentId: string = student.value.id

      await UserStore.updateStudentSubjects({ studentId, form: editStudentHighSchoolForm })
        .catch((err: Error) => {
          console.log('feiled to update student subjects', err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    const handleSubmitDeleteStudent = async (): Promise<void> => {
      CommonStore.startConnection()

      const studentId: string = student.value.id

      await UserStore.deleteStudent({ studentId })
        .then(() => {
          studentDialog.value = false
        })
        .catch((err: Error) => {
          console.log('feiled to delete student', err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    const handleCloseStudentDialog = (): void => {
      studentDialog.value = false
    }

    /**
     * 生徒関連
     */
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
      student,
      students,
      studentDialog,
      studentsTotal,
      studentsPage,
      studentsItemsPerPage,
      teacher,
      teacherDialog,
      teachers,
      teachersTotal,
      teachersPage,
      teachersItemsPerPage,
      editStudentElementarySchoolForm,
      editStudentJuniorHighSchoolForm,
      editStudentHighSchoolForm,
      editTeacherElementarySchoolForm,
      editTeacherJuniorHighSchoolForm,
      editTeacherHighSchoolForm,
      editTeacherRoleForm,
      handleClickNew,
      handleClickShowStudent,
      handleClickShowTeacher,
      handleSubmitStudentElementarySchool,
      handleSubmitStudentJuniorHighSchool,
      handleSubmitStudentHighSchool,
      handleSubmitTeacherElementarySchool,
      handleSubmitTeacherJuniorHighSchool,
      handleSubmitTeacherHighSchool,
      handleSubmitTeacherRole,
      handleSubmitDeleteStudent,
      handleSubmitDeleteTeacher,
      handleCloseStudentDialog,
      handleCloseTeacherDialog,
    }
  },
})
</script>
