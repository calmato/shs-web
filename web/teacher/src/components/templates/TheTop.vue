<template>
  <v-container fill-height class="px-0">
    <v-row class="justiry-center align-stretch">
      <v-dialog v-model="dialog" width="600px" scrollable @click:outside="toggleDialog">
        <the-calendar-detail :detail="detail" @click:close="toggleDialog" />
      </v-dialog>
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
            :start.sync="startAt"
            :end.sync="endAt"
            :type.sync="type"
            :events="events"
            :types="types"
            :weekdays="weekdays"
            @click="showEvent"
          />
        </v-sheet>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { computed, defineComponent, PropType, ref, SetupContext } from '@nuxtjs/composition-api'
import { Dayjs } from 'dayjs'
import TheCalendarBody from '~/components/organisms/TheCalendarBody.vue'
import TheCalendarDetail from '~/components/organisms/TheCalendarDetail.vue'
import TheCalendarHeader from '~/components/organisms/TheCalendarHeader.vue'
import { CalendarType, Date, Event, EventDetail } from '~/types/props/calendar'

export default defineComponent({
  components: {
    TheCalendarBody,
    TheCalendarDetail,
    TheCalendarHeader,
  },

  props: {
    detail: {
      type: Object as PropType<EventDetail>,
      default: () => ({}),
    },
    events: {
      type: Array as PropType<Event[]>,
      default: () => [],
    },
    now: {
      type: Object as PropType<Dayjs>,
      default: null,
    },
    start: {
      type: Object as PropType<Date>,
      default: null,
    },
    end: {
      type: Object as PropType<Date>,
      default: null,
    },
  },

  setup(props, { emit }: SetupContext) {
    const weekdays: number[] = [0, 1, 2, 3, 4, 5, 6]
    const types: CalendarType[] = [
      { name: '月', value: 'month' },
      { name: '週', value: 'week' },
      { name: '日', value: 'day' },
    ]

    const focus = ref<string>('')
    const type = ref<String>('month')
    const today = ref<String>(props.now?.tz().format('YYYY-MM-DD HH:mm:ss'))
    const dialog = ref<boolean>(false)

    const startAt = computed({
      get: () => props.start,
      set: (val: Date) => emit('update:start', val),
    })
    const endAt = computed({
      get: () => props.end,
      set: (val: Date) => emit('update:end', val),
    })

    const toggleDialog = (): void => {
      dialog.value = !dialog.value
    }

    const showEvent = (event: Event): void => {
      emit('click', event)
      toggleDialog()
    }

    const setToday = (): void => {
      focus.value = today.toString()
    }

    return {
      dialog,
      focus,
      type,
      types,
      weekdays,
      startAt,
      endAt,
      toggleDialog,
      showEvent,
      setToday,
    }
  },
})
</script>
