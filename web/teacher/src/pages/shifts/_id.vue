<template>
  <div>
    <!-- PCレイアウト -->
    <pc-shift-detail
      class="hidden-sm-and-down"
      :overlay="overlay"
      :loading="loading"
      :lesson-loading="lessonLoading"
      :dialog="dialog"
      :dialog-key="dialogKey"
      :summary="summary"
      :details="details"
      :rooms="rooms"
      :teachers="teachers"
      :students="students"
      :lesson="lesson"
      :lessons="getLessonDetails()"
      :subjects="subjects"
      :teacher-submission="teacherSubmission"
      :teacher-lessons="teacherLessons"
      :student-submission="studentSubmission"
      :student-lessons="studentLessons"
      :form="form"
      @click:lesson-student="handleClickLessonStudent"
      @click:show-teacher-submissions="handleClickShowTeacherSubmissions"
      @click:show-teacher-lessons="handleClickShowTeacherLessons"
      @click:show-student-submissions="handleClickShowStudentSubmissions"
      @click:show-student-lessons="handleClickShowStudentLessons"
      @click:decided-lesson="handleClickDecidedLesson"
      @click:new-lesson="handleClickNewLesson"
      @click:edit-lesson="handleClickEditLesson"
      @click:submit-lesson="handleClickSubmitLesson"
      @click:delete-lesson="handleClickDeleteLesson"
      @click:close="handleCloseDialog"
    />
    <!-- スマホレイアウト -->
    <mb-shift-detail class="hidden-md-and-up" @click="handleClickTop" />
  </div>
</template>

<script lang="ts">
import {
  computed,
  defineComponent,
  reactive,
  ref,
  useAsync,
  useRoute,
  useRouter,
  useStore,
} from '@nuxtjs/composition-api'
import MbShiftDetail from '~/components/templates/MbShiftDetail.vue'
import PcShiftDetail from '~/components/templates/PcShiftDetail.vue'
import { CommonStore, ShiftStore } from '~/store'
import {
  Lesson,
  PromiseState,
  ShiftDetail,
  ShiftLessonDetail,
  ShiftSummary,
  ShiftUserLesson,
  StudentShift,
  StudentSubmissionDetail,
  Subject,
  TeacherShift,
  TeacherSubmissionDetail,
} from '~/types/store'
import { LessonDetail, ShiftDialogKey } from '~/types/props/shift'
import { ShiftLessonForm, ShiftLessonParams } from '~/types/form'

