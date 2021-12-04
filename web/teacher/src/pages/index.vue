<template>
  <top :events="getEvents()" @click="handleClick" />
</template>

<script lang="ts">
import { computed, defineComponent, SetupContext } from '@nuxtjs/composition-api'
import Top from '~/components/templates/Top.vue'
import { Event } from '~/types/props/calendar'
import { Student, StudentMap } from '~/types/store'
import { Lesson, Subject, SubjectMap } from '~/types/store/lesson'

export default defineComponent({
  components: {
    Top,
  },

  setup(_, { root }: SetupContext) {
    const store = root.$store

    const lessons = computed<Lesson[]>(() => store.getters['lesson/getLessons'])
    const subjects = computed<SubjectMap>(() => store.getters['lesson/getSubjectMap'])
    const students = computed<StudentMap>(() => store.getters['user/getStudentMap'])

    const getEvents = (): Event[] => {
      const events: Event[] = lessons.value.map((lesson: Lesson): Event => {
        const subject: Subject = subjects.value[lesson.subjectId]
        const student: Student = students.value[lesson.studentId]

        return {
          name: `[${subject.name}] ${student.lastname} ${student.firstname}`,
          start: lesson.startAt,
          end: lesson.endAt,
          color: subject.color,
        }
      })
      return events
    }

    const handleClick = (event: Event) => {
      console.log('debug', event)
    }

    return {
      getEvents,
      handleClick,
    }
  },
})
</script>
