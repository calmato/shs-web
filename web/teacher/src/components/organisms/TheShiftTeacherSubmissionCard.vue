<template>
  <v-card>
    <v-toolbar color="primary" dark>提出シフト一覧 - {{ submission.name }}</v-toolbar>
    <v-card-text class="py-4">
      <div class="d-flex align-center">
        <span class="text-subtitle-1">提出シフト一覧</span>
        <span class="ml-auto text-subtitle-1 info--text">合計: {{ submission.submissionTotal }}件</span>
      </div>
      <v-data-table :headers="headers" :items="getItems()" :items-per-page="-1" hide-default-footer>
        <template #[`item.lessons`]="{ item }">
          <v-chip v-for="lesson in item.lessons" :key="lesson.id">
            {{ getDuration(lesson.startTime, lesson.endTime) }}
          </v-chip>
        </template>
      </v-data-table>
    </v-card-text>
    <v-card-actions>
      <v-spacer />
      <v-btn color="secondary" @click="onClose">閉じる</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script lang="ts">
import { defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'
import dayjs from '~/plugins/dayjs'
import { SubmissionTableItem, TableHeader } from '~/types/props/shift'
import { ShiftDetail, ShiftStatus, SubmissionStatus, TeacherSubmissionDetail } from '~/types/store'

export default defineComponent({
  props: {
    submission: {
      type: Object as PropType<TeacherSubmissionDetail>,
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
        submissionTotal: 0,
      }),
    },
  },

  setup(props, { emit }: SetupContext) {
    const headers: TableHeader[] = [
      { text: '日付', value: 'date', align: 'center' },
      { text: '時間', value: 'lessons', align: 'start' },
    ]

    const getItems = (): SubmissionTableItem[] => {
      const format: string = 'DD(ddd)'
      const items: SubmissionTableItem[] = []

      props.submission.shifts.forEach((shift: ShiftDetail) => {
        const date: string = dayjs(shift.date).tz().format(format)
        items.push({ date, lessons: shift.lessons })
      })

      return items
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
      headers,
      getItems,
      getDuration,
      onClose,
    }
  },
})
</script>
