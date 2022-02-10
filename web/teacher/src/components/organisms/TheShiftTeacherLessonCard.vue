<template>
  <v-card>
    <v-toolbar color="primary" dark elevation="0">授業一覧 - {{ getTeacherName(lesson.current) }}</v-toolbar>
    <v-card-text class="py-4">
      <div class="d-flex align-center">
        <span class="text-subtitle-1">授業一覧</span>
        <span class="ml-auto text-subtitle-1 info--text">合計: {{ lesson.total }}</span>
      </div>
      <v-data-table :loading="loading" :headers="headers" :items="getItems()" :items-per-page="-1" hide-default-footer>
        <template #[`item.subjectId`]="{ item }">
          <v-chip :color="getSubjectColor(item.subjectId)">{{ getSubjectName(item.subjectId) }}</v-chip>
        </template>
        <template #[`item.studentId`]="{ item }">
          {{ getStudentName(item.studentId) }}
        </template>
      </v-data-table>
    </v-card-text>
    <v-card-actions>
      <v-spacer />
      <v-btn color="primary" outlined @click="onClose">閉じる</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script lang="ts">
import { defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'
import dayjs from '~/plugins/dayjs'
import { TableHeader, TeacherLessonTableItem } from '~/types/props/shift'
import { ShiftLesson, ShiftUserLesson, Subject } from '~/types/store'

export default defineComponent({
  props: {
    loading: {
      type: Boolean,
      default: false,
    },
    lesson: {
      type: Object as PropType<ShiftUserLesson>,
      default: () => {},
    },
    subjects: {
      type: Array as PropType<Subject[]>,
      default: () => [],
    },
  },

  setup(props, { emit }: SetupContext) {
    const headers: TableHeader[] = [
      { text: '日付', value: 'date', align: 'center' },
      { text: '時間', value: 'duration', align: 'center' },
      { text: '科目', value: 'subjectId', align: 'center' },
      { text: '生徒', value: 'studentId', align: 'center' },
    ]

    const getDuration = (startAt: string, endAt: string): string => {
      const format: string = 'HH:mm'
      const startTime: string = dayjs(startAt).tz().format(format)
      const endTime: string = dayjs(endAt).tz().format(format)
      return `${startTime}~${endTime}`
    }

    const getItems = (): TeacherLessonTableItem[] => {
      const format: string = 'DD(ddd)'
      const items: TeacherLessonTableItem[] = props.lesson.lessons.map((lesson: ShiftLesson) => {
        const date: string = dayjs(lesson.startAt).tz().format(format)
        const duration: string = getDuration(lesson.startAt, lesson.endAt)
        return { date, duration, subjectId: lesson.subjectId, studentId: lesson.studentId }
      })
      return items
    }

    const getTeacherName = (teacherId: string): string => {
      return props.lesson.teachers[teacherId]?.name || ''
    }

    const getStudentName = (studentId: string): string => {
      return props.lesson.students[studentId]?.name || ''
    }

    const getSubjectName = (subjectId: number): string => {
      const subject: Subject | undefined = props.subjects.find((subject: Subject) => subject.id === subjectId)
      return subject?.fullname || ''
    }

    const getSubjectColor = (subjectId: number): string => {
      const subject: Subject | undefined = props.subjects.find((subject: Subject) => subject.id === subjectId)
      return subject?.color || ''
    }

    const onClose = (): void => {
      emit('click:close')
    }

    return {
      headers,
      getItems,
      getTeacherName,
      getStudentName,
      getSubjectName,
      getSubjectColor,
      onClose,
    }
  },
})
</script>