export default defineComponent({
  components: {
    MbShiftDetail,
    PcShiftDetail,
  },

  setup() {
    const router = useRouter()
    const route = useRoute()
    const store = useStore()

    const summaryId = Number(route.value.params.id)

    const overlay = ref<boolean>(true)
    const dialogKey = ref<ShiftDialogKey>('未選択')
    const lessonLoading = ref<boolean>(false)
    const form = reactive<ShiftLessonForm>({ params: { ...ShiftLessonParams } })

    const dialog = computed<boolean>(() => dialogKey.value !== '未選択')
    const summary = computed<ShiftSummary>(() => store.getters['shift/getSummary'])
    const details = computed<ShiftDetail[]>(() => store.getters['shift/getDetails'])
    const rooms = computed<number>(() => store.getters['shift/getRooms'])
    const teachers = computed<TeacherShift[]>(() => store.getters['shift/getTeachers'])
    const students = computed<StudentShift[]>(() => store.getters['shift/getStudents'])
    const lesson = computed<ShiftLessonDetail>(() => store.getters['shift/getLessonDetail'])
    const lessons = computed<Lesson[]>(() => store.getters['shift/getLessons'])
    const subjects = computed<Subject[]>(() => store.getters['lesson/getSubjects'])
    const teacherSubmission = computed<TeacherSubmissionDetail>(() => store.getters['shift/getTeacherSubmission'])
    const teacherLessons = computed<ShiftUserLesson>(() => store.getters['shift/getTeacherLessons'])
    const studentSubmission = computed<StudentSubmissionDetail>(() => store.getters['shift/getStudentSubmission'])
    const studentLessons = computed<ShiftUserLesson>(() => store.getters['shift/getStudentLessons'])
    const loading = computed<boolean>(() => {
      return store.getters['common/getPromiseState'] === PromiseState.LOADING
    })

    useAsync(async () => {
      await listShiftDetails()
      overlay.value = false
    })

    async function listShiftDetails(): Promise<void> {
      CommonStore.startConnection()

      await ShiftStore.listShiftDetails({ summaryId })
        .catch((err: Error) => {
          CommonStore.showErrorInSnackbar(err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    async function listShiftLessons({
      shiftId,
      lessonId,
      room,
    }: {
      shiftId: number
      lessonId: number
      room: number
    }): Promise<void> {
      CommonStore.startConnection()
      openDialog('授業登録')

      await ShiftStore.listShiftLessons({ summaryId, lessonId, shiftId, room })
        .then(() => {
          if (lesson.value.current) {
            form.params = { ...lesson.value.current, summaryId, lessonId, shiftId, room }
          } else {
            form.params = { ...ShiftLessonParams, summaryId, lessonId, shiftId, room }
          }
        })
        .catch((err: Error) => {
          CommonStore.showErrorInSnackbar(err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    const getLessonDetails = (): LessonDetail[] => {
      const details: LessonDetail[] = lessons.value.map((lesson: Lesson) => {
        const teacher: TeacherShift | undefined = teachers.value.find(
          (teacher: TeacherShift): boolean => teacher.id === lesson.teacherId
        )
        const student: StudentShift | undefined = students.value.find(
          (student: StudentShift): boolean => student.id === lesson.studentId
        )
        const subject: Subject | undefined = subjects.value.find(
          (subject: Subject): boolean => subject.id === lesson.subjectId
        )
        return {
          lesson,
          teacher,
          student,
          subject,
        }
      })
      return details
    }

    const openDialog = (key: ShiftDialogKey): void => {
      dialogKey.value = key
    }

    const closeDialog = (): void => {
      dialogKey.value = '未選択'
    }

    const handleClickTop = (): void => {
      router.push('/')
    }

    const handleClickLessonStudent = async (studentId: string): Promise<void> => {
      lessonLoading.value = true

      await ShiftStore.listStudentLessons({ summaryId, studentId })
        .catch((err: Error) => {
          CommonStore.showErrorInSnackbar(err)
        })
        .finally(() => {
          lessonLoading.value = false
          form.params.studentId = studentId
        })
    }

    const handleClickShowTeacherSubmissions = async (teacherId: string): Promise<void> => {
      CommonStore.startConnection()
      openDialog('講師シフト')

      await ShiftStore.showTeacherSubmissions({ summaryId, teacherId })
        .catch((err: Error) => {
          CommonStore.showErrorInSnackbar(err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    const handleClickShowTeacherLessons = async (teacherId: string): Promise<void> => {
      CommonStore.startConnection()
      openDialog('講師授業')

      await ShiftStore.listTeacherLessons({ summaryId, teacherId })
        .catch((err: Error) => {
          CommonStore.showErrorInSnackbar(err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    const handleClickShowStudentSubmissions = async (studentId: string): Promise<void> => {
      CommonStore.startConnection()
      openDialog('生徒授業希望')

      await ShiftStore.showStudentSubmissions({ summaryId, studentId })
        .catch((err: Error) => {
          CommonStore.showErrorInSnackbar(err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    const handleClickShowStudentLessons = async (studentId: string): Promise<void> => {
      CommonStore.startConnection()
      openDialog('生徒授業')

      await ShiftStore.listStudentLessons({ summaryId, studentId })
        .catch((err: Error) => {
          CommonStore.showErrorInSnackbar(err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    const handleClickDecidedLesson = async (): Promise<void> => {
      CommonStore.startConnection()

      const decided: boolean = !summary.value.decided

      await ShiftStore.updateShiftSummaryDecided({ summaryId, decided })
        .then(() => {
          if (decided) {
            CommonStore.showSnackbar({ color: 'success', message: '授業スケジュールを確定しました。' })
          }
        })
        .catch((err: Error) => {
          CommonStore.showErrorInSnackbar(err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    const handleClickNewLesson = async ({ shiftId, room }: { shiftId: number; room: number }): Promise<void> => {
      await listShiftLessons({ shiftId, room, lessonId: 0 })
    }

    const handleClickEditLesson = async ({
      shiftId,
      lessonId,
      room,
    }: {
      shiftId: number
      lessonId: number
      room: number
    }): Promise<void> => {
      await listShiftLessons({ shiftId, room, lessonId })
    }

    const handleClickSubmitLesson = async (): Promise<void> => {
      CommonStore.startConnection()

      await ShiftStore.upsertLesson({ summaryId, form })
        .then(() => {
          closeDialog()
          CommonStore.showSnackbar({ color: 'success', message: '授業を登録しました。' })
        })
        .catch((err: Error) => {
          CommonStore.showErrorInSnackbar(err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    const handleClickDeleteLesson = async (): Promise<void> => {
      CommonStore.startConnection()

      await ShiftStore.deleteLesson({ summaryId, form })
        .then(() => {
          closeDialog()
          CommonStore.showSnackbar({ color: 'success', message: '授業を削除しました。' })
        })
        .catch((err: Error) => {
          CommonStore.showErrorInSnackbar(err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    const handleCloseDialog = (): void => {
      closeDialog()
    }

    return {
      overlay,
      loading,
      lessonLoading,
      dialog,
      dialogKey,
      summary,
      details,
      rooms,
      teachers,
      students,
      lesson,
      lessons,
      subjects,
      teacherSubmission,
      teacherLessons,
      studentSubmission,
      studentLessons,
      form,
      getLessonDetails,
      handleClickTop,
      handleClickLessonStudent,
      handleClickShowTeacherSubmissions,
      handleClickShowTeacherLessons,
      handleClickShowStudentSubmissions,
      handleClickShowStudentLessons,
      handleClickDecidedLesson,
      handleClickNewLesson,
      handleClickEditLesson,
      handleClickSubmitLesson,
      handleClickDeleteLesson,
      handleCloseDialog,
    }
  },
})
</script>
