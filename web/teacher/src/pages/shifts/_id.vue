<template>
  <div>
    <!-- PCレイアウト -->
    <pc-shift-detail
      class="hidden-sm-and-down"
      :summary="summary"
      :details="details"
      :rooms="rooms"
      :teachers="teachers"
      :students="students"
      :lessons="getLessonDetails()"
      :teacher-submission="teacherSubmission"
      :student-submission="studentSubmission"
      :teacher-submissions-dialog="teacherSubmissionsDialog"
      :teacher-lessons-dialog="teacherLessonsDialog"
      :student-submissions-dialog="studentSubmissionsDialog"
      :student-lessons-dialog="studentLessonsDialog"
      :new-lesson-dialog="newLessonDialog"
      :edit-lesson-dialog="editLessonDialog"
      @click:show-teacher-submissions="handleClickShowTeacherSubmissions"
      @click:show-teacher-lessons="handleClickShowTeacherLessons"
      @click:show-student-submissions="handleClickShowStudentSubmissions"
      @click:show-student-lessons="handleClickShowStudentLessons"
      @click:decided-lesson="handleClickDecidedLesson"
      @click:new-lesson="handleClickNewLesson"
      @click:edit-lesson="handleClickEditLesson"
      @click:close-teacher-submissions="handleCloseTeacherSubmissionsDialog"
      @click:close-teacher-lessons="handleCloseTeacherLessonsDialog"
      @click:close-student-submissions="handleCloseStudentSubmissionsDialog"
      @click:close-student-lessons="handleCloseStudentLessonsDialog"
      @click:close-new-lesson="handleCloseNewLessonDialog"
      @click:close-edit-lesson="handleCloseEditLessonDialog"
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
  ShiftSummary,
  StudentShift,
  StudentSubmissionDetail,
  Subject,
  TeacherShift,
  TeacherSubmissionDetail,
} from '~/types/store'
import { LessonDetail } from '~/types/props/shift'

export default defineComponent({
  components: {
    MbShiftDetail,
    PcShiftDetail,
  },

  setup() {
    const router = useRouter()
    const route = useRoute()
    const store = useStore()

    const teacherSubmissionsDialog = ref<boolean>(false)
    const teacherLessonsDialog = ref<boolean>(false)
    const studentSubmissionsDialog = ref<boolean>(false)
    const studentLessonsDialog = ref<boolean>(false)
    const newLessonDialog = ref<boolean>(false)
    const editLessonDialog = ref<boolean>(false)

    const summary = computed<ShiftSummary>(() => store.getters['shift/getSummary'])
    const details = computed<ShiftDetail[]>(() => store.getters['shift/getDetails'])
    const rooms = computed<number>(() => store.getters['shift/getRooms'])
    const teachers = computed<TeacherShift[]>(() => store.getters['shift/getTeachers'])
    const students = computed<StudentShift[]>(() => store.getters['shift/getStudents'])
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

    const handleClickTop = (): void => {
      router.push('/')
    }

    const handleClickShowTeacherSubmissions = async (teacherId: string): Promise<void> => {
      CommonStore.startConnection()

      const summaryId: string = route.value.params.id

      await ShiftStore.showTeacherSubmissions({ summaryId: Number(summaryId), teacherId })
        .then(() => {
          teacherSubmissionsDialog.value = true
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
      teacherLessonsDialog.value = true
    }

    const handleClickShowStudentSubmissions = async (studentId: string): Promise<void> => {
      CommonStore.startConnection()

      const summaryId: string = route.value.params.id

      await ShiftStore.showStudentSubmissions({ summaryId: Number(summaryId), studentId })
        .then(() => {
          studentSubmissionsDialog.value = true
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
      studentLessonsDialog.value = true
    }

    const handleClickDecidedLesson = (): void => {
      console.log('debug', 'decided lesson')
    }

    const handleClickNewLesson = ({ summaryId, room }: { summaryId: number; room: number }): void => {
      console.log('debug', 'new lessons', { summaryId, room })
      newLessonDialog.value = true
    }

    const handleClickEditLesson = ({
      summaryId,
      lessonId,
      room,
    }: {
      summaryId: number
      lessonId: number
      room: number
    }): void => {
      console.log('debug', 'edit lessons', { summaryId, lessonId, room })
      editLessonDialog.value = true
    }

    const handleCloseTeacherSubmissionsDialog = (): void => {
      teacherSubmissionsDialog.value = false
    }

    const handleCloseTeacherLessonsDialog = (): void => {
      teacherLessonsDialog.value = false
    }

    const handleCloseStudentSubmissionsDialog = (): void => {
      studentSubmissionsDialog.value = false
    }

    const handleCloseStudentLessonsDialog = (): void => {
      studentLessonsDialog.value = false
    }

    const handleCloseNewLessonDialog = (): void => {
      newLessonDialog.value = false
    }

    const handleCloseEditLessonDialog = (): void => {
      editLessonDialog.value = false
    }

    return {
      teacherSubmissionsDialog,
      teacherLessonsDialog,
      studentSubmissionsDialog,
      studentLessonsDialog,
      newLessonDialog,
      editLessonDialog,
      summary,
      details,
      rooms,
      teachers,
      students,
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
      handleCloseTeacherSubmissionsDialog,
      handleCloseTeacherLessonsDialog,
      handleCloseStudentSubmissionsDialog,
      handleCloseStudentLessonsDialog,
      handleCloseNewLessonDialog,
      handleCloseEditLessonDialog,
    }
  },
})
</script>
