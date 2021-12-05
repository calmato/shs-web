<template>
  <v-card>
    <v-toolbar color="primary" dark>{{ getLessonDate() }}の授業詳細</v-toolbar>

    <v-card-text>
      <v-row class="py-8">
        <v-col cols="12">
          <h3 class="mb-2">授業名</h3>
          <span>{{ detail.subject }}</span>
          <v-divider />
        </v-col>
        <v-col cols="12">
          <h3 class="mb-2">授業時間</h3>
          <span>{{ getLessonDate() }} {{ getLessonTime() }}</span>
          <v-divider />
        </v-col>
        <v-col cols="12">
          <h3 class="mb-2">講師名</h3>
          <span>{{ detail.teacher }}</span>
          <v-divider />
        </v-col>
        <v-col cols="12">
          <h3 class="mb-2">生徒名</h3>
          <span>{{ detail.student }}</span>
          <v-divider />
        </v-col>
        <v-col cols="12">
          <h3 class="mb-2">備考</h3>
          <span>{{ detail.remark || '未記載' }}</span>
          <v-divider />
        </v-col>
      </v-row>
    </v-card-text>

    <v-card-actions>
      <v-spacer></v-spacer>
      <v-btn color="secondary" @click="onClose">閉じる</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script lang="ts">
import { defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'
import dayjs from '~/plugins/dayjs'
import { EventDetail } from '~/types/props/calendar'

export default defineComponent({
  props: {
    detail: {
      type: Object as PropType<EventDetail>,
      default: () => ({}),
    },
  },

  setup(props, { emit }: SetupContext) {
    const getLessonDate = (): string => {
      return dayjs(props.detail.start).tz().format('MM月DD日')
    }

    const getLessonTime = (): string => {
      const format: string = 'HH:mm'
      const startTime: string = dayjs(props.detail.start).tz().format(format)
      const endTime: string = dayjs(props.detail.end).tz().format(format)
      return `${startTime} ~ ${endTime}`
    }

    const onClose = (): void => {
      emit('click:close')
    }

    return {
      getLessonDate,
      getLessonTime,
      onClose,
    }
  },
})
</script>
