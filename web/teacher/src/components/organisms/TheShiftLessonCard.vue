<template>
  <v-card-text class="d-flex flex-column align-stretch">
    <h2>{{ getTime(summary.startTime) }} ~ {{ getTime(summary.endTime) }}</h2>
    <div v-if="decided" class="d-flex align-center justify-center mt-2">
      <div v-if="detail" class="d-block">
        <div class="info--text text--darken-2 text-subtitle-2">
          {{ getSubjectName(detail) }}
        </div>
        <div class="info--text text--darken-2 text-subtitle-2">{{ getUserName(detail) }}</div>
      </div>
    </div>
    <div v-else class="d-flex align-center justify-center mt-2">
      <a v-if="detail" class="d-block" @click="onClickEdit">
        <div class="info--text text--darken-2 text-subtitle-2">
          {{ getSubjectName(detail) }}
        </div>
        <div class="info--text text--darken-2 text-subtitle-2">{{ getUserName(detail) }}</div>
      </a>
      <v-btn v-else icon color="primary" class="my-1" @click="onClickNew">
        <v-icon>mdi-pencil</v-icon>
      </v-btn>
    </div>
  </v-card-text>
</template>

<script lang="ts">
import { defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'
import dayjs from '~/plugins/dayjs'
import { ShiftDetailLesson } from '~/types/api/v1'
import { LessonDetail } from '~/types/props/shift'

export default defineComponent({
  props: {
    summary: {
      type: Object as PropType<ShiftDetailLesson>,
      default: () => {},
    },
    detail: {
      type: Object as PropType<LessonDetail | undefined>,
      default: () => undefined,
    },
    room: {
      type: Number,
      default: 0,
    },
    decided: {
      type: Boolean,
      default: false,
    },
  },

  setup(props, { emit }: SetupContext) {
    const getTime = (time: string): string => {
      return dayjs(`2000-01-01 ${time}`).tz().format('HH:mm')
    }

    const getSubjectName = (detail: LessonDetail | undefined): string => {
      return detail?.subject?.name || ''
    }

    const getUserName = (detail: LessonDetail | undefined): string => {
      const teacher: string = detail?.teacher?.name || ''
      const student: string = detail?.student?.name || ''
      return `${student} => ${teacher}`
    }

    const onClickNew = (): void => {
      emit('click:new', { shiftId: props.summary.id, room: props.room })
    }

    const onClickEdit = (): void => {
      emit('click:edit', { shiftId: props.summary.id, lessonId: props.detail?.lesson.id || 0, room: props.room })
    }

    return {
      getTime,
      getSubjectName,
      getUserName,
      onClickNew,
      onClickEdit,
    }
  },
})
</script>
