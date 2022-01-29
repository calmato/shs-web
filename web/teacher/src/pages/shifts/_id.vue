<template>
  <div>
    <!-- PCレイアウト -->
    <pc-shift-detail
      class="hidden-sm-and-down"
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
      @click:show-teacher-submissions="handleClickShowTeacherSubmissions"
      @click:show-teacher-lessons="handleClickShowTeacherLessons"
      @click:show-student-submissions="handleClickShowStudentSubmissions"
      @click:show-student-lessons="handleClickShowStudentLessons"
      @click:decided-lesson="handleClickDecidedLesson"
      @click:new-lesson="handleClickNewLesson"
      @click:edit-lesson="handleClickEditLesson"
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

    const dialogKey = ref<ShiftDialogKey>('未選択')
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

    useAsync(async () => {
      await listShiftDetails()
    })

    async function listShiftDetails(): Promise<void> {
      CommonStore.startConnection()

      const summaryId: string = route.value.params.id

      await ShiftStore.listShiftDetails({ summaryId: Number(summaryId) })
        .catch((err: Error) => {
          console.log('feiled to list shift details', err)
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

      const summaryId: string = route.value.params.id

      await ShiftStore.listShiftLessons({ summaryId: Number(summaryId), lessonId, shiftId, room })
        .then(() => {
          if (lesson.value.current) {
            form.params = { ...lesson.value.current, summaryId: 0 }
          } else {
            form.params = { ...ShiftLessonParams }
          }
        })
        .catch((err: Error) => {
          console.log('feiled to list shift lessons', err)
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

    const handleClickShowTeacherSubmissions = async (teacherId: string): Promise<void> => {
      CommonStore.startConnection()

      const summaryId: string = route.value.params.id

      await ShiftStore.showTeacherSubmissions({ summaryId: Number(summaryId), teacherId })
        .then(() => {
          openDialog('講師シフト')
        })
        .catch((err: Error) => {
          console.log('feiled to show teacher submissions', err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    const handleClickShowTeacherLessons = async (teacherId: string): Promise<void> => {
      CommonStore.startConnection()

      const summaryId: string = route.value.params.id

      await ShiftStore.listTeacherLessons({ summaryId: Number(summaryId), teacherId })
        .then(() => {
          openDialog('講師授業')
        })
        .catch((err: Error) => {
          console.log('feiled to list teacher lessons', err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    const handleClickShowStudentSubmissions = async (studentId: string): Promise<void> => {
      CommonStore.startConnection()

      const summaryId: string = route.value.params.id

      await ShiftStore.showStudentSubmissions({ summaryId: Number(summaryId), studentId })
        .then(() => {
          openDialog('生徒授業希望')
        })
        .catch((err: Error) => {
          console.log('feiled to show teacher submissions', err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    const handleClickShowStudentLessons = async (studentId: string): Promise<void> => {
      CommonStore.startConnection()

      const summaryId: string = route.value.params.id

      await ShiftStore.listStudentLessons({ summaryId: Number(summaryId), studentId })
        .then(() => {
          openDialog('生徒授業')
        })
        .catch((err: Error) => {
          console.log('feiled to list student lessons', err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

    const handleClickDecidedLesson = (): void => {
      console.log('debug', 'decided lesson')
    }

    const handleClickNewLesson = async ({ shiftId, room }: { shiftId: number; room: number }): Promise<void> => {
      await listShiftLessons({ shiftId, room, lessonId: 0 }).then(() => {
        openDialog('授業登録')
      })
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
      await listShiftLessons({ shiftId, room, lessonId }).then(() => {
        openDialog('授業登録')
      })
    }

    const handleCloseDialog = (): void => {
      closeDialog()
    }

    return {
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
      handleClickShowTeacherSubmissions,
      handleClickShowTeacherLessons,
      handleClickShowStudentSubmissions,
      handleClickShowStudentLessons,
      handleClickDecidedLesson,
      handleClickNewLesson,
      handleClickEditLesson,
      handleCloseDialog,
    }
  },
})
</script>
