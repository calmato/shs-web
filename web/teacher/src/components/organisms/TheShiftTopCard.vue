<template>
  <v-card>
    <v-card-title>{{ getTitle() }}</v-card-title>
    <v-card-text>
      <div v-show="isEnabledCreateLesson()">
        <span>・授業スケジュール確定状況:</span>
        <span v-if="summary.decided" class="info--text">確定済み</span>
        <span v-else class="error--text">未確定</span>
      </div>
      <div>・シフト募集期間: {{ getDate(summary.openAt) }} ~ {{ getDate(summary.endAt) }}</div>
    </v-card-text>
    <v-card-actions>
      <v-spacer />
      <v-btn color="secondary" @click="onClickEdit">募集期間の修正</v-btn>
      <v-btn v-show="isEnabledCreateLesson()" color="primary" @click="onClickNew">授業登録画面へ</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script lang="ts">
import { defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'
import { ShiftStatus, ShiftSummary } from '~/types/store'
import dayjs from '~/plugins/dayjs'

export default defineComponent({
  props: {
    summary: {
      type: Object as PropType<ShiftSummary>,
      default: () => ({
        id: 0,
        year: 0,
        month: 0,
        decided: false,
        status: ShiftStatus.UNKNOWN,
        openAt: '',
        endAt: '',
        createdAt: '',
        updatedAt: '',
      }),
    },
  },

  setup(props, { emit }: SetupContext) {
    const getTitle = (): string => {
      return `${props.summary.year}年${props.summary.month}月のシフト`
    }

    const getDate = (date: string): string => {
      const format: string = 'YYYY/MM/DD'
      if (!date || date === '') {
        return format
      }
      return dayjs(date).tz().format(format)
    }

    const getDecided = (decided: boolean): string => {
      return decided ? '確定済み' : '未確定'
    }

    const isEnabledCreateLesson = (): boolean => {
      return props.summary?.status === ShiftStatus.FINISHED || props.summary?.status === ShiftStatus.ACCEPTING
    }

    const onClickEdit = (): void => {
      emit('click:edit', props.summary)
    }

    const onClickNew = (): void => {
      emit('click:new', props.summary)
    }

    return {
      getTitle,
      getDate,
      getDecided,
      isEnabledCreateLesson,
      onClickEdit,
      onClickNew,
    }
  },
})
</script>
