<template>
  <div class="shift-lessons">
    <v-container v-for="shift in shifts" :key="shift.date" tag="v-row" class="pa-0 ma-0">
      <v-col cols="1" align="center" class="shift-lessons-date pt-4">{{ getDay(shift.date) }}</v-col>
      <v-col cols="11" class="d-flex flex-column" align="center">
        <!-- 休講日 -->
        <div v-if="shift.isClosed" class="pa-0 ma-0 error--text">
          <v-card tile outlined class="align-center">
            <v-card-text class="text-subtitle-2 error--text">休日</v-card-text>
          </v-card>
        </div>
        <!-- 開講日 -->
        <div
          v-for="num in rooms"
          v-else
          :key="`${shift.date}-${num}`"
          justify="start"
          class="d-flex align-stretch pa-0 ma-0"
        >
          <v-card tile outlined class="d-flex align-center col col-1">
            <v-card-text class="text-subtitle-2">ブース{{ num }}</v-card-text>
          </v-card>
          <v-card v-for="lesson in shift.lessons" :key="lesson.id" tile outlined class="shift-lessons-schedule">
            <the-shift-lesson-card
              :summary="lesson"
              :detail="getLesson(lesson, num)"
              @click:new="onClickNew"
              @click:edit="onClickEdit"
            />
          </v-card>
        </div>
      </v-col>
    </v-container>
  </div>
</template>

<script lang="ts">
import { defineComponent, PropType, SetupContext } from '@nuxtjs/composition-api'
import dayjs from '~/plugins/dayjs'
import TheShiftLessonCard from '~/components/organisms/TheShiftLessonCard.vue'
import { ShiftDetail, ShiftDetailLesson } from '~/types/store'
import { LessonDetail } from '~/types/props/shift'

export default defineComponent({
  components: {
    TheShiftLessonCard,
  },

  props: {
    rooms: {
      type: Number,
      default: 0,
    },
    shifts: {
      type: Array as PropType<ShiftDetail[]>,
      default: () => [],
    },
    lessons: {
      type: Array as PropType<LessonDetail[]>,
      default: () => [],
    },
  },

  setup(props, { emit }: SetupContext) {
    const getDay = (date: string): string => {
      return dayjs(date).tz().format('DD (ddd)')
    }

    const getLesson = (shift: ShiftDetailLesson, room: number): LessonDetail | undefined => {
      return props.lessons.find((detail: LessonDetail) => {
        return detail.lesson.shiftId === shift.id && detail.lesson.room === room
      })
    }

    const onClickNew = ({ summaryId }: { summaryId: number }): void => {
      emit('click:new', { summaryId })
    }

    const onClickEdit = ({ summaryId, lessonId }: { summaryId: number; lessonId: number }): void => {
      emit('click:edit', { summaryId, lessonId })
    }

    return {
      getDay,
      getLesson,
      onClickNew,
      onClickEdit,
    }
  },
})
</script>

<style scoped>
.shift-lessons .col {
  padding: 0;
  margin: 0;
}

.shift-lessons-date {
  border: 1px solid #e5e5e5;
}

.shift-lessons-schedule {
  width: 100%;
}
</style>
