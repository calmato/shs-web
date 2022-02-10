<template>
  <v-card-text class="d-flex flex-column align-stretch">
    <h4 class="d-flex align-center justify-center">
      <v-chip v-if="detail" :color="getSubjectColor(detail)" small class="mr-4">
        {{ getSubjectName(detail) }}
      </v-chip>
      <span>{{ getTime(summary.startTime) }} ~ {{ getTime(summary.endTime) }}</span>
    </h4>
    <div class="d-flex align-center justify-center mt-2">
      <a v-if="detail" class="d-block" @click="onClickEdit">
        <div class="black--text text-decoration-underline">生徒: {{ getStudentName(detail) }}</div>
        <div class="black--text text-decoration-underline">講師: {{ getTeacherName(detail) }}</div>
      </a>
      <v-btn v-if="!decided && !detail" icon color="primary" class="my-1" @click="onClickNew">
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
      return detail?.subject?.fullname || ''
    }

    const getSubjectColor = (detail: LessonDetail | undefined): string => {
      return detail?.subject?.color || ''
    }

    const getTeacherName = (detail: LessonDetail | undefined): string => {
      return detail?.teacher?.name || ''
    }

    const getStudentName = (detail: LessonDetail | undefined): string => {
      return detail?.student?.name || ''
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
      getSubjectColor,
      getTeacherName,
      getStudentName,
      onClickNew,
      onClickEdit,
    }
  },
})
</script>
