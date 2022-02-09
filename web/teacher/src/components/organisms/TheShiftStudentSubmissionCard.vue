<template>
  <v-card>
    <v-toolbar color="primary" dark elevation="0">提出シフト一覧 - {{ submission.name }}</v-toolbar>
    <v-card-text class="py-4">
      <div class="text-subtitle-1">授業希望概要</div>
      <v-data-table
        :headers="suggestedLessonHeaders"
        :items="submission.suggestedLessons"
        :items-per-page="-1"
        hide-default-footer
      >
        <template #[`item.subjectId`]="{ item }">
          <v-chip :color="getSubjectColor(item.subjectId)">
            {{ getSubjectName(item.subjectId) }}
          </v-chip>
        </template>
      </v-data-table>
      <div class="d-flex align-center">
        <span class="text-subtitle-1">授業希望一覧</span>
        <span class="ml-auto text-subtitle-1 info--text">合計: {{ submission.submissionTotal }}件</span>
      </div>
      <v-data-table :headers="submissionHeaders" :items="getSubmissionItems()" :items-per-page="-1" hide-default-footer>
        <template #[`item.lessons`]="{ item }">
          <v-chip v-for="lesson in item.lessons" :key="lesson.id">
            {{ getDuration(lesson.startTime, lesson.endTime) }}
          </v-chip>
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
import { SubmissionTableItem, TableHeader } from '~/types/props/shift'
import { ShiftDetail, ShiftStatus, StudentSubmissionDetail, Subject, SubmissionStatus } from '~/types/store'

export default defineComponent({
  props: {
    submission: {
      type: Object as PropType<StudentSubmissionDetail>,
      default: () => ({
        id: '',
        name: '',
        nameKana: '',
        summary: {
          id: 0,
          year: 0,
          month: 0,
          shiftStatus: ShiftStatus.UNKNOWN,
          submissionStatus: SubmissionStatus.UNKNOWN,
          openAt: '',
          endAt: '',
          createdAt: '',
          updatedAt: '',
        },
        shifts: [],
        suggestedLessons: [],
        submissionTotal: 0,
      }),
    },
    subjects: {
      type: Array as PropType<Subject[]>,
      default: () => [],
    },
  },

  setup(props, { emit }: SetupContext) {
    const suggestedLessonHeaders: TableHeader[] = [
      { text: '科目', value: 'subjectId', align: 'center' },
      { text: 'コマ数', value: 'total', align: 'center' },
    ]

    const submissionHeaders: TableHeader[] = [
      { text: '日付', value: 'date', align: 'start' },
      { text: '時間', value: 'lessons', align: 'start' },
    ]

    const getSubmissionItems = (): SubmissionTableItem[] => {
      const format: string = 'DD(ddd)'
      const items: SubmissionTableItem[] = []

      props.submission.shifts.forEach((shift: ShiftDetail) => {
        const date: string = dayjs(shift.date).tz().format(format)
        items.push({ date, lessons: shift.lessons })
      })

      return items
    }

    const getSubjectName = (subjectId: number): string => {
      const subject: Subject | undefined = props.subjects.find((subject: Subject) => subject.id === subjectId)
      return subject?.fullname || ''
    }

    const getSubjectColor = (subjectId: number): string => {
      const subject: Subject | undefined = props.subjects.find((subject: Subject) => subject.id === subjectId)
      return subject?.color || ''
    }

    const getDuration = (start: string, end: string): string => {
      const format: string = 'HH:mm'
      const startTime: string = dayjs(`2000-01-01 ${start}:00`).tz().format(format)
      const endTime: string = dayjs(`2000-01-01 ${end}:00`).tz().format(format)
      return `${startTime}~${endTime}`
    }

    const onClose = (): void => {
      emit('click:close')
    }

    return {
      suggestedLessonHeaders,
      submissionHeaders,
      getSubmissionItems,
      getSubjectName,
      getSubjectColor,
      getDuration,
      onClose,
    }
  },
})
</script>
