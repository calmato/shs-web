<template>
  <v-row>
    <v-col v-for="item in items.shifts" :key="item.id" cols="12">
      <v-container class="d-flex">
        <v-col class="mb-4" cols="3">
          {{ getDay(item.date) }}
        </v-col>
        <v-row class="d-flex">
          <v-col v-if="item.isClosed" class="mt-3 red--text" cols="9"> 休日 </v-col>
          <v-col v-for="lesson in item.lessons" v-else :key="lesson.id">
            <v-btn-toggle rounded background-color="blue" multiple>
              <v-btn> {{ getLessonDuration(lesson.startTime, lesson.endTime) }} </v-btn>
            </v-btn-toggle>
          </v-col>
        </v-row>
      </v-container>
      <v-divider />
    </v-col>
  </v-row>
</template>

<script lang="ts">
import { PropType } from '@nuxtjs/composition-api'
import { defineComponent } from '@vue/composition-api'
import { SubmissionEditState } from '~/types/props/submission'

export default defineComponent({
  props: {
    items: {
      type: Object as PropType<SubmissionEditState>,
      default: () => ({}),
    },
  },

  setup() {
    const getDay = (date: string): string => {
      const year: number = parseInt(date.slice(0, 4))
      const month: number = parseInt(date.slice(5, 6)) - 1
      const day: number = parseInt(date.slice(7, 8))
      const dayOfWeekStr = ['日', '月', '火', '水', '木', '金', '土']
      const dayName = new Date(year, month, day)
      return day + '(' + dayOfWeekStr[dayName.getDay()] + ')'
    }

    const getLessonDuration = (startTime: string, endTime: string): string => {
      const convertStartTime = startTime.slice(0, 2) + ':' + startTime.slice(2, 4)
      const convertEndTime = endTime.slice(0, 2) + ':' + endTime.slice(2, 4)
      return convertStartTime + '~' + convertEndTime
    }

    return {
      getDay,
      getLessonDuration,
    }
  },
})
</script>
