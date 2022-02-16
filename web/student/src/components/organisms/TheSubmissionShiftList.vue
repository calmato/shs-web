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
    <v-container v-for="shift in shifts" :key="shift.date" cols="12" class="border">
      <v-row align="center">
        <v-col cols="4">{{ getDay(shift.date) }}</v-col>
        <v-col cols="8">
          <div v-if="shift.isClosed" class="error--text">休日</div>
          <v-chip-group v-else rounded color="success" multiple column :value="selectedItems">
            <v-chip
              v-for="lesson in shift.lessons"
              :key="lesson.id"
              :value="lesson.id"
              :small="$vuetify.breakpoint.xs"
              class="my-2"
              @click="onChange(lesson.id)"
            >
              {{ getLessonDuration(lesson.startTime, lesson.endTime) }}
            </v-chip>
          </v-chip-group>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script lang="ts">
import { defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'
import dayjs from '~/plugins/dayjs'
import { SubmissionDetail } from '~/types/store'

export default defineComponent({
  props: {
    loading: {
      type: Boolean,
      default: false,
    },
    shifts: {
      type: Array as PropType<SubmissionDetail[]>,
      default: () => [],
    },
    selectedItems: {
      type: Array as PropType<number[]>,
      default: () => [],
    },
  },

  setup(_, { emit }: SetupContext) {
    const getDay = (date: string): string => {
      return dayjs(date).tz().format('DD(ddd)')
    }

    const getLessonDuration = (start: string, end: string): string => {
      const format: string = 'HH:mm'
      const startTime: string = dayjs(`2000-01-01 ${start}`).tz().format(format)
      const endTime: string = dayjs(`2000-01-01 ${end}`).tz().format(format)
      return `${startTime}~${endTime}`
    }

    const onChange = (lessonId: number): void => {
      emit('click:change-items', lessonId)
    }

    return {
      getDay,
      getLessonDuration,
      onChange,
    }
  },
})
</script>

<style scoped>
.border {
  border-bottom: 1px solid #e5e5e5;
}
</style>
