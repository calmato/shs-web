<template>
  <the-top :events="getEvents()" :now="now" :start.sync="start" :end.sync="end" :detail="detail" @click="handleClick" />
</template>

<script lang="ts">
import { computed, defineComponent, ref, useStore, watch } from '@nuxtjs/composition-api'
import TheTop from '~/components/templates/TheTop.vue'
import dayjs from '~/plugins/dayjs'
import { CommonStore, LessonStore } from '~/store'
import { Date, Event, EventDetail } from '~/types/props/calendar'
import { Student, StudentMap, Teacher, TeacherMap } from '~/types/store'
import { Lesson, Subject, SubjectMap } from '~/types/store/lesson'

export default defineComponent({
  components: {
    TheTop,
  },

  setup() {
    const store = useStore()

    const now = dayjs()

    const start = ref<Date>()
    const end = ref<Date>()
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
    const teachers = computed<TeacherMap>(() => store.getters['lesson/getTeacherMap'])
    const students = computed<StudentMap>(() => store.getters['lesson/getStudentMap'])

    watch(start, (): void => {
      listLessons()
    })

    async function listLessons(): Promise<void> {
      CommonStore.startConnection()

      const format: string = 'YYYYMMDD'
      const since: string = dayjs(start.value?.date).tz().format(format)
      const until: string = dayjs(end.value?.date).tz().format(format)

      await LessonStore.listLessons({ since, until })
        .catch((err: Error) => {
          console.log('feiled to list lessons', err)
        })
        .finally(() => {
          CommonStore.endConnection()
        })
    }

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
      now,
      start,
      end,
      detail,
      getEvents,
      handleClick,
    }
  },
})
</script>
