<template>
  <v-row class="pa-4" justify="center">
    <v-col v-for="summary in summaries" :key="summary.id" cols="12" sm="6" md="4">
      <v-card outlined elevation="4">
        <v-card-title class="justify-center">{{ getTitle(summary.year, summary.month) }}</v-card-title>
        <v-card-text class="text-center">
          提出期限:
          <span class="text-decoration-underline red--text text--lighten-1">{{ getDate(summary.endAt) }}</span>
        </v-card-text>
        <v-card-actions class="mx-8">
          <v-chip dark :color="getSubmissionStatusColor(summary.submissionStatus)">
            {{ getSubmissionStatus(summary.submissionStatus) }}
          </v-chip>
          <v-spacer />
          <v-btn :color="getDetailButtonColor(summary.submissionStatus)" @click="onClick(summary)">
            {{ getDetailButtonMessage(summary.submissionStatus) }}
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-col>
  </v-row>
</template>

<script lang="ts">
import { defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'
import { SubmissionStatus, TeacherShiftSummary } from '~/types/store'
import dayjs from '~/plugins/dayjs'

export default defineComponent({
  props: {
    summaries: {
      type: Array as PropType<TeacherShiftSummary[]>,
      default: () => [],
    },
  },

  setup(_, { emit }: SetupContext) {
    const getTitle = (year: number, month: number): string => {
      return `${year}年${month}月 シフト希望`
    }

    const getDate = (date: string): string => {
      const format: string = 'YYYY/MM/DD'
      if (!date || date === '') {
        return format
      }
      return dayjs(date).tz().format(format)
    }

    const getSubmissionStatus = (status: SubmissionStatus): string => {
      switch (status) {
        case SubmissionStatus.WAITING:
          return '未提出'
        case SubmissionStatus.SUBMITTED:
          return '提出済み'
        default:
          return ''
      }
    }

    const getSubmissionStatusColor = (status: SubmissionStatus): string => {
      switch (status) {
        case SubmissionStatus.WAITING:
          return 'red'
        case SubmissionStatus.SUBMITTED:
          return 'light-green'
        default:
          return ''
      }
    }

    const getDetailButtonMessage = (status: SubmissionStatus): string => {
      switch (status) {
        case SubmissionStatus.WAITING:
          return '入力する'
        case SubmissionStatus.SUBMITTED:
          return '修正する'
        default:
          return ''
      }
    }

    const getDetailButtonColor = (status: SubmissionStatus): string => {
      switch (status) {
        case SubmissionStatus.WAITING:
          return 'primary'
        case SubmissionStatus.SUBMITTED:
          return 'secondary'
        default:
          return ''
      }
    }

    const onClick = (summary: TeacherShiftSummary): void => {
      emit('click', summary)
    }

    return {
      getTitle,
      getDate,
      getSubmissionStatus,
      getSubmissionStatusColor,
      getDetailButtonMessage,
      getDetailButtonColor,
      onClick,
    }
  },
})
</script>
