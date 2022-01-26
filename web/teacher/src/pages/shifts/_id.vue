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
      :teacher-submission="teacherSubmission"
      :student-submission="studentSubmission"
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
import { computed, defineComponent, ref, useAsync, useRoute, useRouter, useStore } from '@nuxtjs/composition-api'
import MbShiftDetail from '~/components/templates/MbShiftDetail.vue'
import PcShiftDetail from '~/components/templates/PcShiftDetail.vue'
import { CommonStore, ShiftStore } from '~/store'
import {
  Lesson,
  ShiftDetail,
  ShiftLessonDetail,
  ShiftSummary,
  StudentShift,
  StudentSubmissionDetail,
  Subject,
  TeacherShift,
  TeacherSubmissionDetail,
} from '~/types/store'
import { LessonDetail, ShiftDialogKey } from '~/types/props/shift'

export default defineComponent({
  components: {
    MbShiftDetail,
    PcShiftDetail,
  },

  setup() {
    const router = useRouter()
    const route = useRoute()
    const store = useStore()

    const dialog = ref<boolean>(false)
    const dialogKey = ref<ShiftDialogKey>('未選択')

    const summary = computed<ShiftSummary>(() => store.getters['shift/getSummary'])
    const details = computed<ShiftDetail[]>(() => store.getters['shift/getDetails'])
    const rooms = computed<number>(() => store.getters['shift/getRooms'])
    const teachers = computed<TeacherShift[]>(() => store.getters['shift/getTeachers'])
    const students = computed<StudentShift[]>(() => store.getters['shift/getStudents'])
    const lesson = computed<ShiftLessonDetail>(() => store.getters['shift/getLessonDetail'])
    const lessons = computed<Lesson[]>(() => store.getters['shift/getLessons'])
    const subjects = computed<Subject[]>(() => store.getters['lesson/getSubjects'])
    const teacherSubmission = computed<TeacherSubmissionDetail>(() => store.getters['shift/getTeacherSubmission'])
    const studentSubmission = computed<StudentSubmissionDetail>(() => store.getters['shift/getStudentSubmission'])

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

    async function listShiftLessons({ shiftId, lessonId, room }: { shiftId: number; lessonId: number; room: number }): Promise<void> {
      CommonStore.startConnection()

      const summaryId: string = route.value.params.id

      await ShiftStore.listShiftLessons({ summaryId: Number(summaryId), lessonId, shiftId, room })
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
      dialog.value = true
      dialogKey.value = key
    }

    const closeDialog = (): void => {
      dialog.value = false
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

    const handleClickShowTeacherLessons = (teacherId: string): void => {
      console.log('debug', 'show teacher lessons', teacherId)
      openDialog('講師授業')
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

    const handleClickShowStudentLessons = (studentId: string): void => {
      console.log('debug', 'show student lessons', studentId)
      openDialog('生徒授業')
    }

    const handleClickDecidedLesson = (): void => {
      console.log('debug', 'decided lesson')
    }

    const handleClickNewLesson = async ({ shiftId, room }: { shiftId: number; room: number }): Promise<void> => {
      await listShiftLessons({ shiftId, room, lessonId: 0 }).then(() => {
        openDialog('授業登録')
      })
    }

    const handleClickEditLesson = async ({ shiftId, lessonId, room }: { shiftId: number; lessonId: number; room: number }): Promise<void> => {
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
      teacherSubmission,
      studentSubmission,
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
