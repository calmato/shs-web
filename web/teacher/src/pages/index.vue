<template>
  <the-top :events="getEvents()" :detail="detail" @click="handleClick" />
</template>

<script lang="ts">
import { computed, defineComponent, ref, SetupContext } from '@nuxtjs/composition-api'
import TheTop from '~/components/templates/TheTop.vue'
import { Event, EventDetail } from '~/types/props/calendar'
import { Student, StudentMap, Teacher, TeacherMap } from '~/types/store'
import { Lesson, Subject, SubjectMap } from '~/types/store/lesson'

export default defineComponent({
  components: {
    TheTop,
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
    // const teachers = computed<TeacherMap>(() => store.getters['user/getTeacherMap'])

    // mock
    const teachers = computed<TeacherMap>(() => {
      return {
        '000000000000000000001': {
          id: '000000000000000000001',
          name: '中村 太郎',
          nameKana: 'なかむら たろう',
          lastName: '中村',
          firstName: '太郎',
          lastNameKana: 'なかむら',
          firstNameKana: 'たろう',
          mail: 'teacher-001@calmato.jp',
          role: 0,
          createdAt: '',
          updatedAt: '',
        },
      }
    })

    const getEvents = (): Event[] => {
      const events: Event[] = lessons.value.map((lesson: Lesson): Event => {
        const subject: Subject = subjects.value[lesson.subjectId]
        const student: Student = students.value[lesson.studentId]

        return {
          lessonId: lesson.id,
          name: `[${subject.name}] ${student.name}`,
          start: lesson.startAt,
          end: lesson.endAt,
          color: subject.color,
        }
      })
      return events
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
        student: `${student.name} (${student.nameKana})`,
        teacher: `${teacher.name} (${teacher.nameKana})`,
        start: lesson.startAt,
        end: lesson.endAt,
        remark: '',
      }

      detail.value = target
    }

    const handleClick = (event: Event) => {
      setDetail(event)
    }

    return {
      detail,
      getEvents,
      handleClick,
    }
  },
})
</script>
