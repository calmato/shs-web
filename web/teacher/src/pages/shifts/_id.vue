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
      @click:show-teacher-submissions="handleClickShowTeacherSubmissions"
      @click:show-teacher-lessons="handleClickShowTeacherLessons"
      @click:show-student-submissions="handleClickShowStudentSubmissions"
      @click:show-student-lessons="handleClickShowStudentLessons"
      @click:decided-lesson="handleClickDecidedLesson"
      @click:new-lesson="handleClickNewLesson"
      @click:edit-lesson="handleClickEditLesson"
    />
    <!-- スマホレイアウト -->
    <mb-shift-detail class="hidden-md-and-up" @click="handleClickTop" />
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, useAsync, useRoute, useRouter, useStore } from '@nuxtjs/composition-api'
import MbShiftDetail from '~/components/templates/MbShiftDetail.vue'
import PcShiftDetail from '~/components/templates/PcShiftDetail.vue'
import { CommonStore, ShiftStore } from '~/store'
import { Lesson, ShiftDetail, ShiftSummary, StudentShift, Subject, TeacherShift } from '~/types/store'
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

    const summary = computed<ShiftSummary>(() => store.getters['shift/getSummary'])
    const details = computed<ShiftDetail[]>(() => store.getters['shift/getDetails'])
    const rooms = computed<number>(() => store.getters['shift/getRooms'])
    const teachers = computed<TeacherShift[]>(() => store.getters['shift/getTeachers'])
    const students = computed<StudentShift[]>(() => store.getters['shift/getStudents'])
    const lessons = computed<Lesson[]>(() => store.getters['shift/getLessons'])
    const subjects = computed<Subject[]>(() => store.getters['lesson/getSubjects'])

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

    const handleClickShowTeacherSubmissions = (teacherId: string): void => {
      console.log('debug', 'show teacher submissions', teacherId)
    }

    const handleClickShowTeacherLessons = (teacherId: string): void => {
      console.log('debug', 'show teacher lessons', teacherId)
    }

    const handleClickShowStudentSubmissions = (studentId: string): void => {
      console.log('debug', 'show student submissions', studentId)
    }

    const handleClickShowStudentLessons = (studentId: string): void => {
      console.log('debug', 'show student lessons', studentId)
    }

    const handleClickDecidedLesson = (): void => {
      console.log('debug', 'decided lesson')
    }

    const handleClickNewLesson = ({ summaryId }: { summaryId: number }): void => {
      console.log('debug', 'new lessons', { summaryId })
    }

    const handleClickEditLesson = ({ summaryId, lessonId }: { summaryId: number; lessonId: number }): void => {
      console.log('debug', 'edit lessons', { summaryId, lessonId })
    }

    return {
      summary,
      details,
      rooms,
      teachers,
      students,
      lessons,
      getLessonDetails,
      handleClickTop,
      handleClickShowTeacherSubmissions,
      handleClickShowTeacherLessons,
      handleClickShowStudentSubmissions,
      handleClickShowStudentLessons,
      handleClickDecidedLesson,
      handleClickNewLesson,
      handleClickEditLesson,
    }
  },
})
</script>
