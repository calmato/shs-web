<template>
  <v-calendar
    ref="calendar"
    v-model="calendarFocus"
    :events="events"
    :event-color="getEventColor"
    :type="calendarType"
    :now="currentTime"
    :weekdays="weekdays"
    @click:event="showEvent"
    @click:more="viewDay"
    @click:date="viewDay"
    @change="updateRange"
  />
</template>

<script lang="ts">
import { computed, defineComponent, onMounted, PropType, SetupContext } from '@nuxtjs/composition-api'
import { Dayjs } from 'dayjs'
import dayjs from '~/plugins/dayjs'
import { Event, Date, CalendarType, CalendarRef } from '~/types/props/calendar'

export default defineComponent({
  props: {
    end: {
      type: Object as PropType<Date>,
      default: () => ({
        date: '2021-07-31',
        time: '23:59:59',
        year: 2021,
        month: 8,
        weekday: 6,
        day: 31,
        hour: 23,
        minute: 59,
        future: false,
        hasDay: false,
        hasTime: false,
        past: false,
        present: false,
      }),
    },
    events: {
      type: Array as PropType<Event[]>,
      default: () => [],
    },
    focus: {
      type: String,
      default: '',
    },
    now: {
      type: Object as PropType<Dayjs>,
      default: null,
    },
    start: {
      type: Object as PropType<Date>,
      default: () => ({
        date: '2021-07-01',
        time: '00:00:00',
        year: 2021,
        month: 8,
        weekday: 4,
        day: 1,
        hour: 0,
        minute: 0,
        future: false,
        hasDay: false,
        hasTime: false,
        past: false,
        present: false,
      }),
    },
    type: {
      type: String,
      default: 'month',
    },
    types: {
      type: Array as PropType<CalendarType[]>,
      default: () => [
        { name: '月', value: 'month' },
        { name: '週', value: 'week' },
        { name: '日', value: 'day' },
      ],
    },
    weekdays: {
      type: Array as PropType<Number[]>,
      default: () => [0, 1, 2, 3, 4, 5, 6],
    },
  },

  setup(props, { emit, refs }: SetupContext) {
    const startTime = computed<Date>({
      get: () => props.start,
      set: (val: Date) => emit('update:start', val),
    })
    const endTime = computed<Date>({
      get: () => props.end,
      set: (val: Date) => emit('update:end', val),
    })
    const currentTime = computed<string>({
      get: () => props.now?.tz().format('YYYY-MM-DD HH:mm:ss') || '2021-07-23 20:00:00',
      set: (val: string) => emit('update:now', dayjs(val)),
    })
    const calendarFocus = computed<string>({
      get: () => props.focus,
      set: (val: string) => emit('update:focus', val),
    })
    const calendarType = computed<string>({
      get: () => props.type,
      set: (val: string) => emit('update:type', val),
    })

    onMounted(() => {
      calendarInstance().checkChange()
    })

    const calendarInstance = (): Vue & CalendarRef => {
      return refs.calendar as Vue & CalendarRef
    }

    const prev = (): void => {
      calendarInstance().prev()
    }

    const next = (): void => {
      calendarInstance().next()
    }

    const getEventColor = (event: Event): string => {
      return event?.color || 'primary'
    }

    const viewDay = (date: Date): void => {
      calendarFocus.value = date.date
      calendarType.value = 'day'
    }

    const showEvent = ({ event }: { event: Event }): void => {
      emit('click', event)
    }

    const updateRange = ({ start, end }: { start: Date; end: Date }): void => {
      startTime.value = start
      endTime.value = end
    }

    return {
      currentTime,
      calendarFocus,
      calendarType,
      getEventColor,
      viewDay,
      showEvent,
      updateRange,
      prev,
      next,
    }
  },
})
</script>
