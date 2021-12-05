<template>
  <top :events="getEvents()" :detail="detail" @click="handleClick" @save="handleSave" />
</template>

<script lang="ts">
import { computed, defineComponent, ref, SetupContext } from '@nuxtjs/composition-api'
import Top from '~/components/templates/Top.vue'
import { Event, EventDetail } from '~/types/props/calendar'
import { Student, StudentMap, Teacher, TeacherMap } from '~/types/store'
import { Lesson, Subject, SubjectMap } from '~/types/store/lesson'

export default defineComponent({
  components: {
    Top,
  },

  setup(_, { root }: SetupContext) {
    const store = root.$store

    const detail = ref<EventDetail>({
      lessonId: 0,
      subject: '',
      student: '',
      teacher: '',
      start: '',
      end: '',
      remark: '',
    })

    const lessons = computed<Lesson[]>(() => store.getters['lesson/getLessons'])
    const subjects = computed<SubjectMap>(() => store.getters['lesson/getSubjectMap'])
    const students = computed<StudentMap>(() => store.getters['user/getStudentMap'])
    const teachers = computed<TeacherMap>(() => store.getters['user/getTeacherMap'])

    const getEvents = (): Event[] => {
      const events: Event[] = lessons.value.map((lesson: Lesson): Event => {
        const subject: Subject = subjects.value[lesson.subjectId]
        const student: Student = students.value[lesson.studentId]

        return {
          lessonId: lesson.id,
          name: `[${subject.name}] ${student.lastname} ${student.firstname}`,
          start: lesson.startAt,
          end: lesson.endAt,
          color: subject.color,
        }
      })
      return events
    }

    const getName = (lastName: string, firstName: string): string => {
      const space: string = lastName && firstName ? ' ' : ''
      return lastName + space + firstName
    }

    const setDetail = (event: Event): void => {
      const lesson = lessons.value.find((lesson: Lesson) => {
        return lesson.id === event.lessonId
      })
      if (!lesson) {
        return
      }

      const subject: Subject = subjects.value[lesson.subjectId]
      const student: Student = students.value[lesson.studentId]
      const teacher: Teacher = teachers.value[lesson.teacherId]

      const target: EventDetail = {
        lessonId: lesson.id,
        subject: subject?.name || '',
        student: getName(student.lastname, student.firstname),
        teacher: getName(teacher.lastname, teacher.firstname),
        start: lesson.startAt,
        end: lesson.endAt,
        remark: '',
      }

      detail.value = target
    }

    const handleClick = (event: Event) => {
      setDetail(event)
    }

    const handleSave = () => {
      console.log('debug', 'save')
    }

    return {
      detail,
      getEvents,
      handleClick,
      handleSave,
    }
  },
})
</script>
