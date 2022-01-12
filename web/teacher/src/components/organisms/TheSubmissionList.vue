<template>
  <div>
    <v-container v-for="shift in shifts" :key="shift.id" cols="12">
      <v-row align="center" class="mb-2">
        <v-col cols="3">{{ getDay(shift.date) }}</v-col>
        <v-col cols="9">
          <v-row class="d-flex">
            <v-col v-if="shift.isClosed" class="red--text"> 休日 </v-col>
            <v-col v-for="lesson in shift.lessons" v-else :key="lesson.id">
              <v-btn-toggle rounded background-color="blue" multiple>
                <v-btn>{{ getLessonDuration(lesson.startTime, lesson.endTime) }}</v-btn>
              </v-btn-toggle>
            </v-col>
          </v-row>
        </v-col>
      </v-row>
      <v-divider />
    </v-container>
  </div>
</template>

<script lang="ts">
import { defineComponent, PropType } from '@nuxtjs/composition-api'
import dayjs from '~/plugins/dayjs'
import { TeacherShiftDetail } from '~/types/store'

export default defineComponent({
  props: {
    shifts: {
      type: Array as PropType<TeacherShiftDetail[]>,
      default: () => [],
    },
  },

  setup() {
    const getDay = (date: string): string => {
      return dayjs(date).tz().format('DD(ddd)')
    }

    const getLessonDuration = (start: string, end: string): string => {
      const format: string = 'HH:mm'
      const startTime: string = dayjs(`2000-01-01 ${start}`).tz().format(format)
      const endTime: string = dayjs(`2000-01-01 ${end}`).tz().format(format)
      return `${startTime}~${endTime}`
    }

    return {
      getDay,
      getLessonDuration,
    }
  },
})
</script>
