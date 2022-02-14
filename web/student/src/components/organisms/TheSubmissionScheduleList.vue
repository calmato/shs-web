<template>
  <div>
    <v-divider />
    <v-container class="d-flex">
      <v-row>
        <v-col cols="4">授業日</v-col>
        <v-col cols="8">授業希望</v-col>
      </v-row>
    </v-container>
    <v-divider />
    <v-container v-for="weekday in weekdays" :key="weekday.value" class="border">
      <v-row align="center">
        <v-col cols="4">{{ weekday.name }}</v-col>
        <v-col cols="8">
          <span v-if="isClosed(weekday.value)" class="error--text">休講日</span>
          <v-chip-group v-else rounded color="success" multiple column :value="getSelectedLessonIds(weekday.value)">
            <v-chip
              v-for="(lesson, index) in getLessons(weekday.value)"
              :key="`${weekday.value}-${index}`"
              :value="newLessonId(weekday.value, lesson.startTime)"
              :small="$vuetify.breakpoint.xs"
              class="my-2"
              @click="onClickLesson(weekday.value, index)"
            >
              {{ getDuration(lesson.startTime, lesson.endTime) }}
            </v-chip>
          </v-chip-group>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script lang="ts">
import { defineComponent, PropType } from '@nuxtjs/composition-api'
import dayjs from '~/plugins/dayjs'
import { SubmissionTemplate, SubmissionTemplateLesson } from '~/types/store'

export default defineComponent({
  props: {
    schedules: {
      type: Array as PropType<SubmissionTemplate[]>,
      default: () => [],
    },
  },

  setup(props) {
    const weekdays: { name: string; value: number }[] = [
      { name: '月', value: 1 },
      { name: '火', value: 2 },
      { name: '水', value: 3 },
      { name: '木', value: 4 },
      { name: '金', value: 5 },
      { name: '土', value: 6 },
      { name: '日', value: 0 },
    ]

    const isClosed = (weekday: number): boolean => {
      const template: SubmissionTemplate = props.schedules.find(
        (schedule: SubmissionTemplate): boolean => schedule.weekday === weekday
      )
      return template?.lessons.length === 0
    }

    const newLessonId = (weekday: number, startTime: string): string => {
      return `${weekday}-${startTime}`
    }

    const getLessons = (weekday: number): SubmissionTemplateLesson[] => {
      const template: SubmissionTemplate = props.schedules.find(
        (schedule: SubmissionTemplate): boolean => schedule.weekday === weekday
      )
      return template?.lessons || []
    }

    const getSelectedLessonIds = (weekday: number): string[] => {
      const lessonIds: string[] = []
      getLessons(weekday).forEach((lesson: SubmissionTemplateLesson): void => {
        if (!lesson.enabled) {
          return
        }
        lessonIds.push(newLessonId(weekday, lesson.startTime))
      })
      return lessonIds
    }

    const getDuration = (start: string, end: string): string => {
      const format: string = 'HH:mm'
      const startTime: string = dayjs(`2000-01-01 ${start}`).tz().format(format)
      const endTime: string = dayjs(`2000-01-01 ${end}`).tz().format(format)
      return `${startTime}~${endTime}`
    }

    const onClickLesson = (weekday: number, lessonIndex: number): void => {
      const templateIndex: number = props.schedules.findIndex(
        (schedule: SubmissionTemplate): boolean => schedule.weekday === weekday
      )
      if (templateIndex === -1) {
        console.log('template is not found')
        return
      }
      const enabled = props.schedules[templateIndex].lessons[lessonIndex].enabled
      props.schedules[templateIndex].lessons[lessonIndex].enabled = !enabled
    }

    return {
      weekdays,
      isClosed,
      newLessonId,
      getLessons,
      getSelectedLessonIds,
      getDuration,
      onClickLesson,
    }
  },
})
</script>

<style lang="scss" scoped>
.border {
  border-bottom: 1px solid #e5e5e5;
}
</style>
