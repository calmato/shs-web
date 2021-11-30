<template>
  <v-row class="fill-height">
    <v-col>
      <v-sheet height="64">
        <the-calendar-header
          :type.sync="type"
          :now="now"
          :start="start"
          :end="end"
          :types="types"
          @click:today="setToday"
          @click:prev="$refs.calendar.prev()"
          @click:next="$refs.calendar.next()"
        />
      </v-sheet>
      <v-sheet height="600">
        <the-calendar-body
          ref="calendar"
          :focus.sync="focus"
          :now.sync="now"
          :start.sync="start"
          :end.sync="end"
          :type.sync="type"
          :events="events"
          :types="types"
          :weekdays="weekdays"
          @click="showEvent"
        />
      </v-sheet>
    </v-col>
  </v-row>
</template>

<script lang="ts">
import { defineComponent, PropType, ref, SetupContext } from '@nuxtjs/composition-api'
import dayjs from 'dayjs'
import TheCalendarBody from '~/components/organisms/TheCalendarBody.vue'
import TheCalendarHeader from '~/components/organisms/TheCalendarHeader.vue'
import { CalendarType, Date, Event } from '~/types/props/calendar'

export default defineComponent({
  components: {
    TheCalendarBody,
    TheCalendarHeader,
  },

  props: {
    events: {
      type: Array as PropType<Event[]>,
      default: () => [],
    },
  },

  setup(_, { emit }: SetupContext) {
    const now = dayjs()
    const weekdays: number[] = [0, 1, 2, 3, 4, 5, 6]
    const types: CalendarType[] = [
      { name: '月', value: 'month' },
      { name: '週', value: 'week' },
      { name: '日', value: 'day' },
    ]

    const start = ref<Date>()
    const end = ref<Date>()
    const focus = ref<string>('')
    const type = ref<String>('month')
    const today = ref<String>(now.format('YYYY-MM-DD hh:mm:ss'))

    const showEvent = (event: Event): void => {
      emit('click', event)
    }

    const setToday = (): void => {
      focus.value = today.toString()
    }

    return {
      now,
      end,
      focus,
      start,
      type,
      types,
      weekdays,
      showEvent,
      setToday,
    }
  },
})
</script>
