<template>
  <pc-shift-detail
    :summary="summary"
    :details="details"
    :rooms="rooms"
    :teachers="teachers"
    :students="students"
    :lessons="getLessonDetails()"
  />
</template>

<script lang="ts">
import { computed, defineComponent, useAsync, useRoute, useStore } from '@nuxtjs/composition-api'
import PcShiftDetail from '~/components/templates/PcShiftDetail.vue'
import { CommonStore, ShiftStore } from '~/store'
import { Lesson, ShiftDetail, ShiftSummary, StudentShift, Subject, TeacherShift } from '~/types/store'
import { LessonDetail } from '~/types/props/shift'

export default defineComponent({
  components: {
    PcShiftDetail,
  },

  setup() {
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

    return {
      summary,
      details,
      rooms,
      teachers,
      students,
      lessons,
      getLessonDetails,
    }
  },
})
</script>
